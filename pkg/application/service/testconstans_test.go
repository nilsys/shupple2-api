package service

import (
	"time"

	"gopkg.in/guregu/null.v3"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

const (
	userID                   = 101
	cfReserveRequestDummyInt = 111
	paymentID                = 121
	cfReturnGiftID           = 131
	cfReturnGiftSnapshotID   = 141
	cfProjectID              = 151
	cfProjectSnapshotID      = 161
	ownerUserID              = 171
)

var (
	validDeadline   = time.Date(9999, 9, 9, 0, 0, 0, 0, time.Local)
	invalidDeadline = time.Date(1999, 9, 9, 0, 0, 0, 0, time.Local)
)

func newUser(userID int) *entity.User {
	user := &entity.User{
		ID:            userID,
		Birthdate:     time.Date(2020, 7, 7, 0, 0, 0, 0, time.Local),
		UserInterests: []*entity.UserInterest{},
		Gender:        model.GenderMale,
	}
	util.FillDummyString(user, userID)
	return user
}

func newCfReserveRequest() *entity.CfReserveRequest {
	req := &entity.CfReserveRequest{}
	util.FillDummyString(req, cfReserveRequestDummyInt)
	return req
}

func newPayment() *entity.Payment {
	return &entity.Payment{
		PaymentTiny: *newPaymentTiny(),
		Owner:       newUser(ownerUserID),
	}
}

func newPaymentTiny() *entity.PaymentTiny {
	tiny := &entity.PaymentTiny{
		ID:     paymentID,
		UserID: userID,
	}
	util.FillDummyString(tiny, paymentID)
	return tiny
}

func newPaymentCfReturnGift(giftType model.CfReturnGiftType, otherStatus model.PaymentCfReturnGiftOtherTypeStatus, ticketStatus model.PaymentCfReturnGiftReservedTicketTypeStatus, deadline time.Time) *entity.PaymentCfReturnGift {
	var tiny *entity.PaymentCfReturnGiftTiny
	switch giftType {
	case model.CfReturnGiftTypeReservedTicket:
		tiny = newPaymentCfReturnGiftTinyForReservedTicket(ticketStatus)
	default:
		tiny = newPaymentCfReturnGiftTinyForOther(otherStatus)
	}
	return &entity.PaymentCfReturnGift{
		PaymentCfReturnGiftTiny: *tiny,
		CfReturnGift:            newCfReturnGiftTiny(giftType),
		CfReturnGiftSnapshot:    newCfReturnGiftSnapshot(deadline),
	}
}

func newPaymentCfReturnGiftTinyForOther(otherStatus model.PaymentCfReturnGiftOtherTypeStatus) *entity.PaymentCfReturnGiftTiny {
	return &entity.PaymentCfReturnGiftTiny{
		PaymentID:              paymentID,
		CfReturnGiftID:         cfReturnGiftID,
		CfReturnGiftSnapshotID: cfReturnGiftSnapshotID,
		CfProjectID:            cfProjectID,
		CfProjectSnapshotID:    cfProjectSnapshotID,
		GiftTypeOtherStatus:    null.IntFrom(int64(otherStatus)),
	}
}

func newPaymentCfReturnGiftTinyForReservedTicket(ticketStatus model.PaymentCfReturnGiftReservedTicketTypeStatus) *entity.PaymentCfReturnGiftTiny {
	return &entity.PaymentCfReturnGiftTiny{
		PaymentID:                    paymentID,
		CfReturnGiftID:               cfReturnGiftID,
		CfReturnGiftSnapshotID:       cfReturnGiftSnapshotID,
		CfProjectID:                  cfProjectID,
		CfProjectSnapshotID:          cfProjectSnapshotID,
		GiftTypeReservedTicketStatus: null.IntFrom(int64(ticketStatus)),
	}
}

func newCfReturnGiftTiny(giftType model.CfReturnGiftType) *entity.CfReturnGiftTiny {
	tiny := &entity.CfReturnGiftTiny{
		ID:               cfReturnGiftID,
		CfProjectID:      cfProjectID,
		LatestSnapshotID: null.IntFrom(int64(cfReturnGiftSnapshotID)),
		GiftType:         giftType,
	}
	util.FillDummyString(tiny, cfReturnGiftID)
	return tiny
}

func newCfReturnGiftSnapshot(deadline time.Time) *entity.CfReturnGiftSnapshotTiny {
	tiny := &entity.CfReturnGiftSnapshotTiny{
		SnapshotID:     cfReturnGiftSnapshotID,
		CfReturnGiftID: cfReturnGiftID,
		Deadline:       null.TimeFrom(deadline),
	}
	util.FillDummyString(tiny, cfReturnGiftSnapshotID)
	return tiny
}

func newMailTemplate() entity.MailTemplate {
	return entity.NewReserveRequestTemplateFromCfReserveRequest(newCfReserveRequest(), newPayment().ChargeID, newPaymentCfReturnGift(model.CfReturnGiftTypeReservedTicket, model.PaymentCfReturnGiftOtherTypeStatusUndefined, model.PaymentCfReturnGiftReservedTicketTypeStatusUnreserved, validDeadline).CfReturnGiftSnapshot.Body)
}
