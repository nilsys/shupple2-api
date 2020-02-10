package service

import (
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
		WordpressService          WordpressService
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

	category := r.WordpressService.ConvertCategory(wpCategories[0])
	if err := r.CategoryCommandRepository.Store(category); err != nil {
		return nil, errors.Wrap(err, "failed to store category")
	}

	return category, nil
}
