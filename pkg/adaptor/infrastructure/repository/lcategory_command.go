package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type LcategoryCommandRepositoryImpl struct {
	DB *gorm.DB
}

var LcategoryCommandRepositorySet = wire.NewSet(
	wire.Struct(new(LcategoryCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.LcategoryCommandRepository), new(*LcategoryCommandRepositoryImpl)),
)

func (r *LcategoryCommandRepositoryImpl) Store(lcategory *entity.Lcategory) error {
	return errors.Wrap(r.DB.Save(lcategory).Error, "failed to save lcategory")
}
