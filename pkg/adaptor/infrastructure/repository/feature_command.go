package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type FeatureCommandRepositoryImpl struct {
	DB *gorm.DB
}

var FeatureCommandRepositorySet = wire.NewSet(
	wire.Struct(new(FeatureCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.FeatureCommandRepository), new(*FeatureCommandRepositoryImpl)),
)

func (r *FeatureCommandRepositoryImpl) Store(feature *entity.Feature) error {
	return errors.Wrap(r.DB.Save(feature).Error, "failed to save feature")
}
