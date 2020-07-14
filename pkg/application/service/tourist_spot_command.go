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

const (
	initializeRate = 3
)

func (r *TouristSpotCommandServiceImpl) ImportFromWordpressByID(id int) (*entity.TouristSpot, error) {
	wpTouristSpot, err := r.WordpressQueryRepository.FindLocationByID(id)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get wordpress touristSpot(id=%d)", id)
	}

	if wpTouristSpot.Status != wordpress.StatusPublish {
		if err := r.TouristSpotCommandRepository.DeleteByID(id); err != nil {
			return nil, errors.Wrapf(err, "failed to delete touristSpot(id=%d)", id)
		}

		return nil, serror.New(nil, serror.CodeImportDeleted, "try to import deleted touristSpot")
	}

	var touristSpot *entity.TouristSpot
	err = r.TransactionService.Do(func(c context.Context) error {
		if err := r.TouristSpotCommandRepository.UndeleteByID(c, id); err != nil {
			return errors.Wrapf(err, "failed to undelete toursist_spot(id=%d)", id)
		}

		touristSpot, err = r.TouristSpotCommandRepository.Lock(c, id)
		if err != nil {
			if !serror.IsErrorCode(err, serror.CodeNotFound) {
				return errors.Wrap(err, "failed to get touristSpot")
			}
			touristSpot = &entity.TouristSpot{}
			touristSpot.Rate = initializeRate
			touristSpot.VendorRate = initializeRate
		}

		if err := r.WordpressService.PatchTouristSpot(touristSpot, wpTouristSpot); err != nil {
			return errors.Wrap(err, "failed  to patch touristSpot")
		}

		if err := r.TouristSpotCommandRepository.Store(c, touristSpot); err != nil {
			return errors.Wrap(err, "failed to store touristSpot")
		}

		return nil
	})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return touristSpot, nil
}
