package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type MetasearchAreaQueryRepositoryImpl struct {
	DB *gorm.DB
}

var MetasearchAreaQueryRepositorySet = wire.NewSet(
	wire.Struct(new(MetasearchAreaQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.MetasearchAreaQueryRepository), new(*MetasearchAreaQueryRepositoryImpl)),
)

func (r *MetasearchAreaQueryRepositoryImpl) FindByMetasearchAreaID(metasearchAreaID int, metasearchAreaType model.AreaCategoryType) (*entity.MetasearchArea, error) {
	var row entity.MetasearchArea
	if err := r.DB.Where("metasearch_area_id = ? AND metasearch_area_type = ?", metasearchAreaID, metasearchAreaType).First(&row).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "metasearchArea(id=%d, type=%s)", metasearchAreaID, metasearchAreaType)
	}
	return &row, nil
}

func (r *MetasearchAreaQueryRepositoryImpl) FindByAreaCategoryID(areaCategoryID int, areaCategoryType model.AreaCategoryType) ([]*entity.MetasearchArea, error) {
	var rows []*entity.MetasearchArea
	q := r.DB.
		Joins("JOIN area_category ac ON area_category_id = ac.id").
		Where("area_category_id = ?", areaCategoryID).
		Where("ac.type = ?", areaCategoryType)

	if err := q.Find(&rows).Error; err != nil {
		return nil, errors.Wrapf(err, "failed to find metasearchArea(area_category_id=%d, area_category_type=%s)", areaCategoryID, areaCategoryType)
	}
	return rows, nil
}
