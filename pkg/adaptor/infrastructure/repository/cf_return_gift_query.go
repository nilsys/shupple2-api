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

func (r *CfReturnGiftQueryRepositoryImpl) FindByID(id int) (*entity.CfReturnGift, error) {
	var row entity.CfReturnGift
	if err := r.DB(context.Background()).First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "cf_return_gift(id=%d)", id)
	}
	return &row, nil
}

func (r *CfReturnGiftQueryRepositoryImpl) FindSoldCountByReturnGiftIDs(c context.Context, ids []int) (*entity.CfReturnGiftSoldCountList, error) {
	var rows entity.CfReturnGiftSoldCountList

	if err := r.DB(c).Table("payment_cf_return_gift").Select("cf_return_gift_id, SUM(amount) AS sold_count").Group("cf_return_gift_id").Where("cf_return_gift_id IN (?)", ids).Find(&rows.List).Error; err != nil {
		return nil, errors.Wrap(err, "failed find cf_return_gift sold count")
	}

	return &rows, nil
}

func (r *CfReturnGiftQueryRepositoryImpl) FindByCfProjectID(projectID int) (*entity.CfReturnGiftWithCountList, error) {
	var rows entity.CfReturnGiftWithCountList

	if err := r.DB(context.Background()).
		Select("*").
		Joins("LEFT JOIN (SELECT payment_cf_return_gift.cf_return_gift_id AS id, COUNT(DISTINCT user_id) AS supporter_count, SUM(payment_cf_return_gift.amount) AS sold_count FROM payment INNER JOIN payment_cf_return_gift ON payment.id = payment_cf_return_gift.payment_id GROUP BY payment_cf_return_gift.cf_return_gift_id) pc ON cf_return_gift.id = pc.id INNER JOIN cf_return_gift_snapshot ON cf_return_gift.latest_snapshot_id = cf_return_gift_snapshot.id").
		Where("cf_return_gift.cf_project_id = ?", projectID).Order("cf_return_gift_snapshot.sort_order").Find(&rows.List).Error; err != nil {
		return nil, errors.Wrap(err, "failed find cf_return_gift")
	}

	return &rows, nil
}
