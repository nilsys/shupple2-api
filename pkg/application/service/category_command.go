package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	CategoryCommandService interface {
		ImportFromWordpressByID(wordpressCategoryID int) error
	}

	CategoryCommandServiceImpl struct {
		AreaCategoryCommandService
		ThemeCategoryCommandService
		repository.AreaCategoryCommandRepository
		repository.WordpressQueryRepository
	}
)

var CategoryCommandServiceSet = wire.NewSet(
	wire.Struct(new(CategoryCommandServiceImpl), "*"),
	wire.Bind(new(CategoryCommandService), new(*CategoryCommandServiceImpl)),
)

func (r *CategoryCommandServiceImpl) ImportFromWordpressByID(id int) error {
	wpCategories, err := r.WordpressQueryRepository.FindCategoriesByIDs([]int{id})
	if err != nil || len(wpCategories) == 0 {
		return serror.NewResourcesNotFoundError(err, "wordpress category(id=%d)", id)
	}

	_, err = r.AreaCategoryCommandService.ImportFromWordpress(wpCategories[0])
	if err != nil {
		if !serror.IsErrorCode(err, serror.CodeInvalidCategoryType) {
			return errors.Wrap(err, "failed to import category as area category")
		}

		_, err := r.ThemeCategoryCommandService.ImportFromWordpress(wpCategories[0])
		return errors.Wrap(err, "failed to import category as theme category")
	}

	return nil
}
