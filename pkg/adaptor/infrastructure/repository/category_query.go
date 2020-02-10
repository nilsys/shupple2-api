package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type CategoryQueryRepositoryImpl struct {
	DB *gorm.DB
}

var CategoryQueryRepositorySet = wire.NewSet(
	wire.Struct(new(CategoryQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.CategoryQueryRepository), new(*CategoryQueryRepositoryImpl)),
)

func (r *CategoryQueryRepositoryImpl) FindByIDs(ids ...int) ([]*entity.Category, error) {
	var categories []*entity.Category

	if err := r.DB.Where("id IN (?)", ids).Find(&categories).Error; err != nil {
		return nil, errors.Wrapf(err, "Failed get categories")
	}

	return categories, nil
}

func (r *CategoryQueryRepositoryImpl) FindByID(id int) (*entity.Category, error) {
	var row entity.Category
	if err := r.DB.First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "category(id=%d)", id)
	}
	return &row, nil
}
