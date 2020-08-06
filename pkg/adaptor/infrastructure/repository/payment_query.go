package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type PaymentQueryRepositoryImpl struct {
	DAO
}

var PaymentQueryRepositorySet = wire.NewSet(
	wire.Struct(new(PaymentQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.PaymentQueryRepository), new(*PaymentQueryRepositoryImpl)),
)

func (r *PaymentQueryRepositoryImpl) FindByUserID(userID int, query *query.FindListPaginationQuery) (*entity.PaymentList, error) {
	var rows entity.PaymentList
	if err := r.DB(context.Background()).Where("user_id = ?", userID).
		Offset(query.Offset).Limit(query.Limit).Order("created_at DESC").
		Find(&rows.List).Offset(0).Count(&rows.TotalNumber).Error; err != nil {
		return nil, errors.Wrap(err, "failed find payment.user_id")
	}
	return &rows, nil
}

func (r *PaymentQueryRepositoryImpl) FindByID(id int) (*entity.Payment, error) {
	var row entity.Payment
	if err := r.DB(context.Background()).First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "payment(id=%d)", id)
	}
	return &row, nil
}

func (r *PaymentQueryRepositoryImpl) FindPaymentCfReturnGiftByPaymentIDAndCfReturnGift(paymentID, cfReturnGiftID int) (*entity.PaymentCfReturnGift, error) {
	var row entity.PaymentCfReturnGift
	if err := r.DB(context.Background()).Where("payment_id = ? AND cf_return_gift_id = ?", paymentID, cfReturnGiftID).First(&row).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "payment_cf_return_gift(payment_id=%d,cf_return_gift_id=%d)")
	}
	return &row, nil
}

func (r *PaymentQueryRepositoryImpl) FindTinyByID(id int) (*entity.PaymentTiny, error) {
	var row entity.PaymentTiny
	if err := r.DB(context.Background()).First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "payment(id=%d)", id)
	}
	return &row, nil
}

func (r *PaymentQueryRepositoryImpl) LockPaymentCfReturnGift(c context.Context, paymentID, cfReturnGiftID int) (*entity.PaymentCfReturnGift, error) {
	var row entity.PaymentCfReturnGift
	if err := r.LockDB(c).Where("payment_id = ? AND cf_return_gift_id = ?", paymentID, cfReturnGiftID).First(&row).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "payment_cf_return_gift(payment_id=%d,cf_return_gift_id=%d)")
	}
	return &row, nil
}
