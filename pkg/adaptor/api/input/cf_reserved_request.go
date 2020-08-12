package input

type (
	CfReserveRequest struct {
		PaymentID        IDParam
		CfReturnGiftID   int    `json:"cfReturnGiftId" validate:"required"`
		FirstName        string `json:"firstName" validate:"required"`
		LastName         string `json:"lastName" validate:"required"`
		FirstNameKana    string `json:"firstNameKana" validate:"required"`
		LastNameKana     string `json:"lastNameKana" validate:"required"`
		Email            string `json:"email" validate:"required"`
		PhoneNumber      string `json:"phoneNumber" validate:"required"`
		Checkin          string `json:"checkin" validate:"required"`
		Checkout         string `json:"checkout" validate:"required"`
		StayDays         int    `json:"stayDays" validate:"required"`
		AdultMemberCount int    `json:"adultMemberCount"`
		ChildMemberCount int    `json:"childMemberCount"`
	}
)
