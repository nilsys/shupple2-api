package output

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

type (
	Payment struct {
		ID                   int                    `json:"id"`
		ShippingAddress      *ShippingAddress       `json:"shippingAddress"`
		Card                 *Card                  `json:"card"`
		TotalPrice           int                    `json:"totalPrice"`
		CommissionPrice      int                    `json:"commissionPrice"`
		ChargeID             string                 `json:"chargeId"`
		Remark               string                 `json:"remark"`
		PaymentCfReturnGifts []*PaymentCfReturnGift `json:"paymentCfReturnGifts"`
		OrderedAt            model.TimeResponse     `json:"orderedAt"`
	}

	PaymentCfReturnGift struct {
		CfReturnGift                 *CfReturnGift                                     `json:"cfReturnGift"`
		Amount                       int                                               `json:"amount"`
		InquiryCode                  string                                            `json:"inquiryCode"`
		GiftTypeOtherStatus          model.PaymentCfReturnGiftOtherTypeStatus          `json:"giftTypeOtherStatus"`
		GiftTypeReservedTicketStatus model.PaymentCfReturnGiftReservedTicketTypeStatus `json:"giftTypeReservedTicketStatus"`
		OwnerConfirmedAt             model.TimeResponse                                `json:"ownerConfirmedAt"`
	}

	PaymentList struct {
		TotalNumber int        `json:"totalNumber"`
		Payments    []*Payment `json:"payments"`
	}
)
