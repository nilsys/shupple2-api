package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type PaymentCommandRepositoryImpl struct {
	DAO
}

var PaymentCommandRepositorySet = wire.NewSet(
	wire.Struct(new(PaymentCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.PaymentCommandRepository), new(*PaymentCommandRepositoryImpl)),
)

func (r *PaymentCommandRepositoryImpl) Store(c context.Context, payment *entity.Payment) error {
	if err := r.DB(c).Save(payment).Error; err != nil {
		return errors.Wrap(err, "failed store payment")
	}
	return nil
}

func (r *PaymentCommandRepositoryImpl) StorePaymentReturnGiftList(c context.Context, list []*entity.PaymentCfReturnGift, paymentID int) error {
	for _, gift := range list {
		gift.PaymentID = paymentID
		if err := r.DB(c).Save(gift).Error; err != nil {
			return errors.Wrap(err, "failed store payment_return_gift")
		}
	}
	return nil
}
