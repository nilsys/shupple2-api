package service

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	CategoryCommandService interface {
		ImportFromWordpressByID(wordpressCategoryID int) (*entity.Category, error)
	}

	CategoryCommandServiceImpl struct {
		CategoryCommandRepository repository.CategoryCommandRepository
		WordpressQueryRepository  repository.WordpressQueryRepository
		WordpressService
		TransactionService
	}
)

var CategoryCommandServiceSet = wire.NewSet(
	wire.Struct(new(CategoryCommandServiceImpl), "*"),
	wire.Bind(new(CategoryCommandService), new(*CategoryCommandServiceImpl)),
)

func (r *CategoryCommandServiceImpl) ImportFromWordpressByID(id int) (*entity.Category, error) {
	wpCategories, err := r.WordpressQueryRepository.FindCategoriesByIDs([]int{id})
	if err != nil || len(wpCategories) == 0 {
		return nil, serror.NewResourcesNotFoundError(err, "wordpress category(id=%d)", id)
	}

	var category *entity.Category
	err = r.TransactionService.Do(func(c context.Context) error {
		category, err = r.CategoryCommandRepository.Lock(c, id)
		if err != nil {
			if !serror.IsErrorCode(err, serror.CodeNotFound) {
				return errors.Wrap(err, "failed to get category")
			}
			category = &entity.Category{}
		}

		if err := r.WordpressService.PatchCategory(category, wpCategories[0]); err != nil {
			return errors.Wrap(err, "failed  to patch category")
		}

		if err := r.CategoryCommandRepository.Store(c, category); err != nil {
			return errors.Wrap(err, "failed to store category")
		}

		return nil
	})

	return category, nil
}
