package repository

import "github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

type (
	ShippingCommandRepository interface {
		StoreShippingAddress(address *entity.ShippingAddress) error
	}

	ShippingQueryRepository interface {
		FindLatestShippingAddressByUserID(userID int) (*entity.ShippingAddress, error)
	}
)
