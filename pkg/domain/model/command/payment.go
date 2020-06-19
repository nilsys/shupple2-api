package command

type (
	Payment struct {
		ReturnGiftID        int
		ReturnGiftSummaryID int
		Amount              int
	}

	PaymentList struct {
		Payments []*Payment
	}
)

func (p *PaymentList) ReturnIDs() []int {
	ids := make([]int, len(p.Payments))
	for i, payment := range p.Payments {
		ids[i] = payment.ReturnGiftID
	}
	return ids
}
