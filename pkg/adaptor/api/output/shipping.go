package output

type (
	ShippingAddress struct {
		ID            int    `json:"id"`
		FirstName     string `json:"firstName"`
		LastName      string `json:"lastName"`
		FirstNameKana string `json:"firstNameKana"`
		LastNameKana  string `json:"lastNameKana"`
		PhoneNumber   string `json:"phoneNumber"`
		PostalNumber  string `json:"postalNumber"`
		Prefecture    string `json:"prefecture"`
		City          string `json:"city"`
		Address       string `json:"address"`
		Building      string `json:"building"`
		Email         string `json:"email"`
	}
)
