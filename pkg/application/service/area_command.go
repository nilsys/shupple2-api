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
	AreaCategoryCommandService interface {
		ImportFromWordpress(wpCategory *wordpress.Category) (*entity.AreaCategory, error)
		Delete(id int) error
	}

	AreaCategoryCommandServiceImpl struct {
		repository.AreaCategoryCommandRepository
		repository.AreaCategoryQueryRepository
		repository.WordpressQueryRepository
		WordpressService
		TransactionService
	}
)

var AreaCategoryCommandServiceSet = wire.NewSet(
	wire.Struct(new(AreaCategoryCommandServiceImpl), "*"),
	wire.Bind(new(AreaCategoryCommandService), new(*AreaCategoryCommandServiceImpl)),
)

func (r *AreaCategoryCommandServiceImpl) Delete(id int) error {
	return r.AreaCategoryCommandRepository.DeleteByID(id)
}

func (r *AreaCategoryCommandServiceImpl) ImportFromWordpress(wpCategory *wordpress.Category) (*entity.AreaCategory, error) {
	isAreaCategory, err := r.isAreaCategory(wpCategory)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decide category type")
	}
	if !isAreaCategory {
		return nil, serror.New(nil, serror.CodeInvalidCategoryType, "passed category is not area category")
	}

	var areaCategory *entity.AreaCategory
	err = r.TransactionService.Do(func(c context.Context) error {
		areaCategory, err = r.AreaCategoryCommandRepository.Lock(c, wpCategory.ID)
		if err != nil {
			if !serror.IsErrorCode(err, serror.CodeNotFound) {
				return errors.Wrap(err, "failed to get areaCategory")
			}
			areaCategory = &entity.AreaCategory{}
		}

		if err := r.WordpressService.PatchAreaCategory(areaCategory, wpCategory); err != nil {
			return errors.Wrap(err, "failed  to patch areaCategory")
		}

		if err := r.AreaCategoryCommandRepository.Store(c, areaCategory); err != nil {
			return errors.Wrap(err, "failed to store areaCategory")
		}

		return nil
	})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return areaCategory, nil
}

/**
AreaCategoryであるかどうかを判定
1. 親カテゴリがエリアカテゴリならエリアカテゴリ
2. 親カテゴリが存在しない場合、wordpress側でのCategoryTypeがjapan or worldならAreaCategory
*/
func (r *AreaCategoryCommandServiceImpl) isAreaCategory(wpCategory *wordpress.Category) (bool, error) {
	if wpCategory.Parent != 0 {
		_, err := r.AreaCategoryQueryRepository.FindByID(wpCategory.Parent)
		if err != nil {
			if serror.IsErrorCode(err, serror.CodeNotFound) {
				return false, nil
			}
			return false, errors.Wrap(err, "failed to find parent area category")
		}

		return true, nil
	}

	result := wpCategory.Type == wordpress.CategoryTypeJapan || wpCategory.Type == wordpress.CategoryTypeWorld
	return result, nil
}
