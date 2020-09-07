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

// GiftTypeOtherStatusがNULLの場合は更新できない
func (r *PaymentCommandRepositoryImpl) MarkPaymentCfReturnGiftAsCancel(c context.Context, paymentID, cfReturnGiftID int) error {
	if err := r.DB(c).Exec("UPDATE payment_cf_return_gift SET gift_type_other_status = ? WHERE payment_id = ? AND cf_return_gift_id = ? AND gift_type_other_status IS NOT NULL",
		model.PaymentCfReturnGiftOtherTypeStatusCanceled, paymentID, cfReturnGiftID).Error; err != nil {
		return errors.Wrap(err, "failed update payment_cf_return_gift.is_canceled")
	}
	return nil
}

// GiftTypeReservedTicketStatusがNULLの場合は更新できない
func (r *PaymentCommandRepositoryImpl) MarkPaymentCfReturnGiftAsReserved(c context.Context, paymentID, cfReturnGiftID int) error {
	if err := r.DB(c).Exec("UPDATE payment_cf_return_gift SET gift_type_reserved_ticket_status = ?, user_reserve_requested_at = NOW() WHERE payment_id = ? AND cf_return_gift_id = ? AND gift_type_reserved_ticket_status IS NOT NULL",
		model.PaymentCfReturnGiftReservedTicketTypeStatusReserved, paymentID, cfReturnGiftID).Error; err != nil {
		return errors.Wrap(err, "failed update payment_cf_return_gift.gift_type_reserved_ticket_status")
	}
	return nil
}

// 有効期限切れのPaymentCfReturnGiftを全て期限切れステータスへ
func (r *PaymentCommandRepositoryImpl) MarkExpiredAllPaymentCfReturnGiftAsExpired() error {
	if err := r.DB(context.Background()).
		Exec(`UPDATE payment_cf_return_gift SET gift_type_reserved_ticket_status = ? WHERE gift_type_reserved_ticket_status IS NOT NULL AND (payment_id, cf_return_gift_id) IN
			(SELECT * FROM (SELECT payment_cf_return_gift.payment_id, payment_cf_return_gift.cf_return_gift_id FROM payment_cf_return_gift INNER JOIN cf_return_gift_snapshot ON payment_cf_return_gift.cf_return_gift_snapshot_id = cf_return_gift_snapshot.id AND cf_return_gift_snapshot.deadline < NOW()) t)`, model.PaymentCfReturnGiftReservedTicketTypeStatusExpired).
		Error; err != nil {
		return errors.Wrap(err, "failed update payment_cf_return_git.gift_type_reserved_ticket_status")
	}

	return nil
}
