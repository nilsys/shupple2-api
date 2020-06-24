package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type (
	PaymentCommandRepository interface {
		Store(c context.Context, payment *entity.Payment) error
		StorePaymentReturnGiftList(c context.Context, list []*entity.PaymentCfReturnGift, paymentID int) error
	}
)
