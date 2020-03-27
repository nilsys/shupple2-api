package service

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity/wordpress"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	ThemeCategoryCommandService interface {
		ImportFromWordpress(wpCategory *wordpress.Category) (*entity.ThemeCategory, error)
	}

	ThemeCategoryCommandServiceImpl struct {
		repository.ThemeCategoryCommandRepository
		repository.ThemeCategoryQueryRepository
		repository.WordpressQueryRepository
		WordpressService
		TransactionService
	}
)

var ThemeCategoryCommandServiceSet = wire.NewSet(
	wire.Struct(new(ThemeCategoryCommandServiceImpl), "*"),
	wire.Bind(new(ThemeCategoryCommandService), new(*ThemeCategoryCommandServiceImpl)),
)

func (r *ThemeCategoryCommandServiceImpl) ImportFromWordpress(wpCategory *wordpress.Category) (*entity.ThemeCategory, error) {
	isThemeCategory, err := r.isThemeCategory(wpCategory)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decide category type")
	}
	if !isThemeCategory {
		return nil, serror.New(nil, serror.CodeInvalidCategoryType, "passed category is not theme category")
	}

	var themeCategory *entity.ThemeCategory
	err = r.TransactionService.Do(func(c context.Context) error {
		themeCategory, err = r.ThemeCategoryCommandRepository.Lock(c, wpCategory.ID)
		if err != nil {
			if !serror.IsErrorCode(err, serror.CodeNotFound) {
				return errors.Wrap(err, "failed to get themeCategory")
			}
			themeCategory = &entity.ThemeCategory{}
		}

		if err := r.WordpressService.PatchThemeCategory(themeCategory, wpCategory); err != nil {
			return errors.Wrap(err, "failed  to patch themeCategory")
		}

		if err := r.ThemeCategoryCommandRepository.Store(c, themeCategory); err != nil {
			return errors.Wrap(err, "failed to store themeCategory")
		}

		return nil
	})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return themeCategory, nil
}

/**
ThemeCategoryであるかどうかを判定
1. 親カテゴリがエリアカテゴリならエリアカテゴリ
2. 親カテゴリが存在しない場合、wordpress側でのCategoryTypeが未指定 or themeならtheme category
*/
func (r *ThemeCategoryCommandServiceImpl) isThemeCategory(wpCategory *wordpress.Category) (bool, error) {
	if wpCategory.Parent != 0 {
		_, err := r.ThemeCategoryQueryRepository.FindByID(wpCategory.Parent)
		if err != nil {
			if serror.IsErrorCode(err, serror.CodeNotFound) {
				return false, nil
			}
			return false, errors.Wrap(err, "failed to find parent theme category")
		}

		return true, nil
	}

	result := (wpCategory.Type == wordpress.CategoryTypeUndefined || wpCategory.Type == wordpress.CategoryTypeTheme)
	return result, nil
}
