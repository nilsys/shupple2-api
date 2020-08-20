package service

import (
	"context"

	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	ShippingCommandService interface {
		StoreShippingAddress(user *entity.User, address *entity.ShippingAddress) error
	}

	ShippingCommandServiceImpl struct {
		repository.ShippingCommandRepository
		repository.ShippingQueryRepository
	}
)

var ShippingCommandServiceSet = wire.NewSet(
	wire.Struct(new(ShippingCommandServiceImpl), "*"),
	wire.Bind(new(ShippingCommandService), new(*ShippingCommandServiceImpl)),
)

// 履歴を残す為、dbには更新ではなく、追加する
func (s *ShippingCommandServiceImpl) StoreShippingAddress(user *entity.User, address *entity.ShippingAddress) error {
	return s.ShippingCommandRepository.StoreShippingAddress(context.Background(), address)
}
