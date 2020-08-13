package input

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	CfInnReserveRequest struct {
		PaymentID        IDParam
		CfReturnGiftID   int             `json:"cfReturnGiftId" validate:"required"`
		FirstName        string          `json:"firstName" validate:"required"`
		LastName         string          `json:"lastName" validate:"required"`
		FirstNameKana    string          `json:"firstNameKana" validate:"required"`
		LastNameKana     string          `json:"lastNameKana" validate:"required"`
		Email            string          `json:"email" validate:"required"`
		PhoneNumber      string          `json:"phoneNumber" validate:"required"`
		Checkin          model.TimeFront `json:"checkin" validate:"required"`
		Checkout         model.TimeFront `json:"checkout" validate:"required"`
		StayDays         int             `json:"stayDays" validate:"required"`
		AdultMemberCount int             `json:"adultMemberCount"`
		ChildMemberCount int             `json:"childMemberCount"`
		Remark           string          `json:"remark"`
	}
)
