package input

type (
	ShippingAddress struct {
		FirstName     string `json:"firstName" validate:"required"`
		LastName      string `json:"lastName" validate:"required"`
		FirstNameKana string `json:"firstNameKana" validate:"required"`
		LastNameKana  string `json:"lastNameKana" validate:"required"`
		PhoneNumber   string `json:"phoneNumber" validate:"required"`
		PostalNumber  string `json:"postalNumber" validate:"required"`
		Prefecture    string `json:"prefecture" validate:"required"`
		City          string `json:"city" validate:"required"`
		Address       string `json:"address" validate:"required"`
		Building      string `json:"building"`
		Email         string `json:"email" validate:"required"`
	}
)
