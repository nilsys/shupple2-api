package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
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
		WordpressService             WordpressService
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

	touristSpot, err := r.WordpressService.ConvertLocation(wpTouristSpots[0])
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert touristSpot")
	}

	if err := r.TouristSpotCommandRepository.Store(touristSpot); err != nil {
		return nil, errors.Wrap(err, "failed to store touristSpot")
	}

	return touristSpot, nil
}
