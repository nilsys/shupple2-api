package service

import (
	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	ShippingQueryService interface {
		ShowShippingAddressByUserID(user *entity.User) (*entity.ShippingAddress, error)
	}

	ShippingQueryServiceImpl struct {
		repository.ShippingQueryRepository
	}
)

var ShippingQueryServiceSet = wire.NewSet(
	wire.Struct(new(ShippingQueryServiceImpl), "*"),
	wire.Bind(new(ShippingQueryService), new(*ShippingQueryServiceImpl)),
)

func (s *ShippingQueryServiceImpl) ShowShippingAddressByUserID(user *entity.User) (*entity.ShippingAddress, error) {
	return s.ShippingQueryRepository.FindLatestShippingAddressByUserID(user.ID)
}
