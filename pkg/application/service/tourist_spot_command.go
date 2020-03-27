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
	TouristSpotCommandService interface {
		ImportFromWordpressByID(wordpressTouristSpotID int) (*entity.TouristSpot, error)
	}

	TouristSpotCommandServiceImpl struct {
		TouristSpotCommandRepository repository.TouristSpotCommandRepository
		WordpressQueryRepository     repository.WordpressQueryRepository
		WordpressService
		TransactionService
	}
)

var TouristSpotCommandServiceSet = wire.NewSet(
	wire.Struct(new(TouristSpotCommandServiceImpl), "*"),
	wire.Bind(new(TouristSpotCommandService), new(*TouristSpotCommandServiceImpl)),
)

func (r *TouristSpotCommandServiceImpl) ImportFromWordpressByID(id int) (*entity.TouristSpot, error) {
	wpTouristSpots, err := r.WordpressQueryRepository.FindLocationsByIDs([]int{id})
	if err != nil || len(wpTouristSpots) == 0 {
		return nil, serror.NewResourcesNotFoundError(err, "wordpress touristSpot(id=%d)", id)
	}

	if wpTouristSpots[0].Status != wordpress.StatusPublish {
		if err := r.TouristSpotCommandRepository.DeleteByID(id); err != nil {
			return nil, errors.Wrapf(err, "failed to delete touristSpot(id=%d)", id)
		}

		return nil, serror.New(nil, serror.CodeImportDeleted, "try to import deleted touristSpot")
	}

	var touristSpot *entity.TouristSpot
	err = r.TransactionService.Do(func(c context.Context) error {
		touristSpot, err = r.TouristSpotCommandRepository.Lock(c, id)
		if err != nil {
			if !serror.IsErrorCode(err, serror.CodeNotFound) {
				return errors.Wrap(err, "failed to get touristSpot")
			}
			touristSpot = &entity.TouristSpot{}
		}

		if err := r.WordpressService.PatchTouristSpot(touristSpot, wpTouristSpots[0]); err != nil {
			return errors.Wrap(err, "failed  to patch touristSpot")
		}

		if err := r.TouristSpotCommandRepository.Store(c, touristSpot); err != nil {
			return errors.Wrap(err, "failed to store touristSpot")
		}

		return nil
	})

	return touristSpot, nil
}
