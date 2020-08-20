package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type ShippingCommandRepositoryImpl struct {
	DAO
}

var ShippingCommandRepositorySet = wire.NewSet(
	wire.Struct(new(ShippingCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.ShippingCommandRepository), new(*ShippingCommandRepositoryImpl)),
)

func (r *ShippingCommandRepositoryImpl) StoreShippingAddress(ctx context.Context, address *entity.ShippingAddress) error {
	if err := r.DB(ctx).Save(address).Error; err != nil {
		return errors.Wrap(err, "failed store address")
	}
	return nil
}
