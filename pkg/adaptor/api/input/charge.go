package input

type (
	Charge struct {
		CfReturnGiftID         int `json:"cfReturnGiftId"`
		CfReturnGiftSnapshotID int `json:"cfReturnGiftSnapshotId"`
		Amount                 int `json:"amount"`
	}

	CreateCharge struct {
		List               []Charge `json:"payments" validate:"gt=0"`
		SupportCommentBody string   `json:"supportCommentBody"`
		Remark             string   `json:"remark" validate:"lt=500"`
		AssociateID        string   `json:"associateId"`
	}

	InstantCreateCharge struct {
		CreateCharge
		ShippingAddress *ShippingAddress `json:"shippingAddress" validate:"required"`
		CardToken       string           `json:"cardToken" validate:"required"`
	}

	RefundCharge struct {
		IDParam
		CfReturnGiftID int `json:"cfReturnGiftId" validate:"required"`
	}
)
