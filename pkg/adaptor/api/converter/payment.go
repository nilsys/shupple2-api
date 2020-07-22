package converter

import (
	"github.com/payjp/payjp-go/v1"
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

	return &command.PaymentList{List: result, Body: payments.SupportCommentBody}
}

func (c Converters) ConvertListPaymentToQuery(i *input.ListPayment) *query.FindListPaginationQuery {
	return &query.FindListPaginationQuery{
		Limit:  i.GetLimit(),
		Offset: i.GetOffSet(),
	}
}

func (c Converters) ConvertPaymentListToOutput(payments *entity.PaymentList, cardIDMap map[string]*payjp.CardResponse) []*output.Payment {
	response := make([]*output.Payment, len(payments.List))
	for i, tiny := range payments.List {
		response[i] = c.ConvertPaymentToOutput(tiny, cardIDMap[tiny.Card.CardID])
	}
	return response
}

func (c Converters) ConvertPaymentToOutput(payment *entity.Payment, card *payjp.CardResponse) *output.Payment {
	return &output.Payment{
		ID:                   payment.ID,
		ShippingAddress:      c.ConvertShippingAddressToOutput(payment.ShippingAddress),
		Card:                 c.ConvertCardToOutput(card),
		TotalPrice:           payment.TotalPrice,
		ChargeID:             payment.ChargeID,
		PaymentCfReturnGifts: c.ConvertPaymentCfReturnGiftToOutput(payment.PaymentCfReturnGift),
		OrderedAt:            model.TimeResponse(payment.CreatedAt),
	}
}

func (c Converters) ConvertPaymentCfReturnGiftToOutput(payment []*entity.PaymentCfReturnGift) []*output.PaymentCfReturnGift {
	response := make([]*output.PaymentCfReturnGift, len(payment))
	for i, tiny := range payment {
		response[i] = &output.PaymentCfReturnGift{
			CfReturnGift: c.convertCfReturnGiftToOutput(&entity.CfReturnGift{
				CfReturnGiftTiny: *tiny.CfReturnGift,
				Snapshot:         tiny.CfReturnGiftSnapshot,
				// MEMO: ユーザーの購入一覧APIではsoldcount使わないので0入れてる。。
			}, 0),
			Amount:                       tiny.Amount,
			GiftTypeOtherStatus:          tiny.ResolveGiftTypeOtherStatus(),
			GiftTypeReservedTicketStatus: tiny.ResolveGiftTypeReservedTicketStatus(),
			OwnerConfirmedAt:             model.TimeResponse(*tiny.OwnerConfirmedAt),
		}
	}
	return response
}

func (c Converters) ConvertReserveRequestToEntity(reserveReq *input.CfReserveRequest) *entity.CfReserveRequest {
	return &entity.CfReserveRequest{
		FirstName:        reserveReq.FirstName,
		LastName:         reserveReq.LastName,
		FirstNameKana:    reserveReq.FirstNameKana,
		LastNameKana:     reserveReq.LastNameKana,
		Email:            reserveReq.Email,
		PhoneNumber:      reserveReq.PhoneNumber,
		Checkin:          reserveReq.Checkin,
		Checkout:         reserveReq.Checkout,
		StayDays:         reserveReq.StayDays,
		AdultMemberCount: reserveReq.AdultMemberCount,
		ChildMemberCount: reserveReq.ChildMemberCount,
	}
}
