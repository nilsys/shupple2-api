package entity

import (
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
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
		OwnerDepositRequestedAt *time.Time
		Times
	}

	PaymentCfReturnGiftTiny struct {
		PaymentID                    int `gorm:"primary_key"`
		CfReturnGiftID               int `gorm:"primary_key"`
		CfReturnGiftSnapshotID       int
		CfProjectID                  int
		CfProjectSnapshotID          int
		Amount                       int
		OwnerConfirmedAt             *time.Time
		GiftTypeOtherStatus          model.PaymentCfReturnGiftOtherTypeStatus
		GiftTypeReservedTicketStatus model.PaymentCfReturnGiftReservedTicketTypeStatus
		TimesWithoutDeletedAt
	}

	PaymentCfReturnGift struct {
		PaymentCfReturnGiftTiny
		CfReturnGift         *CfReturnGiftTiny         `gorm:"foreignkey:ID;association_foreignkey:CfReturnGiftID"`
		CfReturnGiftSnapshot *CfReturnGiftSnapshotTiny `gorm:"foreignkey:ID;association_foreignkey:CfReturnGiftSnapshotID"`
	}

	Payment struct {
		PaymentTiny
		Owner *User `gorm:"foreignkey:ID;association_foreignkey:ProjectOwnerID"`
	}
)

func NewPayment(userID, projectOwnerID, cardID, addressID int, chargeID string, price int) *PaymentTiny {
	return &PaymentTiny{
		UserID:            userID,
		ProjectOwnerID:    projectOwnerID,
		CardID:            cardID,
		ChargeID:          chargeID,
		ShippingAddressID: addressID,
		TotalPrice:        price,
	}
}

// PaymentIDが先に取得できない為、後でいれる想定
func NewPaymentReturnGiftForOther(giftID, giftSnapshotID, projectID, projectSnapshotID, amount int) *PaymentCfReturnGiftTiny {
	return &PaymentCfReturnGiftTiny{
		CfReturnGiftID:         giftID,
		CfReturnGiftSnapshotID: giftSnapshotID,
		CfProjectID:            projectID,
		CfProjectSnapshotID:    projectSnapshotID,
		Amount:                 amount,
	}
}

// PaymentIDが先に取得できない為、後でいれる想定
func NewPaymentReturnGiftForReservedTicket(giftID, giftSnapshotID, projectID, projectSnapshotID, amount int) *PaymentCfReturnGiftTiny {
	now := time.Now()
	return &PaymentCfReturnGiftTiny{
		CfReturnGiftID:               giftID,
		CfReturnGiftSnapshotID:       giftSnapshotID,
		CfProjectID:                  projectID,
		CfProjectSnapshotID:          projectSnapshotID,
		Amount:                       amount,
		OwnerConfirmedAt:             &now,
		GiftTypeReservedTicketStatus: model.PaymentCfReturnGiftReservedTicketTypeStatusUnreserved,
	}
}
