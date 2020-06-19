package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type ShippingQueryRepositoryImpl struct {
	DAO
}

var ShippingQueryRepositorySet = wire.NewSet(
	wire.Struct(new(ShippingQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.ShippingQueryRepository), new(*ShippingQueryRepositoryImpl)),
)

func (r *ShippingQueryRepositoryImpl) FindLatestShippingAddressByUserID(c context.Context, userID int) (*entity.ShippingAddress, error) {
	var row entity.ShippingAddress
	if err := r.DB(c).Where("user_id = ?", userID).Order("created_at desc").First(&row).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "shipping_address(user_id=%d)", userID)
	}
	return &row, nil
}
