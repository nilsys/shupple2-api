package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type VlogCommandRepositoryImpl struct {
	DAO
}

var VlogCommandRepositorySet = wire.NewSet(
	wire.Struct(new(VlogCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.VlogCommandRepository), new(*VlogCommandRepositoryImpl)),
)

func (r *VlogCommandRepositoryImpl) Lock(c context.Context, id int) (*entity.Vlog, error) {
	var row entity.Vlog
	if err := r.LockDB(c).First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "vlog(id=%d)", id)
	}
	return &row, nil
}

func (r *VlogCommandRepositoryImpl) Store(c context.Context, vlog *entity.Vlog) error {
	return errors.Wrap(r.DB(c).Save(vlog).Error, "failed to save vlog")
}

func (r *VlogCommandRepositoryImpl) DeleteByID(id int) error {
	e := &entity.Vlog{}
	e.ID = id
	return errors.Wrapf(r.DB(context.Background()).Delete(e).Error, "failed to delete vlog(id=%d)", id)
}

func (r *VlogCommandRepositoryImpl) UpdateViewsByID(id, views int) error {
	if err := r.DB(context.Background()).Exec("UPDATE vlog SET views = ?, updated_at = updated_at WHERE id = ?", views, id).Error; err != nil {
		return errors.Wrap(err, "failed to update vlog.views")
	}
	return nil
}

func (r *VlogCommandRepositoryImpl) IncrementFavoriteCount(c context.Context, vlogID int) error {
	return errors.Wrap(r.DB(c).Exec("UPDATE vlog SET favorite_count = favorite_count + 1 WHERE id = ?", vlogID).Error, "failed to update")
}

func (r *VlogCommandRepositoryImpl) DecrementFavoriteCount(c context.Context, vlogID int) error {
	return errors.Wrap(r.DB(c).Exec("UPDATE vlog SET favorite_count = favorite_count - 1 WHERE id = ?", vlogID).Error, "failed to update")
}

func (r *VlogCommandRepositoryImpl) UpdateMonthlyViewsByID(id, views int) error {
	if err := r.DB(context.Background()).Exec("UPDATE vlog SET monthly_views = ?, updated_at = updated_at WHERE id = ?", views, id).Error; err != nil {
		return errors.Wrap(err, "failed to update vlog.monthly_views")
	}
	return nil
}

func (r *VlogCommandRepositoryImpl) UpdateWeeklyViewsByID(id, views int) error {
	if err := r.DB(context.Background()).Exec("UPDATE vlog SET weekly_views = ?, updated_at = updated_at WHERE id = ?", views, id).Error; err != nil {
		return errors.Wrap(err, "failed to update vlog.weekly_views")
	}
	return nil
}
