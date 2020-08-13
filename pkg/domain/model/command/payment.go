package command

type (
	Payment struct {
		ReturnGiftID         int
		ReturnGiftSnapshotID int
		Amount               int
	}

	PaymentList struct {
		List   []*Payment
		Body   string
		Remark string
	}
)

func (p *PaymentList) ReturnIDs() []int {
	ids := make([]int, len(p.List))
	for i, payment := range p.List {
		ids[i] = payment.ReturnGiftID
	}
	return ids
}
