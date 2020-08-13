package converter

import (
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/command"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
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

	return &command.PaymentList{
		List:   result,
		Body:   payments.SupportCommentBody,
		Remark: payments.Remark,
	}
}

func (c Converters) ConvertListPaymentToQuery(i *input.ListPayment) *query.FindListPaginationQuery {
	return &query.FindListPaginationQuery{
		Limit:  i.GetLimit(),
		Offset: i.GetOffSet(),
	}
}

func (c Converters) ConvertPaymentListToOutput(payments *entity.PaymentList) output.PaymentList {
	response := make([]*output.Payment, len(payments.List))
	for i, tiny := range payments.List {
		response[i] = c.ConvertPaymentToOutput(tiny)
	}

	return output.PaymentList{
		TotalNumber: payments.TotalNumber,
		Payments:    response,
	}
}

func (c Converters) ConvertPaymentToOutput(payment *entity.Payment) *output.Payment {
	return &output.Payment{
		ID:                   payment.ID,
		ShippingAddress:      c.ConvertShippingAddressToOutput(payment.ShippingAddress),
		Card:                 c.ConvertCardToOutput(payment.Card),
		TotalPrice:           payment.TotalPrice,
		CommissionPrice:      payment.CommissionPrice,
		ChargeID:             payment.ChargeID,
		Remark:               payment.Remark,
		PaymentCfReturnGifts: c.ConvertPaymentCfReturnGiftToOutput(payment.PaymentCfReturnGift),
		OrderedAt:            model.TimeResponse(payment.CreatedAt),
	}
}

func (c Converters) ConvertPaymentCfReturnGiftToOutput(payment []*entity.PaymentCfReturnGift) []*output.PaymentCfReturnGift {
	response := make([]*output.PaymentCfReturnGift, len(payment))
	for i, tiny := range payment {
		var ownerConfirmedAt time.Time
		if tiny.OwnerConfirmedAt != nil {
			ownerConfirmedAt = *tiny.OwnerConfirmedAt
		}

		response[i] = &output.PaymentCfReturnGift{
			CfReturnGift: c.convertCfReturnGiftToOutput(&entity.CfReturnGift{
				CfReturnGiftTiny: *tiny.CfReturnGift,
				Snapshot:         tiny.CfReturnGiftSnapshot,
			}),
			Amount:                       tiny.Amount,
			InquiryCode:                  tiny.InquiryCode,
			GiftTypeOtherStatus:          tiny.ResolveGiftTypeOtherStatus(),
			GiftTypeReservedTicketStatus: tiny.ResolveGiftTypeReservedTicketStatus(),
			OwnerConfirmedAt:             model.TimeResponse(ownerConfirmedAt),
		}
	}
	return response
}

func (c Converters) ConvertReserveRequestToEntity(reserveReq *input.CfInnReserveRequest, userID int) *entity.CfInnReserveRequest {
	return &entity.CfInnReserveRequest{
		UserID:           userID,
		PaymentID:        reserveReq.PaymentID.ID,
		CfReturnGiftID:   reserveReq.CfReturnGiftID,
		FirstName:        reserveReq.FirstName,
		LastName:         reserveReq.LastName,
		FirstNameKana:    reserveReq.FirstNameKana,
		LastNameKana:     reserveReq.LastNameKana,
		Email:            reserveReq.Email,
		PhoneNumber:      reserveReq.PhoneNumber,
		CheckinAt:        time.Time(reserveReq.Checkin),
		CheckoutAt:       time.Time(reserveReq.Checkout),
		StayDays:         reserveReq.StayDays,
		AdultMemberCount: reserveReq.AdultMemberCount,
		ChildMemberCount: reserveReq.ChildMemberCount,
	}
}
