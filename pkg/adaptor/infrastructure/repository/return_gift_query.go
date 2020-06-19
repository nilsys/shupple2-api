package repository

import (
	"context"

	"github.com/pkg/errors"

	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type ReturnGiftQueryRepositoryImpl struct {
	DAO
}

var ReturnGiftQueryRepositorySet = wire.NewSet(
	wire.Struct(new(ReturnGiftQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.ReturnGiftQueryRepository), new(*ReturnGiftQueryRepositoryImpl)),
)

func (r *ReturnGiftQueryRepositoryImpl) LockReturnGiftsWitLatestSummary(c context.Context, ids []int) (*entity.CfReturnGiftList, error) {
	var rows entity.CfReturnGiftList
	if err := r.LockDB(c).Where("id IN (?)", ids).Find(&rows.List).Error; err != nil {
		return nil, errors.Wrap(err, "failed find return_gift")
	}
	return &rows, nil
}

func (r *ReturnGiftQueryRepositoryImpl) FindSoldCountByReturnGiftIDs(c context.Context, ids []int) (*entity.CfReturnGiftSoldCountList, error) {
	var rows entity.CfReturnGiftSoldCountList

	if err := r.DB(c).Table("payment_return_gift").Select("return_gift_id, SUM(amount) AS sold_count").Group("return_gift_id").Where("return_gift_id IN (?)", ids).Find(&rows.List).Error; err != nil {
		return nil, errors.Wrap(err, "failed find return_gift sold count")
	}

	return &rows, nil
}
