package service

import (
	"context"
	"strconv"

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
		repository.CfInnReserveRequestQueryRepository
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

	isReserved, err := s.CfInnReserveRequestQueryRepository.IsExistByPaymentIDAndCfReturnGiftID(paymentID, cfReturnGiftID)
	if err != nil {
		return errors.Wrap(err, "failed find reserve_request")
	}
	if isReserved {
		return serror.New(nil, serror.CodeInvalidParam, "already requested")
	}

	paymentCfReturnGift, err := s.PaymentQueryRepository.FindPaymentCfReturnGiftByPaymentIDAndCfReturnGift(paymentID, cfReturnGiftID)
	if err != nil {
		return errors.Wrap(err, "failed find payment_cf_return_gift")
	}

	if paymentCfReturnGift.CfReturnGift.GiftType != model.CfReturnGiftTypeReservedTicket {
		return serror.New(nil, serror.CodeInvalidParam, "not reserved ticket")
	}

	if paymentCfReturnGift.CfReturnGiftSnapshot.IsExpired() {
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

		// プロジェクトオーナーへのメールテンプレート
		toOwnerEmail := entity.NewReserveRequestForOwnerTemplate(request.FullNameMailFmt(), request.FullNameKanaMailFmt(), request.Email, request.PhoneNumber, payment.ChargeID, paymentCfReturnGift.InquiryCode,
			paymentCfReturnGift.CfReturnGiftSnapshot.Title, paymentCfReturnGift.CfReturnGiftSnapshot.Body, model.TimeFront(request.CheckinAt).ToString(), model.TimeFront(request.CheckoutAt).ToString(),
			strconv.Itoa(request.StayDays), strconv.Itoa(request.AdultMemberCount), strconv.Itoa(request.ChildMemberCount), request.Remark,
		)

		// ユーザーへのメールテンプレート
		toUserEmail := entity.NewReserveRequestForUserTemplate(request.FullNameMailFmt(), request.FullNameKanaMailFmt(), request.Email, request.PhoneNumber, payment.ChargeID, paymentCfReturnGift.InquiryCode,
			paymentCfReturnGift.CfReturnGiftSnapshot.Title, paymentCfReturnGift.CfReturnGiftSnapshot.Body, model.TimeFront(request.CheckinAt).ToString(), model.TimeFront(request.CheckoutAt).ToString(),
			strconv.Itoa(request.StayDays), strconv.Itoa(request.AdultMemberCount), strconv.Itoa(request.ChildMemberCount), request.Remark,
		)

		if err := s.MailCommandRepository.SendTemplateMail([]string{payment.Owner.Email}, toOwnerEmail); err != nil {
			return errors.Wrap(err, "failed send email from ses")
		}

		if err := s.MailCommandRepository.SendTemplateMail([]string{request.Email}, toUserEmail); err != nil {
			return errors.Wrap(err, "failed send email from ses")
		}

		return nil
	})
}
