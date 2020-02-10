package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type FeatureQueryRepositoryImpl struct {
	DB *gorm.DB
}

var FeatureQueryRepositorySet = wire.NewSet(
	wire.Struct(new(FeatureQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.FeatureQueryRepository), new(*FeatureQueryRepositoryImpl)),
)

func (r *FeatureQueryRepositoryImpl) FindByID(id int) (*entity.Feature, error) {
	var row entity.Feature
	if err := r.DB.First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "feature(id=%d)", id)
	}
	return &row, nil
}
