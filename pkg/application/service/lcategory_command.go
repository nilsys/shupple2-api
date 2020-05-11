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
	SpotCategoryCommandService interface {
		ImportFromWordpressByID(wordpressSpotCategoryID int) error
	}

	SpotCategoryCommandServiceImpl struct {
		SpotCategoryCommandRepository repository.SpotCategoryCommandRepository
		WordpressQueryRepository      repository.WordpressQueryRepository
		WordpressService
		TransactionService
	}
)

var SpotCategoryCommandServiceSet = wire.NewSet(
	wire.Struct(new(SpotCategoryCommandServiceImpl), "*"),
	wire.Bind(new(SpotCategoryCommandService), new(*SpotCategoryCommandServiceImpl)),
)

func (r *SpotCategoryCommandServiceImpl) ImportFromWordpressByID(id int) error {
	wpSpotCategory, err := r.WordpressQueryRepository.FindLocationCategoryByID(id)
	if err != nil {
		return errors.Wrapf(err, "failed to get wordpress spotCategory(id=%d)", id)
	}

	var spotCategory *entity.SpotCategory
	return r.TransactionService.Do(func(c context.Context) error {
		spotCategory, err = r.SpotCategoryCommandRepository.Lock(c, id)
		if err != nil {
			if !serror.IsErrorCode(err, serror.CodeNotFound) {
				return errors.Wrap(err, "failed to get spotCategory")
			}
			spotCategory = &entity.SpotCategory{}
		}

		if err := r.WordpressService.PatchSpotCategory(spotCategory, wpSpotCategory); err != nil {
			return errors.Wrap(err, "failed  to patch spotCategory")
		}

		if err := r.SpotCategoryCommandRepository.Store(c, spotCategory); err != nil {
			return errors.Wrap(err, "failed to store spotCategory")
		}

		return nil
	})
}
