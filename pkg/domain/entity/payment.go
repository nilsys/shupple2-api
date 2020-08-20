package entity

import (
	"time"

	"gopkg.in/guregu/null.v3"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

const (
	// https://pay.jp/docs/api/#%E6%94%AF%E6%89%95%E3%81%84%E6%83%85%E5%A0%B1%E3%82%92%E6%9B%B4%E6%96%B0
	chargeRefundExpired = 4320 * time.Hour
)

type (
	PaymentTiny struct {
		ID                      int `gorm:"primary_key"`
		UserID                  int
		ProjectOwnerID          int
		CardID                  int
		ChargeID                string
		ShippingAddressID       int
		TotalPrice              int
		CommissionPrice         int
		Remark                  string
		OwnerDepositRequestedAt *time.Time
		Times
	}

	// StatusはCfReturnGift.GiftTypeに対応したfieldに値が入る
	// 他のStatusは必ずNULLになる
	// 全てのStatusがNULLの場合は無い
	// app内ではenumが用意されているのでそれを介して扱う事
	// Ex.) CfReturnGift.GiftType == Other の場合
	//     	GiftTypeOtherStatusはNOT NULL
	//      それ以外のStatusはNULL
	PaymentCfReturnGiftTiny struct {
		PaymentID                    int `gorm:"primary_key"`
		CfReturnGiftID               int `gorm:"primary_key"`
		CfReturnGiftSnapshotID       int
		CfProjectID                  int
		CfProjectSnapshotID          int
		Amount                       int
		OwnerConfirmedAt             *time.Time
		UserReserveRequestedAt       *time.Time
		GiftTypeOtherStatus          null.Int
		GiftTypeReservedTicketStatus null.Int
		InquiryCode                  string
		TimesWithoutDeletedAt
	}

	Payment struct {
		PaymentTiny
		ShippingAddress     *ShippingAddress       `gorm:"foreignkey:ID;association_foreignkey:ShippingAddressID"`
		Card                *Card                  `gorm:"foreignkey:ID;association_foreignkey:CardID"`
		PaymentCfReturnGift []*PaymentCfReturnGift `gorm:"foreignkey:PaymentID;association_foreignkey:ID"`
		Owner               *User                  `gorm:"foreignkey:ID;association_foreignkey:ProjectOwnerID"`
	}

	PaymentCfReturnGift struct {
		PaymentCfReturnGiftTiny
		CfReturnGift         *CfReturnGiftTiny         `gorm:"foreignkey:ID;association_foreignkey:CfReturnGiftID"`
		CfReturnGiftSnapshot *CfReturnGiftSnapshotTiny `gorm:"foreignkey:ID;association_foreignkey:CfReturnGiftSnapshotID"`
	}

	PaymentList struct {
		TotalNumber int
		List        []*Payment
	}
)

func NewPaymentTiny(userID, projectOwnerID, cardID, addressID int, chargeID string, price, commissionPrice int, remark string) *PaymentTiny {
	return &PaymentTiny{
		UserID:            userID,
		ProjectOwnerID:    projectOwnerID,
		CardID:            cardID,
		ChargeID:          chargeID,
		ShippingAddressID: addressID,
		TotalPrice:        price,
		CommissionPrice:   commissionPrice,
		Remark:            remark,
	}
}

func (p *PaymentTiny) TableName() string {
	return "payment"
}

func (p *PaymentCfReturnGiftTiny) TableName() string {
	return "payment_cf_return_gift"
}

func (p *PaymentCfReturnGiftTiny) ResolveGiftTypeOtherStatus() model.PaymentCfReturnGiftOtherTypeStatus {
	return model.PaymentCfReturnGiftOtherTypeStatus(int(p.GiftTypeOtherStatus.Int64))
}

func (p *PaymentCfReturnGiftTiny) ResolveGiftTypeReservedTicketStatus() model.PaymentCfReturnGiftReservedTicketTypeStatus {
	return model.PaymentCfReturnGiftReservedTicketTypeStatus(int(p.GiftTypeReservedTicketStatus.Int64))
}

// PaymentIDが先に取得できない為、後でいれる想定
func NewPaymentReturnGiftForOther(giftID, giftSnapshotID, projectID, projectSnapshotID, amount int, inquiryCode string) *PaymentCfReturnGiftTiny {
	return &PaymentCfReturnGiftTiny{
		CfReturnGiftID:         giftID,
		CfReturnGiftSnapshotID: giftSnapshotID,
		CfProjectID:            projectID,
		CfProjectSnapshotID:    projectSnapshotID,
		Amount:                 amount,
		GiftTypeOtherStatus:    null.IntFrom(int64(model.PaymentCfReturnGiftOtherTypeStatusOwnerUnconfirmed)),
		InquiryCode:            inquiryCode,
	}
}

func (p *PaymentCfReturnGift) TotalPrice() int {
	return p.CfReturnGiftSnapshot.Price * p.Amount
}

// PaymentIDが先に取得できない為、後でいれる想定
func NewPaymentReturnGiftForReservedTicket(giftID, giftSnapshotID, projectID, projectSnapshotID, amount int, inquiryCode string) *PaymentCfReturnGiftTiny {
	now := time.Now()
	return &PaymentCfReturnGiftTiny{
		CfReturnGiftID:               giftID,
		CfReturnGiftSnapshotID:       giftSnapshotID,
		CfProjectID:                  projectID,
		CfProjectSnapshotID:          projectSnapshotID,
		Amount:                       amount,
		OwnerConfirmedAt:             &now,
		GiftTypeReservedTicketStatus: null.IntFrom(int64(model.PaymentCfReturnGiftReservedTicketTypeStatusUnreserved)),
		InquiryCode:                  inquiryCode,
	}
}

func (p *PaymentCfReturnGift) CanCancel() bool {
	// キャンセル可能なリターンで無い場合
	if !p.CfReturnGiftSnapshot.IsCancelable || p.CfReturnGift.GiftType == model.CfReturnGiftTypeReservedTicket {
		return false
	}

	if !p.ResolveGiftTypeOtherStatus().CanTransit(model.PaymentCfReturnGiftOtherTypeStatusCanceled) {
		return false
	}

	// 発送が行われた場合
	if p.OwnerConfirmedAt != nil {
		return false
	}

	// キャンセル期限(180日)を過ぎている場合
	if p.CreatedAt.Add(chargeRefundExpired).Before(time.Now()) {
		return false
	}

	return true
}
