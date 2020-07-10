package entity

import "fmt"

type (
	CfReserveRequest struct {
		FirstName        string
		LastName         string
		FirstNameKana    string
		LastNameKana     string
		Email            string
		PhoneNumber      string
		Checkin          string
		Checkout         string
		StayDays         int
		AdultMemberCount int
		ChildMemberCount int
	}
)

func (c *CfReserveRequest) FullNameMailFmt() string {
	return fmt.Sprintf("%s %s", c.LastName, c.FirstName)
}

func (c *CfReserveRequest) FullNameKanaMailFmt() string {
	return fmt.Sprintf("%s %s", c.LastNameKana, c.FirstNameKana)
}
