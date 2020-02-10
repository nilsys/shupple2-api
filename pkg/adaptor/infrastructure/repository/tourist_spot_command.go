package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type TouristSpotCommandRepositoryImpl struct {
	DB *gorm.DB
}

var TouristSpotCommandRepositorySet = wire.NewSet(
	wire.Struct(new(TouristSpotCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.TouristSpotCommandRepository), new(*TouristSpotCommandRepositoryImpl)),
)

func (r *TouristSpotCommandRepositoryImpl) Store(touristSpot *entity.TouristSpot) error {
	return errors.Wrap(r.DB.Save(touristSpot).Error, "failed to save touristSpot")
}
