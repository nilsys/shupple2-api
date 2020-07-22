package input

type (
	Charge struct {
		CfReturnGiftID         int `json:"cfReturnGiftId"`
		CfReturnGiftSnapshotID int `json:"cfReturnGiftSnapshotId"`
		Amount                 int `json:"amount"`
	}

	CaptureCharge struct {
		Payments           []Charge `json:"payments" validate:"gt=0"`
		SupportCommentBody string   `json:"supportCommentBody"`
	}

	RefundCharge struct {
		IDParam
		CfReturnGiftID int `json:"cfReturnGiftId" validate:"required"`
	}

	ListPayment struct {
		PerPage int `query:"perPage"`
		Page    int `query:"page"`
	}
)

const defaultListPaymentPerPage = 10

func (i ListPayment) GetLimit() int {
	if i.PerPage == 0 {
		return defaultListPaymentPerPage
	}
	return i.PerPage
}

func (i ListPayment) GetOffSet() int {
	if i.Page == 1 || i.Page == 0 {
		return 0
	}
	return i.GetLimit() * (i.Page - 1)
}
