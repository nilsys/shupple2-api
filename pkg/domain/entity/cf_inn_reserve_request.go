package entity

import (
	"fmt"
	"time"
)

type (
	CfInnReserveRequest struct {
		UserID           int
		PaymentID        int
		CfReturnGiftID   int
		FirstName        string
		LastName         string
		FirstNameKana    string
		LastNameKana     string
		Email            string
		PhoneNumber      string
		CheckinAt        time.Time
		CheckoutAt       time.Time
		StayDays         int
		AdultMemberCount int
		ChildMemberCount int
		Remark           string
		TimesWithoutDeletedAt
	}
)

func (c *CfInnReserveRequest) FullNameMailFmt() string {
	return fmt.Sprintf("%s %s", c.LastName, c.FirstName)
}

func (c *CfInnReserveRequest) FullNameKanaMailFmt() string {
	return fmt.Sprintf("%s %s", c.LastNameKana, c.FirstNameKana)
}
