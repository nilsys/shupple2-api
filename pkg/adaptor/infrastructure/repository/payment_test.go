package repository

import (
	"context"
	"time"

	"gopkg.in/guregu/null.v3"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"

	"github.com/stayway-corp/stayway-media-api/pkg/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

var _ = Describe("PaymentRepositoryImpl", func() {
	var (
		queryRepo   *PaymentQueryRepositoryImpl
		commandRepo *PaymentCommandRepositoryImpl
	)
	BeforeEach(func() {
		queryRepo = tests.PaymentQueryRepositoryImpl
		commandRepo = tests.PaymentCommandRepositoryImpl
		truncate(db)
		Expect(db.Save(newUser(userID)).Error).To(Succeed())
		Expect(db.Save(newShippingAddress(shippingAddressID, userID)).Error).To(Succeed())
		Expect(db.Save(newCard()).Error).To(Succeed())
		Expect(db.Save(newCfProjectTiny(cfProjectID, userID)).Error).To(Succeed())
		Expect(db.Save(newCfReturnGiftTiny()).Error).To(Succeed())
		Expect(db.Save(newCfReturnGiftSnapshot()).Error).To(Succeed())
		Expect(db.Save(newCfProjectSnapshotTiny()).Error).To(Succeed())
		Expect(db.Save(newPayment()).Error).To(Succeed())
		Expect(db.Save(newPaymentCfReturnGift()).Error).To(Succeed())
	})

	base := newPayment()

	Describe("Store", func() {
		It("保存する", func() {
			Expect(commandRepo.Store(context.Background(), base))
			actual, err := queryRepo.FindTinyByID(paymentID)
			Expect(err).To(Succeed())

			Expect(actual).To(entity.EqualEntity(base))
		})
	})

	Describe("MarkPaymentCfReturnGiftAsCancel", func() {
		It("gif_type_other_statusをcanceledへupdate", func() {
			expect, err := queryRepo.FindPaymentCfReturnGiftByPaymentIDAndCfReturnGift(paymentID, cfReturnGiftID)
			Expect(err).To(Succeed())

			Expect(commandRepo.MarkPaymentCfReturnGiftAsCancel(context.Background(), paymentID, cfReturnGiftID))

			actual, err := queryRepo.FindPaymentCfReturnGiftByPaymentIDAndCfReturnGift(paymentID, cfReturnGiftID)
			Expect(err).To(Succeed())

			expect.GiftTypeOtherStatus = null.IntFrom(int64(model.PaymentCfReturnGiftOtherTypeStatusCanceled))

			Expect(actual).To(entity.EqualEntity(expect))
		})
	})

	Describe("MarkPaymentCfReturnGiftAsReserved", func() {
		It("gift_type_reserved_ticket_statusをreservedへupdate", func() {
			expect, err := queryRepo.FindPaymentCfReturnGiftByPaymentIDAndCfReturnGift(paymentID, cfReturnGiftID)
			Expect(err).To(Succeed())

			Expect(commandRepo.MarkPaymentCfReturnGiftAsReserved(context.Background(), paymentID, cfReturnGiftID))

			actual, err := queryRepo.FindPaymentCfReturnGiftByPaymentIDAndCfReturnGift(paymentID, cfReturnGiftID)
			Expect(err).To(Succeed())

			expect.GiftTypeReservedTicketStatus = null.IntFrom(int64(model.PaymentCfReturnGiftReservedTicketTypeStatusReserved))

			Expect(actual.UserReserveRequestedAt).NotTo(Equal(BeZero()))

			actual.UserReserveRequestedAt = &time.Time{}
			expect.UserReserveRequestedAt = &time.Time{}
			Expect(actual).To(entity.EqualEntity(expect))
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

func newPayment() *entity.PaymentTiny {
	tmp := time.Date(2020, 7, 7, 0, 0, 0, 0, time.Local)
	payment := &entity.PaymentTiny{
		ID:                      paymentID,
		UserID:                  userID,
		ProjectOwnerID:          userID,
		CardID:                  cardID,
		ShippingAddressID:       shippingAddressID,
		OwnerDepositRequestedAt: &tmp,
	}
	util.FillDummyString(payment, paymentID)
	return payment
}

func newCfReturnGiftTiny() *entity.CfReturnGiftTiny {
	gift := &entity.CfReturnGiftTiny{
		ID:          cfReturnGiftID,
		CfProjectID: cfProjectID,
	}
	util.FillDummyString(gift, cfReturnGiftID)
	return gift
}

func newCfReturnGiftSnapshot() *entity.CfReturnGiftSnapshotTiny {
	tiny := &entity.CfReturnGiftSnapshotTiny{
		SnapshotID:     cfReturnGiftSnapshotID,
		CfReturnGiftID: cfReturnGiftID,
		Deadline:       null.TimeFrom(time.Date(2020, 7, 7, 0, 0, 0, 0, time.Local)),
	}
	util.FillDummyString(tiny, cfReturnGiftSnapshotID)
	return tiny
}

func newPaymentCfReturnGift() *entity.PaymentCfReturnGiftTiny {
	tmp := time.Date(2020, 7, 7, 0, 0, 0, 0, time.Local)
	return &entity.PaymentCfReturnGiftTiny{
		PaymentID:                    paymentID,
		CfReturnGiftID:               cfReturnGiftID,
		CfReturnGiftSnapshotID:       cfReturnGiftSnapshotID,
		CfProjectID:                  cfProjectID,
		CfProjectSnapshotID:          cfProjectSnapshotID,
		GiftTypeOtherStatus:          null.IntFrom(0),
		GiftTypeReservedTicketStatus: null.IntFrom(0),
		OwnerConfirmedAt:             &tmp,
		UserReserveRequestedAt:       &tmp,
	}
}

func newCfProjectTiny(id, userID int) entity.CfProjectTiny {
	return entity.CfProjectTiny{
		ID:     id,
		UserID: userID,
	}
}
