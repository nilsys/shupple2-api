package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	CategoryCommandService interface {
		ImportFromWordpressByID(wordpressCategoryID int, deleted bool) error
	}

	CategoryCommandServiceImpl struct {
		AreaCategoryCommandService
		ThemeCategoryCommandService
		repository.WordpressQueryRepository
	}
)

var CategoryCommandServiceSet = wire.NewSet(
	wire.Struct(new(CategoryCommandServiceImpl), "*"),
	wire.Bind(new(CategoryCommandService), new(*CategoryCommandServiceImpl)),
)

func (r *CategoryCommandServiceImpl) ImportFromWordpressByID(id int, deleted bool) error {
	if deleted {
		if err := r.AreaCategoryCommandService.Delete(id); err != nil {
			return errors.Wrap(err, "failed to delete wordpress category")
		}
		if err := r.ThemeCategoryCommandService.Delete(id); err != nil {
			return errors.Wrap(err, "failed to delete wordpress category")
		}
		return nil
	}

	wpCategory, err := r.WordpressQueryRepository.FindCategoryByID(id)
	if err != nil {
		return errors.Wrapf(err, "failed to get wordpress category(id=%d)", id)
	}

	_, err = r.AreaCategoryCommandService.ImportFromWordpress(wpCategory)
	if err != nil {
		if !serror.IsErrorCode(err, serror.CodeInvalidCategoryType) {
			return errors.Wrap(err, "failed to import category as area category")
		}

		_, err := r.ThemeCategoryCommandService.ImportFromWordpress(wpCategory)
		return errors.Wrap(err, "failed to import category as theme category")
	}

	return nil
}
