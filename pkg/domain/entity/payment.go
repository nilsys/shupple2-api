package entity

type (
	Payment struct {
		ID                int `gorm:"primary_key"`
		UserID            int
		ProjectOwnerID    int
		CardID            int
		ChargeID          string
		ShippingAddressID int
		TotalPrice        int
		Times
	}

	PaymentCfReturnGift struct {
		PaymentID              int `gorm:"primary_key"`
		CfReturnGiftID         int `gorm:"primary_key"`
		CfReturnGiftSnapshotID int
		CfProjectID            int
		CfProjectSnapshotID    int
		Amount                 int
		IsCanceled             bool
		IsOwnerConfirmed       bool
		TimesWithoutDeletedAt
	}
)

func NewPayment(userID, projectOwnerID, cardID, addressID int, chargeID string, price int) *Payment {
	return &Payment{
		UserID:            userID,
		ProjectOwnerID:    projectOwnerID,
		CardID:            cardID,
		ChargeID:          chargeID,
		ShippingAddressID: addressID,
		TotalPrice:        price,
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
