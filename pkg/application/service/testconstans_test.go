package service

import (
	"strconv"
	"time"

	"gopkg.in/guregu/null.v3"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

const (
	userID                         = 101
	cfReserveRequestDummyInt       = 111
	paymentID                      = 121
	cfReturnGiftID                 = 131
	cfReturnGiftSnapshotID         = 141
	cfProjectID                    = 151
	cfProjectSnapshotID            = 161
	ownerUserID                    = 171
	cardID                         = 181
	shippingAddressID              = 191
	chargeResID                    = "201"
	cfReturnGiftPrice              = 1000
	cfProjectSupportCommentBody    = "dummy"
	cfReturnGiftFullAmount         = 100
	systemFee                      = 220
	paymentIDForChargeTest         = 0
	paymentCfReturnGiftAmount      = 1
	paymentCfReturnGiftInquiryCode = "dummy"
)

var (
	validDeadline   = time.Date(9999, 9, 9, 0, 0, 0, 0, time.Local)
	invalidDeadline = time.Date(1999, 9, 9, 0, 0, 0, 0, time.Local)
	validSoldCount  = 0
)

func newUser(userID int) *entity.User {
	user := &entity.User{
		UserTiny: entity.UserTiny{
			ID:        userID,
			Birthdate: time.Date(2020, 7, 7, 0, 0, 0, 0, time.Local),
			Gender:    model.GenderMale,
		},
		UserInterests: []*entity.UserInterest{},
	}
	util.FillDummyString(user, userID)
	return user
}

func newCfReserveRequest() *entity.CfInnReserveRequest {
	req := &entity.CfInnReserveRequest{}
	util.FillDummyString(req, cfReserveRequestDummyInt)
	return req
}

func newPayment() *entity.Payment {
	return &entity.Payment{
		PaymentTiny:         *newPaymentTiny(),
		Owner:               newUser(ownerUserID),
		PaymentCfReturnGift: []*entity.PaymentCfReturnGift{newPaymentCfReturnGift(model.CfReturnGiftTypeOther, model.PaymentCfReturnGiftOtherTypeStatusOwnerUnconfirmed, model.PaymentCfReturnGiftReservedTicketTypeStatusUndefined, validDeadline)},
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
		InquiryCode:            paymentCfReturnGiftInquiryCode,
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
		InquiryCode:                  paymentCfReturnGiftInquiryCode,
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
		Price:          cfReturnGiftPrice,
		FullAmount:     cfReturnGiftFullAmount,
	}
	util.FillDummyString(tiny, cfReturnGiftSnapshotID)
	return tiny
}

func newReserveRequestForOwnerTemplate() entity.MailTemplate {
	req := newCfInnReserveRequest()
	pCfReturnGift := newPaymentCfReturnGift(model.CfReturnGiftTypeReservedTicket, model.PaymentCfReturnGiftOtherTypeStatusUndefined, model.PaymentCfReturnGiftReservedTicketTypeStatusUnreserved, validDeadline)
	return entity.NewReserveRequestForOwnerTemplate(req.FullNameMailFmt(), req.FullNameKanaMailFmt(), req.Email, req.PhoneNumber, newPayment().ChargeID, paymentCfReturnGiftInquiryCode, pCfReturnGift.CfReturnGiftSnapshot.Title,
		pCfReturnGift.CfReturnGiftSnapshot.Body, model.TimeFront(req.CheckinAt).ToString(), model.TimeFront(req.CheckoutAt).ToString(), strconv.Itoa(req.StayDays),
		strconv.Itoa(req.AdultMemberCount), strconv.Itoa(req.ChildMemberCount), req.Remark,
	)
}

func newReserveRequestForUserMailTemplate() entity.MailTemplate {
	req := newCfInnReserveRequest()
	pCfReturnGift := newPaymentCfReturnGift(model.CfReturnGiftTypeReservedTicket, model.PaymentCfReturnGiftOtherTypeStatusUndefined, model.PaymentCfReturnGiftReservedTicketTypeStatusUnreserved, validDeadline)
	return entity.NewReserveRequestForUserTemplate(req.FullNameMailFmt(), req.FullNameKanaMailFmt(), req.Email, req.PhoneNumber, newPayment().ChargeID, paymentCfReturnGiftInquiryCode, pCfReturnGift.CfReturnGiftSnapshot.Title,
		pCfReturnGift.CfReturnGiftSnapshot.Body, model.TimeFront(req.CheckinAt).ToString(), model.TimeFront(req.CheckoutAt).ToString(), strconv.Itoa(req.StayDays),
		strconv.Itoa(req.AdultMemberCount), strconv.Itoa(req.ChildMemberCount), req.Remark,
	)
}

func newThanksPurchaseTemplate() entity.MailTemplate {
	address := newShippingAddress()
	gift := newCfReturnGiftSnapshot(validDeadline)
	return entity.NewThanksPurchaseTemplate(newUser(ownerUserID).Name, "<br>"+gift.Title+"<br>お問い合わせ番号: "+paymentCfReturnGiftInquiryCode+"<br>有効期限: "+model.TimeFront(gift.Deadline.Time).ToString()+"<br>", chargeResID, util.WithComma(systemFee), util.WithComma(cfReturnGiftPrice+systemFee), address.Email, address.FullAddress(), address.PhoneNumber, address.FullName())
}

func newCfProjectTiny() entity.CfProjectTiny {
	project := entity.CfProjectTiny{
		ID:     cfProjectID,
		UserID: ownerUserID,
	}
	util.FillDummyString(&project, cfProjectID)
	return project
}

func newCfProjectSnapshot() entity.CfProjectSnapshot {
	snapshot := entity.CfProjectSnapshot{
		CfProjectSnapshotTiny: entity.CfProjectSnapshotTiny{
			SnapshotID:  cfProjectSnapshotID,
			CfProjectID: cfProjectID,
			UserID:      ownerUserID,
		},
	}
	util.FillDummyString(&snapshot, cfProjectSnapshotID)
	return snapshot
}

func newCfProject() *entity.CfProject {
	return &entity.CfProject{
		CfProjectTiny: newCfProjectTiny(),
		Snapshot:      newCfProjectSnapshot(),
	}
}

func newPaymentTinyForChargeTest() *entity.PaymentTiny {
	return &entity.PaymentTiny{
		UserID:            userID,
		ProjectOwnerID:    ownerUserID,
		CardID:            cardID,
		ChargeID:          chargeResID,
		ShippingAddressID: shippingAddressID,
		TotalPrice:        cfReturnGiftPrice,
		CommissionPrice:   systemFee,
	}
}

func newPaymentCfReturnGiftTinyForOtherForChargeTest(otherStatus model.PaymentCfReturnGiftOtherTypeStatus) *entity.PaymentCfReturnGiftTiny {
	return &entity.PaymentCfReturnGiftTiny{
		CfReturnGiftID:         cfReturnGiftID,
		CfReturnGiftSnapshotID: cfReturnGiftSnapshotID,
		CfProjectID:            cfProjectID,
		Amount:                 paymentCfReturnGiftAmount,
		InquiryCode:            paymentCfReturnGiftInquiryCode,
		GiftTypeOtherStatus:    null.IntFrom(int64(otherStatus)),
	}
}

func newCfInnReserveRequest() *entity.CfInnReserveRequest {
	req := newCfReserveRequest()

	return &entity.CfInnReserveRequest{
		UserID:           req.UserID,
		PaymentID:        req.PaymentID,
		CfReturnGiftID:   req.CfReturnGiftID,
		FirstName:        req.FirstName,
		LastName:         req.LastName,
		FirstNameKana:    req.FirstNameKana,
		LastNameKana:     req.LastNameKana,
		Email:            req.Email,
		PhoneNumber:      req.PhoneNumber,
		CheckinAt:        req.CheckinAt,
		CheckoutAt:       req.CheckoutAt,
		StayDays:         req.StayDays,
		AdultMemberCount: req.AdultMemberCount,
		ChildMemberCount: req.ChildMemberCount,
		Remark:           req.Remark,
	}
}
