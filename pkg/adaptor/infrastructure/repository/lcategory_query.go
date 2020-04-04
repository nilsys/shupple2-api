package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type SpotCategoryQueryRepositoryImpl struct {
	DB *gorm.DB
}

var SpotCategoryQueryRepositorySet = wire.NewSet(
	wire.Struct(new(SpotCategoryQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.SpotCategoryQueryRepository), new(*SpotCategoryQueryRepositoryImpl)),
)

func (r *SpotCategoryQueryRepositoryImpl) FindByID(id int) (*entity.SpotCategory, error) {
	var row entity.SpotCategory
	if err := r.DB.First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "spotCategory(id=%d)", id)
	}
	return &row, nil
}
