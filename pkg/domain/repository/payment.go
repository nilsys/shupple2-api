package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type (
	PaymentQueryRepository interface {
		FindByUserID(userID int, query *query.FindListPaginationQuery) (*entity.PaymentList, error)
		FindByID(id int) (*entity.Payment, error)
		FindTinyByID(id int) (*entity.PaymentTiny, error)
		FindPaymentCfReturnGiftByPaymentIDAndCfReturnGift(paymentID, cfReturnGiftID int) (*entity.PaymentCfReturnGift, error)
		LockPaymentCfReturnGift(ctx context.Context, paymentID, cfReturnGiftID int) (*entity.PaymentCfReturnGift, error)
	}

	PaymentCommandRepository interface {
		Store(c context.Context, payment *entity.PaymentTiny) error
		StorePaymentReturnGiftList(c context.Context, list []*entity.PaymentCfReturnGiftTiny, paymentID int) error
		MarkPaymentCfReturnGiftAsCancel(c context.Context, paymentID, cfReturnID int) error
		MarkPaymentCfReturnGiftAsReserved(c context.Context, paymentID, cfReturnGiftID int) error
	}
)
