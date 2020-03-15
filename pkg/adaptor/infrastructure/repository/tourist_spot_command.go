package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type TouristSpotCommandRepositoryImpl struct {
	DAO
}

var TouristSpotCommandRepositorySet = wire.NewSet(
	wire.Struct(new(TouristSpotCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.TouristSpotCommandRepository), new(*TouristSpotCommandRepositoryImpl)),
)

func (r *TouristSpotCommandRepositoryImpl) Store(touristSpot *entity.TouristSpot) error {
	return errors.Wrap(r.DB(context.TODO()).Save(touristSpot).Error, "failed to save touristSpot")
}

func (r *TouristSpotCommandRepositoryImpl) UpdateScoreByID(c context.Context, id int) error {
	if err := r.DB(c).Exec("UPDATE  tourist_spot SET rate = (select AVG(score) from review where tourist_spot_id = ?) WHERE id = ?;", id, id).Error; err != nil {
		return errors.Wrap(err, "failed to find or create hashtag_category")
	}
	return nil
}

func (r *TouristSpotCommandRepositoryImpl) DeleteByID(id int) error {
	e := &entity.TouristSpot{}
	e.ID = id
	return errors.Wrapf(r.DB(context.TODO()).Delete(e).Error, "failed to delete touristSpot(id=%d)", id)
}
