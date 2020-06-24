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
)
