package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type (
	ShippingCommandRepository interface {
		StoreShippingAddress(ctx context.Context, address *entity.ShippingAddress) error
	}

	ShippingQueryRepository interface {
		FindLatestShippingAddressByUserID(c context.Context, userID int) (*entity.ShippingAddress, error)
	}
)
