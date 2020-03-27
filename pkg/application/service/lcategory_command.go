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
	LcategoryCommandService interface {
		ImportFromWordpressByID(wordpressLcategoryID int) (*entity.Lcategory, error)
	}

	LcategoryCommandServiceImpl struct {
		LcategoryCommandRepository repository.LcategoryCommandRepository
		WordpressQueryRepository   repository.WordpressQueryRepository
		WordpressService
		TransactionService
	}
)

var LcategoryCommandServiceSet = wire.NewSet(
	wire.Struct(new(LcategoryCommandServiceImpl), "*"),
	wire.Bind(new(LcategoryCommandService), new(*LcategoryCommandServiceImpl)),
)

func (r *LcategoryCommandServiceImpl) ImportFromWordpressByID(id int) (*entity.Lcategory, error) {
	wpLcategories, err := r.WordpressQueryRepository.FindLocationCategoriesByIDs([]int{id})
	if err != nil || len(wpLcategories) == 0 {
		return nil, serror.NewResourcesNotFoundError(err, "wordpress lcategory(id=%d)", id)
	}

	var lcategory *entity.Lcategory
	err = r.TransactionService.Do(func(c context.Context) error {
		lcategory, err = r.LcategoryCommandRepository.Lock(c, id)
		if err != nil {
			if !serror.IsErrorCode(err, serror.CodeNotFound) {
				return errors.Wrap(err, "failed to get lcategory")
			}
			lcategory = &entity.Lcategory{}
		}

		if err := r.WordpressService.PatchLcategory(lcategory, wpLcategories[0]); err != nil {
			return errors.Wrap(err, "failed  to patch lcategory")
		}

		if err := r.LcategoryCommandRepository.Store(c, lcategory); err != nil {
			return errors.Wrap(err, "failed to store lcategory")
		}

		return nil
	})

	return lcategory, nil
}
