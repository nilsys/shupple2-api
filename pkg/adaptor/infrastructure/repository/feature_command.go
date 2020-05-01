package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type FeatureCommandRepositoryImpl struct {
	DAO
}

var FeatureCommandRepositorySet = wire.NewSet(
	wire.Struct(new(FeatureCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.FeatureCommandRepository), new(*FeatureCommandRepositoryImpl)),
)

func (r *FeatureCommandRepositoryImpl) Lock(c context.Context, id int) (*entity.Feature, error) {
	var row entity.Feature
	if err := r.LockDB(c).First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "feature(id=%d)", id)
	}
	return &row, nil
}

func (r *FeatureCommandRepositoryImpl) Store(c context.Context, feature *entity.Feature) error {
	return errors.Wrap(r.DB(c).Save(feature).Error, "failed to save feature")
}

func (r *FeatureCommandRepositoryImpl) DeleteByID(id int) error {
	e := &entity.Feature{}
	e.ID = id
	return errors.Wrapf(r.DB(context.Background()).Delete(e).Error, "failed to delete feature(id=%d)", id)
}

func (r *FeatureCommandRepositoryImpl) UpdateViewsByID(id, views int) error {
	if err := r.DB(context.Background()).Exec("UPDATE feature SET views = ? WHERE id = ?", views, id).Error; err != nil {
		return errors.Wrap(err, "failed to update feature.views")
	}
	return nil
}
