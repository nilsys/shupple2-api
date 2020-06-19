package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type ShippingQueryRepositoryImpl struct {
	DB *gorm.DB
}

var ShippingQueryRepositorySet = wire.NewSet(
	wire.Struct(new(ShippingQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.ShippingQueryRepository), new(*ShippingQueryRepositoryImpl)),
)

func (r *ShippingQueryRepositoryImpl) FindLatestShippingAddressByUserID(userID int) (*entity.ShippingAddress, error) {
	var row entity.ShippingAddress
	if err := r.DB.Where("user_id = ?", userID).Order("created_at desc").First(&row).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "shipping_address(user_id=%d)", userID)
	}
	return &row, nil
}
