package service

import (
	"context"
	"time"

	"github.com/google/wire"

	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	CfInnReserveRequestCommandService interface {
		RequestReserve(user *entity.User, paymentID, cfReturnGiftID int, request *entity.CfInnReserveRequest) error
	}

	CfInnReserveRequestCommandServiceImpl struct {
		repository.CfInnReserveRequestCommandRepository
		repository.PaymentQueryRepository
		repository.PaymentCommandRepository
		repository.MailCommandRepository
		TransactionService
	}
)

var CfInnReserveRequestCommandServiceSet = wire.NewSet(
	wire.Struct(new(CfInnReserveRequestCommandServiceImpl), "*"),
	wire.Bind(new(CfInnReserveRequestCommandService), new(*CfInnReserveRequestCommandServiceImpl)),
)

func (s *CfInnReserveRequestCommandServiceImpl) RequestReserve(user *entity.User, paymentID, cfReturnGiftID int, request *entity.CfInnReserveRequest) error {
	payment, err := s.PaymentQueryRepository.FindByID(context.Background(), paymentID)
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

	if deadline := paymentCfReturnGift.CfReturnGiftSnapshot.Deadline; deadline.Valid && deadline.Time.Before(time.Now()) {
		return serror.New(nil, serror.CodeInvalidParam, "expired")
	}
	if !paymentCfReturnGift.ResolveGiftTypeReservedTicketStatus().CanTransit(model.PaymentCfReturnGiftReservedTicketTypeStatusReserved) {
		return serror.New(nil, serror.CodeInvalidParam, "can't transit")
	}

	return s.TransactionService.Do(func(ctx context.Context) error {
		if err := s.PaymentCommandRepository.MarkPaymentCfReturnGiftAsReserved(ctx, paymentID, cfReturnGiftID); err != nil {
			return errors.Wrap(err, "failed mark as reserved")
		}

		if err := s.CfInnReserveRequestCommandRepository.Store(ctx, request); err != nil {
			return errors.Wrap(err, "failed store cf_reserve_request")
		}

		if err := s.MailCommandRepository.SendTemplateMail([]string{payment.Owner.Email}, entity.NewReserveRequestTemplateFromCfReserveRequest(request, payment.ChargeID, paymentCfReturnGift.CfReturnGiftSnapshot.Title)); err != nil {
			return errors.Wrap(err, "failed send email from ses")
		}

		return nil
	})
}
