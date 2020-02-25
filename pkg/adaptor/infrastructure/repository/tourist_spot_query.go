package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type TouristSpotQueryRepositoryImpl struct {
	DB *gorm.DB
}

var TouristSpotQueryRepositorySet = wire.NewSet(
	wire.Struct(new(TouristSpotQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.TouristSpotQueryRepository), new(*TouristSpotQueryRepositoryImpl)),
)

func (r *TouristSpotQueryRepositoryImpl) FindByID(id int) (*entity.TouristSpot, error) {
	var row entity.TouristSpot
	if err := r.DB.First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "touristSpot(id=%d)", id)
	}
	return &row, nil
}

// name部分一致検索
func (r *TouristSpotQueryRepositoryImpl) SearchByName(name string) ([]*entity.TouristSpot, error) {
	var rows []*entity.TouristSpot

	if err := r.DB.Where("MATCH(name) AGAINST(?)", name).Limit(defaultSearchSuggestionsNumber).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find tourist_spot list by like name")
	}

	return rows, nil
}
