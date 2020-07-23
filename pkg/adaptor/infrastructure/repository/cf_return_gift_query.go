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

func (r *CfReturnGiftQueryRepositoryImpl) FindByCfProjectID(projectID int) (*entity.CfReturnGiftList, error) {
	var rows entity.CfReturnGiftList

	if err := r.DB(context.Background()).Where("cf_project_id = ?", projectID).Find(&rows.List).Error; err != nil {
		return nil, errors.Wrap(err, "failed find cf_return_gift")
	}

	return &rows, nil
}
