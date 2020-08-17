package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type (
	CfInnReserveRequestCommandRepository interface {
		Store(ctx context.Context, request *entity.CfInnReserveRequest) error
	}

	CfInnReserveRequestQueryRepository interface {
		IsExistByPaymentIDAndCfReturnGiftID(paymentID, cfReturnGiftID int) (bool, error)
	}
)
