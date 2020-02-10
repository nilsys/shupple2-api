package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type LcategoryQueryRepositoryImpl struct {
	DB *gorm.DB
}

var LcategoryQueryRepositorySet = wire.NewSet(
	wire.Struct(new(LcategoryQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.LcategoryQueryRepository), new(*LcategoryQueryRepositoryImpl)),
)

func (r *LcategoryQueryRepositoryImpl) FindByID(id int) (*entity.Lcategory, error) {
	var row entity.Lcategory
	if err := r.DB.First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "lcategory(id=%d)", id)
	}
	return &row, nil
}
