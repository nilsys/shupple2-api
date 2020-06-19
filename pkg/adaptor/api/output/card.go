package output

type (
	Card struct {
		ID      string `json:"id"`
		Last4   string `json:"last4"`
		Expired string `json:"expired"`
	}
)

func NewCard(id, last4, expired string) *Card {
	return &Card{
		ID:      id,
		Last4:   last4,
		Expired: expired,
	}
}
