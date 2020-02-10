package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type CategoryCommandRepositoryImpl struct {
	DB *gorm.DB
}

var CategoryCommandRepositorySet = wire.NewSet(
	wire.Struct(new(CategoryCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.CategoryCommandRepository), new(*CategoryCommandRepositoryImpl)),
)

func (r *CategoryCommandRepositoryImpl) Store(category *entity.Category) error {
	return errors.Wrap(r.DB.Save(category).Error, "failed to save category")
}
