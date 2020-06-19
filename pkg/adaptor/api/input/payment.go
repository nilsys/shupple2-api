package input

type (
	CreateCharge struct {
		CfReturnGiftID         int `json:"cfReturnGiftId"`
		CfReturnGiftSnapshotID int `json:"cfReturnGiftSnapshotId"`
		Amount                 int `json:"amount"`
	}

	CreateChargeList struct {
		Payments []CreateCharge `json:"payments"`
	}

	CapturePayment struct {
		ID int `param:"id"`
	}
)
