package input

type (
	StoreCard struct {
		CardToken string `form:"payjp-token" validate:"required"`
	}
)
