package output

type (
	Card struct {
		ID      int    `json:"id"`
		Last4   string `json:"last4"`
		Expired string `json:"expired"`
	}
)
