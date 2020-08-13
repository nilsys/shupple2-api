package service

import (
	"context"

	"github.com/golang/mock/gomock"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/mock"
)

var _ = Describe("PaymentServiceImpl", func() {
	var (
		cmdSvc            *CfInnReserveRequestCommandServiceImpl
		paymentQueryRepo  repository.PaymentQueryRepository
		paymentCmdRepo    repository.PaymentCommandRepository
		innReserveCmdRepo repository.CfInnReserveRequestCommandRepository
		mailCmdRepo       repository.MailCommandRepository
		err               error
	)

	BeforeEach(func() {
		cmdSvc = tests.CfInnReserveRequestCommandServiceImpl
		paymentQueryRepo = tests.CfInnReserveRequestCommandServiceImpl.PaymentQueryRepository
		paymentCmdRepo = tests.CfInnReserveRequestCommandServiceImpl.PaymentCommandRepository
		innReserveCmdRepo = tests.CfInnReserveRequestCommandServiceImpl.CfInnReserveRequestCommandRepository
		mailCmdRepo = tests.CfInnReserveRequestCommandServiceImpl.MailCommandRepository
	})

	Describe("RequestReserve", func() {
		BeforeEach(func() {
			paymentQueryRepo.(*mock.MockPaymentQueryRepository).EXPECT().FindByID(context.Background(), paymentID).Return(newPayment(), nil)
			innReserveCmdRepo.(*mock.MockCfInnReserveRequestCommandRepository).EXPECT().Store(context.Background(), newCfInnReserveRequest()).Return(nil)
		})

		Context("正常系", func() {
			BeforeEach(func() {
				paymentQueryRepo.(*mock.MockPaymentQueryRepository).EXPECT().FindPaymentCfReturnGiftByPaymentIDAndCfReturnGift(paymentID, cfReturnGiftID).Return(newPaymentCfReturnGift(model.CfReturnGiftTypeReservedTicket, model.PaymentCfReturnGiftOtherTypeStatusUndefined, model.PaymentCfReturnGiftReservedTicketTypeStatusUnreserved, validDeadline), nil)
				paymentCmdRepo.(*mock.MockPaymentCommandRepository).EXPECT().MarkPaymentCfReturnGiftAsReserved(gomock.Any(), paymentID, cfReturnGiftID).Return(nil)
				mailCmdRepo.(*mock.MockMailCommandRepository).EXPECT().SendTemplateMail([]string{newUser(ownerUserID).Email}, newReserveRequestTemplateMailTemplate()).Return(nil)
			})

			It("エラーなし", func() {
				err = cmdSvc.RequestReserve(newUser(userID), paymentID, cfReturnGiftID, newCfReserveRequest())
				Expect(err).To(Succeed())
			})
		})

		Context("権限無し", func() {

			It("'forbidden'というエラーメッセージが返る", func() {
				err = cmdSvc.RequestReserve(newUser(ownerUserID), paymentID, cfReturnGiftID, newCfReserveRequest())
				Expect(err.Error()).To(Equal("forbidden"))
			})
		})

		Context("宿泊予約券では無い場合", func() {
			BeforeEach(func() {
				paymentCfReturnGift := newPaymentCfReturnGift(model.CfReturnGiftTypeOther, model.PaymentCfReturnGiftOtherTypeStatusOwnerUnconfirmed, model.PaymentCfReturnGiftReservedTicketTypeStatusUndefined, validDeadline)
				paymentQueryRepo.(*mock.MockPaymentQueryRepository).EXPECT().FindPaymentCfReturnGiftByPaymentIDAndCfReturnGift(paymentID, cfReturnGiftID).Return(paymentCfReturnGift, nil)
			})

			It("'not reserved ticket'というエラーメッセージが返る", func() {
				err = cmdSvc.RequestReserve(newUser(userID), paymentID, cfReturnGiftID, newCfReserveRequest())
				Expect(err.Error()).To(Equal("not reserved ticket"))
			})
		})

		Context("有効期限が切れている場合", func() {
			BeforeEach(func() {
				paymentCfReturnGift := newPaymentCfReturnGift(model.CfReturnGiftTypeReservedTicket, model.PaymentCfReturnGiftOtherTypeStatusUndefined, model.PaymentCfReturnGiftReservedTicketTypeStatusUnreserved, invalidDeadline)
				paymentQueryRepo.(*mock.MockPaymentQueryRepository).EXPECT().FindPaymentCfReturnGiftByPaymentIDAndCfReturnGift(paymentID, cfReturnGiftID).Return(paymentCfReturnGift, nil)
			})

			It("'expired'というエラーメッセージが返る", func() {
				err = cmdSvc.RequestReserve(newUser(userID), paymentID, cfReturnGiftID, newCfReserveRequest())
				Expect(err.Error()).To(Equal("expired"))
			})
		})

		Context("予約済な場合", func() {
			BeforeEach(func() {
				paymentCfReturnGift := newPaymentCfReturnGift(model.CfReturnGiftTypeReservedTicket, model.PaymentCfReturnGiftOtherTypeStatusUndefined, model.PaymentCfReturnGiftReservedTicketTypeStatusReserved, validDeadline)
				paymentQueryRepo.(*mock.MockPaymentQueryRepository).EXPECT().FindPaymentCfReturnGiftByPaymentIDAndCfReturnGift(paymentID, cfReturnGiftID).Return(paymentCfReturnGift, nil)
			})

			It("予約済な場合", func() {
				err = cmdSvc.RequestReserve(newUser(userID), paymentID, cfReturnGiftID, newCfReserveRequest())
				Expect(err.Error()).To(Equal("can't transit"))
			})
		})

	})
})
