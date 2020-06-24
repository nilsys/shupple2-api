package repository

import (
	"context"

	"github.com/pkg/errors"

	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type CfReturnGiftQueryRepositoryImpl struct {
	DAO
}

var CfReturnGiftQueryRepositorySet = wire.NewSet(
	wire.Struct(new(CfReturnGiftQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.CfReturnGiftQueryRepository), new(*CfReturnGiftQueryRepositoryImpl)),
)

func (r *CfReturnGiftQueryRepositoryImpl) LockCfReturnGiftList(c context.Context, ids []int) (*entity.CfReturnGiftList, error) {
	var rows entity.CfReturnGiftList
	if err := r.LockDB(c).Where("id IN (?)", ids).Find(&rows.List).Error; err != nil {
		return nil, errors.Wrap(err, "failed find cf_return_gift")
	}
	return &rows, nil
}

func (r *CfReturnGiftQueryRepositoryImpl) FindSoldCountByReturnGiftIDs(c context.Context, ids []int) (*entity.CfReturnGiftSoldCountList, error) {
	var rows entity.CfReturnGiftSoldCountList

	if err := r.DB(c).Table("payment_cf_return_gift").Select("cf_return_gift_id, SUM(amount) AS sold_count").Group("cf_return_gift_id").Where("cf_return_gift_id IN (?)", ids).Find(&rows.List).Error; err != nil {
		return nil, errors.Wrap(err, "failed find cf_return_gift sold count")
	}

	return &rows, nil
}
