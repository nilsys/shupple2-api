package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type CfInnReserveRequestQueryRepositoryImpl struct {
	DB *gorm.DB
}

var CfInnReserveRequestQueryRepositorySet = wire.NewSet(
	wire.Struct(new(CfInnReserveRequestQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.CfInnReserveRequestQueryRepository), new(*CfInnReserveRequestQueryRepositoryImpl)),
)

func (r *CfInnReserveRequestQueryRepositoryImpl) IsExistByPaymentIDAndCfReturnGiftID(paymentID, cfReturnGiftID int) (bool, error) {
	var row entity.CfInnReserveRequest
	err := r.DB.Where("payment_id = ? AND cf_return_gift_id = ?", paymentID, cfReturnGiftID).First(&row).Error

	return ErrorToIsExist(err, "cf_inn_reserve_request(payment_id=%d,cf_return_gift_id=%d)", paymentID, cfReturnGiftID)
}
