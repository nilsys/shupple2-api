package entity

type (
	Payment struct {
		ID                int `gorm:"primary_key"`
		UserID            int
		CardID            int
		ChargeID          string
		ShippingAddressID int
		Times
	}

	PaymentCfReturnGift struct {
		PaymentID              int `gorm:"primary_key"`
		CfReturnGiftID         int `gorm:"primary_key"`
		CfReturnGiftSnapshotID int
		CfProjectID            int
		CfProjectSnapshotID    int
		Amount                 int
		TimesWithoutDeletedAt
	}
)

func NewPayment(userID, cardID, addressID int, chargeID string) *Payment {
	return &Payment{
		UserID:            userID,
		CardID:            cardID,
		ChargeID:          chargeID,
		ShippingAddressID: addressID,
	}
}

// PaymentIDが先に取得できない為、後でいれる想定
func NewPaymentReturnGift(giftID, giftSnapshotID, projectID, projectSnapshotID, amount int) *PaymentCfReturnGift {
	return &PaymentCfReturnGift{
		CfReturnGiftID:         giftID,
		CfReturnGiftSnapshotID: giftSnapshotID,
		CfProjectID:            projectID,
		CfProjectSnapshotID:    projectSnapshotID,
		Amount:                 amount,
	}
}
