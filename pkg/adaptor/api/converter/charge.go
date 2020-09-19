package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/command"
)

func (c Converters) ConvertCaptureChargeToCmd(charge *input.CreateCharge) *command.CreateCharge {
	idAmountMap := make(map[int]int, len(charge.List))
	idSummaryIDMap := make(map[int]int, len(charge.List))

	for _, payment := range charge.List {
		idSummaryIDMap[payment.CfReturnGiftID] = payment.CfReturnGiftSnapshotID
		idAmountMap[payment.CfReturnGiftID] += payment.Amount
	}

	result := make([]*command.Charge, 0, len(idAmountMap))
	for id, amount := range idAmountMap {
		result = append(result, &command.Charge{
			ReturnGiftID:         id,
			ReturnGiftSnapshotID: idSummaryIDMap[id],
			Amount:               amount,
		})
	}

	return &command.CreateCharge{
		List:               result,
		SupportCommentBody: charge.SupportCommentBody,
		Remark:             charge.Remark,
		AssociateID:        charge.AssociateID,
	}
}
