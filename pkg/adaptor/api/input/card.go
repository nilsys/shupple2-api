package input

type (
	StoreCard struct {
		CardToken string `json:"cardToken" validate:"required"`
	}
)
