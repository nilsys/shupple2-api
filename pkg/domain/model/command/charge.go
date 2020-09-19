package command

type (
	Charge struct {
		ReturnGiftID         int
		ReturnGiftSnapshotID int
		Amount               int
	}

	CreateCharge struct {
		List               []*Charge
		SupportCommentBody string
		Remark             string
		AssociateID        string
	}
)

func (p *CreateCharge) ReturnIDs() []int {
	ids := make([]int, len(p.List))
	for i, payment := range p.List {
		ids[i] = payment.ReturnGiftID
	}
	return ids
}
