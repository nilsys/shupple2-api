package service

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	payjp2 "github.com/payjp/payjp-go/v1"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/command"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository/payjp"
	"github.com/stayway-corp/stayway-media-api/pkg/mock"
	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

var _ = Describe("ChargeServiceImpl", func() {
	var (
		cmdSvc                *ChargeCommandServiceImpl
		paymentQueryRepo      repository.PaymentQueryRepository
		paymentCmdRepo        repository.PaymentCommandRepository
		mailCmdRepo           repository.MailCommandRepository
		cardQueryRepo         repository.CardQueryRepository
		chargeCmdRepo         payjp.ChargeCommandRepository
		cfReturnGiftQueryRepo repository.CfReturnGiftQueryRepository
		userQueryRepo         repository.UserQueryRepository
		cfReturnGiftCmdRepo   repository.CfReturnGiftCommandRepository
		shippingQueryRepo     repository.ShippingQueryRepository
		cfProjectCmdRepo      repository.CfProjectCommandRepository
		err                   error
	)

	BeforeEach(func() {
		cmdSvc = tests.ChargeCommandServiceImpl
		paymentQueryRepo = tests.ChargeCommandServiceImpl.PaymentQueryRepository
		paymentCmdRepo = tests.ChargeCommandServiceImpl.PaymentCommandRepository
		mailCmdRepo = tests.ChargeCommandServiceImpl.MailCommandRepository
		cardQueryRepo = tests.ChargeCommandServiceImpl.CardQueryRepository
		chargeCmdRepo = tests.ChargeCommandServiceImpl.ChargeCommandRepository
		cfReturnGiftQueryRepo = tests.ChargeCommandServiceImpl.CfReturnGiftQueryRepository
		userQueryRepo = tests.ChargeCommandServiceImpl.UserQueryRepository
		cfReturnGiftCmdRepo = tests.ChargeCommandServiceImpl.CfReturnGiftCommandRepository
		shippingQueryRepo = tests.ChargeCommandServiceImpl.ShippingQueryRepository
		cfProjectCmdRepo = tests.ChargeCommandServiceImpl.CfProjectCommandRepository
	})

	Describe("Create", func() {
		BeforeEach(func() {
			shippingQueryRepo.(*mock.MockShippingQueryRepository).EXPECT().FindLatestShippingAddressByUserID(gomock.Any(), userID).Return(newShippingAddress(), nil)
		})

		Context("正常系", func() {
			BeforeEach(func() {
				cfReturnGiftCmdRepo.(*mock.MockCfReturnGiftCommandRepository).EXPECT().LockByIDs(gomock.Any(), []int{cfReturnGiftID}).Return(newCfReturnGiftList(), nil)
				cfProjectCmdRepo.(*mock.MockCfProjectCommandRepository).EXPECT().Lock(gomock.Any(), cfProjectID).Return(newCfProject(), nil)
				userQueryRepo.(*mock.MockUserQueryRepository).EXPECT().FindByID(ownerUserID).Return(newUser(ownerUserID), nil)
				cfReturnGiftQueryRepo.(*mock.MockCfReturnGiftQueryRepository).EXPECT().FindSoldCountByReturnGiftIDs(gomock.Any(), []int{cfReturnGiftID}).Return(newCfReturnGiftSoldCountList(validSoldCount), nil)
				cardQueryRepo.(*mock.MockCardQueryRepository).EXPECT().FindLatestByUserID(gomock.Any(), userID).Return(newCard(), nil)
				chargeCmdRepo.(*mock.MockChargeCommandRepository).EXPECT().Create(newUser(userID).PayjpCustomerID(), newCard().CardID, cfReturnGiftPrice+cmdSvc.CfProjectConfig.SystemFee).Return(newChargeRes(true), nil)
				paymentCmdRepo.(*mock.MockPaymentCommandRepository).EXPECT().Store(gomock.Any(), newPaymentTinyForChargeTest()).Return(nil)
				paymentCmdRepo.(*mock.MockPaymentCommandRepository).EXPECT().StorePaymentReturnGiftList(gomock.Any(), []*entity.PaymentCfReturnGiftTiny{newPaymentCfReturnGiftTinyForOtherForChargeTest(model.PaymentCfReturnGiftOtherTypeStatusOwnerUnconfirmed)}, paymentIDForChargeTest).Return(nil)
				cfProjectCmdRepo.(*mock.MockCfProjectCommandRepository).EXPECT().StoreSupportComment(gomock.Any(), newCfProjectSupportCmt()).Return(nil)
				cfProjectCmdRepo.(*mock.MockCfProjectCommandRepository).EXPECT().IncrementSupportCommentCount(gomock.Any(), cfProjectID).Return(nil)
				cfProjectCmdRepo.(*mock.MockCfProjectCommandRepository).EXPECT().IncrementAchievedPrice(gomock.Any(), cfProjectID, cfReturnGiftPrice).Return(nil)
				paymentQueryRepo.(*mock.MockPaymentQueryRepository).EXPECT().FindByID(gomock.Any(), paymentIDForChargeTest).Return(newPayment(), nil)
				chargeCmdRepo.(*mock.MockChargeCommandRepository).EXPECT().Capture(newChargeRes(true).ID).Return(nil)
				gomock.InOrder(
					mailCmdRepo.(*mock.MockMailCommandRepository).EXPECT().SendTemplateMail([]string{newShippingAddress().Email}, newThanksPurchaseTemplate()).Return(nil),
					mailCmdRepo.(*mock.MockMailCommandRepository).EXPECT().SendTemplateMail([]string{newUser(ownerUserID).Email}, newThanksPurchaseForOwnerTemplate()).Return(nil),
				)
			})

			It("エラーなし", func() {
				_, err = cmdSvc.Create(newUser(userID), newPaymentListCmd())
				Expect(err).To(Succeed())
			})
		})
	})
})

