package service

import (
	"context"
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	PaymentCommandService interface {
		ReservePaymentCfReturnGift(user *entity.User, paymentID, cfReturnGiftID int, request *entity.CfReserveRequest) error
	}

	PaymentCommandServiceImpl struct {
		repository.PaymentCommandRepository
		repository.PaymentQueryRepository
		repository.MailCommandRepository
		TransactionService
	}
)

var PaymentCommandServiceSet = wire.NewSet(
	wire.Struct(new(PaymentCommandServiceImpl), "*"),
	wire.Bind(new(PaymentCommandService), new(*PaymentCommandServiceImpl)),
)

func (s *PaymentCommandServiceImpl) ReservePaymentCfReturnGift(user *entity.User, paymentID, cfReturnGiftID int, request *entity.CfReserveRequest) error {
	payment, err := s.PaymentQueryRepository.FindByID(paymentID)
	if err != nil {
		return errors.Wrap(err, "failed find payment")
	}
	if !user.IsSelfID(payment.UserID) {
		return serror.New(nil, serror.CodeForbidden, "forbidden")
	}

	paymentCfReturnGift, err := s.PaymentQueryRepository.FindPaymentCfReturnGiftByPaymentIDAndCfReturnGift(paymentID, cfReturnGiftID)
	if err != nil {
		return errors.Wrap(err, "failed find payment_cf_return_gift")
	}

	if paymentCfReturnGift.CfReturnGift.GiftType != model.CfReturnGiftTypeReservedTicket {
		return serror.New(nil, serror.CodeInvalidParam, "not reserved ticket")
	}
	if paymentCfReturnGift.CfReturnGiftSnapshot.Deadline.After(time.Now()) {
		return serror.New(nil, serror.CodeInvalidParam, "expired")
	}
	if !paymentCfReturnGift.GiftTypeReservedTicketStatus.CanTransit(model.PaymentCfReturnGiftReservedTicketTypeStatusReserved) {
		return serror.New(nil, serror.CodeInvalidParam, "can't transit")
	}

	return s.TransactionService.Do(func(ctx context.Context) error {
		if err := s.PaymentCommandRepository.MarkPaymentCfReturnGiftAsReserved(ctx, paymentID, cfReturnGiftID); err != nil {
			return errors.Wrap(err, "failed mark as reserved")
		}

		if err := s.MailCommandRepository.SendTemplateMail(payment.Owner.Email, entity.NewReserveRequestTemplateFromCfReserveRequest(request, payment.ChargeID, paymentCfReturnGift.CfReturnGiftSnapshot.Body)); err != nil {
			return errors.Wrap(err, "failed send email from ses")
		}

		return nil
	})
}
