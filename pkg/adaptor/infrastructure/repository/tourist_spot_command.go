package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
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

const (
	defaultTouristSpotRate = 3
)

func (r *TouristSpotCommandRepositoryImpl) Lock(c context.Context, id int) (*entity.TouristSpot, error) {
	var row entity.TouristSpot
	if err := r.LockDB(c).First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "touristSpot(id=%d)", id)
	}
	return &row, nil
}

func (r *TouristSpotCommandRepositoryImpl) Store(c context.Context, touristSpot *entity.TouristSpot) error {
	return errors.Wrap(r.DB(c).Save(touristSpot).Error, "failed to save touristSpot")
}

// review.scoreの平均値でrateを更新
// 0(reviewが0件)になった場合0の代わりにdefaultTouristSpotRateで更新する
func (r *TouristSpotCommandRepositoryImpl) UpdateScoreByID(c context.Context, id int64) error {
	if err := r.DB(c).Exec("UPDATE tourist_spot SET rate = COALESCE( NULLIF( (select AVG(score) from review where tourist_spot_id = ? AND deleted_at IS NULL), 0), ?) WHERE id = ?;", id, defaultTouristSpotRate, id).Error; err != nil {
		return errors.Wrap(err, "failed to update tourist spot rate")
	}
	return nil
}

func (r *TouristSpotCommandRepositoryImpl) UndeleteByID(c context.Context, id int) error {
	e := &entity.TouristSpot{}
	e.ID = id
	return errors.Wrapf(
		r.DB(c).Unscoped().Model(e).Update("deleted_at", gorm.Expr("NULL")).Error,
		"failed to delete tourist_spot(id=%d)", id)
}

func (r *TouristSpotCommandRepositoryImpl) DeleteByID(id int) error {
	e := &entity.TouristSpot{}
	e.ID = id
	return errors.Wrapf(r.DB(context.TODO()).Delete(e).Error, "failed to delete touristSpot(id=%d)", id)
}