func newCard() *entity.Card {
	card := &entity.Card{
		ID:     cardID,
		UserID: userID,
	}
	util.FillDummyString(card, cardID)
	return card
}

func newCfReturnGiftList() *entity.CfReturnGiftList {
	return &entity.CfReturnGiftList{
		List: []*entity.CfReturnGift{{
			CfReturnGiftTiny: *newCfReturnGiftTiny(model.CfReturnGiftTypeOther),
			Snapshot:         newCfReturnGiftSnapshot(validDeadline),
		},
		},
	}
}

func newCfReturnGiftSoldCountList(soldCnt int) *entity.CfReturnGiftSoldCountList {
	return &entity.CfReturnGiftSoldCountList{
		List: []*entity.CfReturnGiftSoldCount{newCfReturnGiftSoldCount(soldCnt)},
	}
}

func newCfReturnGiftSoldCount(soldCnt int) *entity.CfReturnGiftSoldCount {
	return &entity.CfReturnGiftSoldCount{
		CfReturnGiftID: cfReturnGiftID,
		SoldCount:      soldCnt,
	}
}

func newShippingAddress() *entity.ShippingAddress {
	address := entity.ShippingAddress{
		ID:     shippingAddressID,
		UserID: userID,
	}
	util.FillDummyString(&address, shippingAddressID)
	return &address
}

func newChargeRes(paid bool) *payjp2.ChargeResponse {
	return &payjp2.ChargeResponse{
		ID:   chargeResID,
		Paid: paid,
	}
}

func newCfProjectSupportCmt() *entity.CfProjectSupportCommentTiny {
	cmt := &entity.CfProjectSupportCommentTiny{
		UserID:      userID,
		CfProjectID: cfProjectID,
		Body:        cfProjectSupportCommentBody,
	}
	return cmt
}

func newPaymentListCmd() *command.CreateCharge {
	return &command.CreateCharge{
		List:               []*command.Charge{{ReturnGiftID: cfReturnGiftID, ReturnGiftSnapshotID: cfReturnGiftSnapshotID, Amount: paymentCfReturnGiftAmount}},
		SupportCommentBody: cfProjectSupportCommentBody,
	}
}
