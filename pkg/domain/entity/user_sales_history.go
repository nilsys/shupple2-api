package entity

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"gopkg.in/guregu/null.v3"
)

type (
	UserSalesHistoryTiny struct {
		UserID           int
		PaymentID        null.Int
		DepositRequestID null.Int
		Reason           model.UserSalesReasonType
		Price            int
		TimesWithoutDeletedAt
	}
)

func (u *UserSalesHistoryTiny) TableName() string {
	return "user_sales_history"
}

func NewUserSalesHistoryTiny(userID int, paymentID, depositRequestID null.Int, reason model.UserSalesReasonType, price int) *UserSalesHistoryTiny {
	return &UserSalesHistoryTiny{
		UserID:           userID,
		PaymentID:        paymentID,
		DepositRequestID: depositRequestID,
		Reason:           reason,
		Price:            price,
	}
}
