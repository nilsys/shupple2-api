package model

func (x PaymentCfReturnGiftReservedTicketTypeStatus) CanTransit(to PaymentCfReturnGiftReservedTicketTypeStatus) bool {
	allowState := x.allowedTransit()
	_, ok := allowState[to]
	return ok
}

func (x PaymentCfReturnGiftReservedTicketTypeStatus) allowedTransit() map[PaymentCfReturnGiftReservedTicketTypeStatus]struct{} {
	switch x {
	case PaymentCfReturnGiftReservedTicketTypeStatusUnreserved:
		return map[PaymentCfReturnGiftReservedTicketTypeStatus]struct{}{
			PaymentCfReturnGiftReservedTicketTypeStatusReserved: {},
		}
	case PaymentCfReturnGiftReservedTicketTypeStatusReserved:
		return map[PaymentCfReturnGiftReservedTicketTypeStatus]struct{}{}
	default:
		return map[PaymentCfReturnGiftReservedTicketTypeStatus]struct{}{}
	}
}

func (x PaymentCfReturnGiftOtherTypeStatus) CanTransit(to PaymentCfReturnGiftOtherTypeStatus) bool {
	allowState := x.allowedTransit()
	_, ok := allowState[to]
	return ok
}

func (x PaymentCfReturnGiftOtherTypeStatus) allowedTransit() map[PaymentCfReturnGiftOtherTypeStatus]struct{} {
	switch x {
	case PaymentCfReturnGiftOtherTypeStatusOwnerUnconfirmed:
		return map[PaymentCfReturnGiftOtherTypeStatus]struct{}{
			PaymentCfReturnGiftOtherTypeStatusOwnerConfirmed: {},
			PaymentCfReturnGiftOtherTypeStatusCanceled:       {},
		}
	case PaymentCfReturnGiftOtherTypeStatusOwnerConfirmed:
		return map[PaymentCfReturnGiftOtherTypeStatus]struct{}{
			PaymentCfReturnGiftOtherTypeStatusOwnerUnconfirmed: {},
		}
	case PaymentCfReturnGiftOtherTypeStatusCanceled:
		return map[PaymentCfReturnGiftOtherTypeStatus]struct{}{}
	default:
		return map[PaymentCfReturnGiftOtherTypeStatus]struct{}{}
	}
}
