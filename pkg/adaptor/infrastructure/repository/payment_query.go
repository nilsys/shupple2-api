package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type PaymentQueryRepositoryImpl struct {
	DB *gorm.DB
}

var PaymentQueryRepositorySet = wire.NewSet(
	wire.Struct(new(PaymentQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.PaymentQueryRepository), new(*PaymentQueryRepositoryImpl)),
)

func (r *PaymentQueryRepositoryImpl) FindByID(id int) (*entity.Payment, error) {
	var row entity.Payment
	if err := r.DB.First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "payment(id=%d)", id)
	}
	return &row, nil
}

func (r *PaymentQueryRepositoryImpl) FindTinyByID(id int) (*entity.PaymentTiny, error) {
	var row entity.PaymentTiny
	if err := r.DB.First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "payment(id=%d)", id)
	}
	return &row, nil
}

func (r *PaymentQueryRepositoryImpl) FindPaymentCfReturnGiftByPaymentIDAndCfReturnGift(paymentID, cfReturnGiftID int) (*entity.PaymentCfReturnGift, error) {
	var row entity.PaymentCfReturnGift
	if err := r.DB.Where("payment_id = ? AND cf_return_gift_id = ?", paymentID, cfReturnGiftID).First(&row).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "payment_cf_return_gift(payment_id=%d,cf_return_gift_id=%d)")
	}
	return &row, nil
}
