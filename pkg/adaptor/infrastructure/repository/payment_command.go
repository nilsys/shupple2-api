package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"

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

func (r *PaymentCommandRepositoryImpl) Store(c context.Context, payment *entity.PaymentTiny) error {
	if err := r.DB(c).Save(payment).Error; err != nil {
		return errors.Wrap(err, "failed store payment")
	}
	return nil
}

func (r *PaymentCommandRepositoryImpl) StorePaymentReturnGiftList(c context.Context, list []*entity.PaymentCfReturnGiftTiny, paymentID int) error {
	for _, gift := range list {
		gift.PaymentID = paymentID
		if err := r.DB(c).Save(gift).Error; err != nil {
			return errors.Wrap(err, "failed store payment_return_gift")
		}
	}
	return nil
}

func (r *PaymentCommandRepositoryImpl) MarkPaymentCfReturnGiftAsReserved(c context.Context, paymentID, cfReturnGiftID int) error {
	if err := r.DB(c).Exec("UPDATE payment_cf_return_gift SET gift_type_reserved_ticket_status = ? WHERE payment_id = ? AND cf_return_gift_id = ?", model.PaymentCfReturnGiftReservedTicketTypeStatusReserved, paymentID, cfReturnGiftID).Error; err != nil {
		return errors.Wrap(err, "failed update payment_cf_return_gift.gift_type_reserved_ticket_status")
	}
	return nil
}
