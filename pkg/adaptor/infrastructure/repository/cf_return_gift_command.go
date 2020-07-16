package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type CfReturnGiftCommandRepositoryImpl struct {
	DAO
}

var CfReturnGiftCommandRepositorySet = wire.NewSet(
	wire.Struct(new(CfReturnGiftCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.CfReturnGiftCommandRepository), new(*CfReturnGiftCommandRepositoryImpl)),
)

func (r *CfReturnGiftCommandRepositoryImpl) Store(cfReturnGift *entity.CfReturnGift) error {
	return Transaction(r.DB(context.Background()), func(db *gorm.DB) error {
		if err := db.Set("gorm:insert_modifier", "ignore").Create(&cfReturnGift.CfReturnGiftTiny).Error; err != nil {
			return errors.Wrap(err, "failed to insert cf_return_gift")
		}

		cfReturnGift.Snapshot.SnapshotID = 0
		if err := db.Create(&cfReturnGift.Snapshot).Error; err != nil {
			return errors.Wrap(err, "failed to insert cf_return_gift_snapshot")
		}

		if err := db.Exec("UPDATE cf_return_gift SET latest_snapshot_id = ? WHERE id = ?", cfReturnGift.Snapshot.SnapshotID, cfReturnGift.ID).Error; err != nil {
			return errors.Wrap(err, "failed to update latest_snapshot_id")
		}

		return nil
	})
}

func (r *CfReturnGiftCommandRepositoryImpl) LockByIDs(c context.Context, ids []int) (*entity.CfReturnGiftList, error) {
	var rows entity.CfReturnGiftList
	if err := r.LockDB(c).Where("id IN (?)", ids).Find(&rows.List).Error; err != nil {
		return nil, errors.Wrap(err, "failed to lock cf_return_gift")
	}
	return &rows, nil
}

func (r *CfReturnGiftCommandRepositoryImpl) Lock(c context.Context, id int) (*entity.CfReturnGift, error) {
	var rows entity.CfReturnGift
	if err := r.LockDB(c).Find(&rows, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "cf_snapshot(id=%d)", id)
	}
	return &rows, nil
}

func (r *CfReturnGiftCommandRepositoryImpl) UndeleteByID(c context.Context, id int) error {
	e := &entity.CfReturnGiftTiny{}
	e.ID = id
	return errors.Wrapf(
		r.DB(c).Unscoped().Model(e).Update("deleted_at", gorm.Expr("NULL")).Error,
		"failed to cfproject post(id=%d)", id)
}

func (r *CfReturnGiftCommandRepositoryImpl) DeleteByID(id int) error {
	return errors.Wrapf(r.DB(context.Background()).Delete(&entity.CfReturnGiftTiny{ID: id}).Error, "failed to delete cfproject(id=%d)", id)
}
