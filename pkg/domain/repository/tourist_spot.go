package repository

import "github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

type (
	TouristSpotCommandRepository interface {
		Store(touristSpot *entity.TouristSpot) error
	}

	TouristSpotQueryRepository interface {
		FindByID(id int) (*entity.TouristSpot, error)
	}
)
