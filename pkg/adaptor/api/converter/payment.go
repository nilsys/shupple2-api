package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/command"
)

func (c Converters) ConvertPaymentsToCmd(payments *input.CaptureCharge) *command.PaymentList {
	idAmountMap := make(map[int]int, len(payments.Payments))
	idSummaryIDMap := make(map[int]int, len(payments.Payments))

	for _, payment := range payments.Payments {
		idSummaryIDMap[payment.CfReturnGiftID] = payment.CfReturnGiftSnapshotID
		idAmountMap[payment.CfReturnGiftID] += payment.Amount
	}

	result := make([]*command.Payment, 0, len(idAmountMap))
	for id, amount := range idAmountMap {
		result = append(result, &command.Payment{
			ReturnGiftID:         id,
			ReturnGiftSnapshotID: idSummaryIDMap[id],
			Amount:               amount,
		})
	}

	return &command.PaymentList{List: result, Body: payments.SupportCommentBody}
}
