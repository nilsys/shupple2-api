package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type BatchOptionQueryRepositoryImpl struct {
	DB *gorm.DB
}

var BatchOptionQueryRepositorySet = wire.NewSet(
	wire.Struct(new(BatchOptionQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.BatchOptionQueryRepository), new(*BatchOptionQueryRepositoryImpl)),
)

func (r *BatchOptionQueryRepositoryImpl) FirstOrCreateByName(name model.BatchOptionName) (string, error) {
	var row entity.BatchOption

	if err := r.DB.Where("name = ?", name).FirstOrCreate(&row).Error; err != nil {
		return "", ErrorToFindSingleRecord(err, "batch_option(name=%d)", name)
	}

	return row.Value, nil
}
