package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type BatchOptionCommandRepositoryImpl struct {
	DB *gorm.DB
}

var BatchOptionCommandRepositorySet = wire.NewSet(
	wire.Struct(new(BatchOptionCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.BatchOptionCommandRepository), new(*BatchOptionCommandRepositoryImpl)),
)

func (r *BatchOptionCommandRepositoryImpl) UpdateByName(name model.BatchOptionName, val string) error {
	if err := r.DB.Exec("UPDATE batch_option SET value = ? WHERE name = ?", val, name).Error; err != nil {
		return errors.Wrap(err, "failed update batch_option")
	}
	return nil
}
