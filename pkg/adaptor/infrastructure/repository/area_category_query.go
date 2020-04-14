package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type AreaCategoryQueryRepositoryImpl struct {
	DB *gorm.DB
}

var AreaCategoryQueryRepositorySet = wire.NewSet(
	wire.Struct(new(AreaCategoryQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.AreaCategoryQueryRepository), new(*AreaCategoryQueryRepositoryImpl)),
)

func (r *AreaCategoryQueryRepositoryImpl) FindByID(id int) (*entity.AreaCategory, error) {
	var row entity.AreaCategory
	if err := r.DB.First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "areaCategory(id=%d)", id)
	}
	return &row, nil
}

func (r *AreaCategoryQueryRepositoryImpl) FindByIDAndType(id int, areaCategoryType model.AreaCategoryType) (*entity.AreaCategory, error) {
	var row entity.AreaCategory
	if err := r.DB.Where("id = ? AND type = ?", id, areaCategoryType).First(&row).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "areaCategory(id=%d, type=%s)", id, areaCategoryType)
	}
	return &row, nil
}

func (r *AreaCategoryQueryRepositoryImpl) FindDetailByIDAndType(id int, areaCategoryType model.AreaCategoryType) (*entity.AreaCategoryDetail, error) {
	var row entity.AreaCategoryDetail
	var areaRow entity.AreaCategory
	var subAreaRow entity.AreaCategory
	var subSubAreaRow entity.AreaCategory
	if err := r.DB.Where("id = ? AND type = ?", id, areaCategoryType).First(&row.AreaCategory).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "areaCategory(id=%d, type=%s)", id, areaCategoryType)
	}
	if err := r.DB.Where("id = ? AND type = ?", row.AreaID, model.AreaCategoryTypeArea).First(&areaRow).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "areaCategory(id=%d, type=%s)", id, model.AreaCategoryTypeArea)
	}
	row.SetArea(&areaRow)

	// sub_area_idがnullじゃない時
	if row.SubAreaID.Valid {
		if err := r.DB.Where("id = ? AND type = ?", row.SubAreaID, model.AreaCategoryTypeSubArea).First(&subAreaRow).Error; err != nil {
			return nil, ErrorToFindSingleRecord(err, "areaCategory(id=%d, type=%s)", id, model.AreaCategoryTypeSubArea)
		}
		row.SetSubArea(&subAreaRow)
	}

	// sub_sub_area_idがnullじゃない時
	if row.SubSubAreaID.Valid {
		if err := r.DB.Where("id = ? AND type = ?", row.SubSubAreaID, model.AreaCategoryTypeSubSubArea).First(&subSubAreaRow).Error; err != nil {
			return nil, ErrorToFindSingleRecord(err, "areaCategory(id=%d, type=%s)", id, model.AreaCategoryTypeSubSubArea)
		}
		row.SetSubSubArea(&subSubAreaRow)
	}

	return &row, nil
}

func (r *AreaCategoryQueryRepositoryImpl) FindBySlug(slug string) (*entity.AreaCategory, error) {
	var row entity.AreaCategory
	if err := r.DB.Where("slug = ?", slug).First(&row).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "areaCategory(slug=%s)", slug)
	}
	return &row, nil
}

func (r *AreaCategoryQueryRepositoryImpl) FindByIDs(ids []int) ([]*entity.AreaCategory, error) {
	var areaCategories []*entity.AreaCategory

	if err := r.DB.Where("id IN (?)", ids).Find(&areaCategories).Error; err != nil {
		return nil, errors.Wrapf(err, "Failed get areaCategories")
	}

	return areaCategories, nil
}

func (r *AreaCategoryQueryRepositoryImpl) FindAreaListByAreaGroup(areaGroup model.AreaGroup, limit int, excludeID []int) ([]*entity.AreaCategory, error) {
	var rows []*entity.AreaCategory
	q := r.buildQueryByLimitAndExcludeID(limit, excludeID)

	if err := q.Where("area_group = ? AND type = ?", areaGroup, model.AreaCategoryTypeArea).Find(&rows).Error; err != nil {
		return nil, errors.Wrapf(err, "failed to area list by areaGroup= %s", areaGroup)
	}

	return rows, nil
}

func (r *AreaCategoryQueryRepositoryImpl) FindSubAreaListByAreaID(areaID int, limit int, excludeID []int) ([]*entity.AreaCategory, error) {
	var rows []*entity.AreaCategory
	q := r.buildQueryByLimitAndExcludeID(limit, excludeID)

	if err := q.Where("area_id = ? AND type = ?", areaID, model.AreaCategoryTypeSubArea).Find(&rows).Error; err != nil {
		return nil, errors.Wrapf(err, "failed to sub_area list by area_id = %d", areaID)
	}

	return rows, nil
}

func (r *AreaCategoryQueryRepositoryImpl) FindSubSubAreaListBySubAreaID(subAreaID int, limit int, excludeID []int) ([]*entity.AreaCategory, error) {
	var rows []*entity.AreaCategory
	q := r.buildQueryByLimitAndExcludeID(limit, excludeID)

	if err := q.Where("sub_area_id = ? AND type = ?", subAreaID, model.AreaCategoryTypeSubSubArea).Find(&rows).Error; err != nil {
		return nil, errors.Wrapf(err, "failed to sub_sub_area list by sub_area_id = %d", subAreaID)
	}

	return rows, nil
}

// limitやexcludeIDを指定する場合のqueryBuilder
func (r *AreaCategoryQueryRepositoryImpl) buildQueryByLimitAndExcludeID(limit int, excludeID []int) *gorm.DB {
	if limit == 0 {
		limit = defaultAcquisitionNumber
	}

	query := r.DB.Limit(limit)

	if len(excludeID) > 0 {
		query = query.Not("id", excludeID)
	}

	return query
}

func (r *AreaCategoryQueryRepositoryImpl) SearchByName(name string) ([]*entity.AreaCategory, error) {
	var rows []*entity.AreaCategory

	if err := r.DB.Where("MATCH(name) AGAINST(?)", name).Limit(defaultSearchSuggestionsNumber).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find areaCategory list by like name")
	}

	return rows, nil
}

func (r *AreaCategoryQueryRepositoryImpl) SearchAreaByName(name string) ([]*entity.AreaCategory, error) {
	var rows []*entity.AreaCategory

	q := r.buildSearchByNameAndType(name, model.AreaCategoryTypeArea)

	if err := q.Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to search area by name")
	}

	return rows, nil
}

func (r *AreaCategoryQueryRepositoryImpl) SearchSubAreaByName(name string) ([]*entity.AreaCategory, error) {
	var rows []*entity.AreaCategory

	q := r.buildSearchByNameAndType(name, model.AreaCategoryTypeSubArea)

	if err := q.Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to search area by name")
	}

	return rows, nil
}

func (r *AreaCategoryQueryRepositoryImpl) SearchSubSubAreaByName(name string) ([]*entity.AreaCategory, error) {
	var rows []*entity.AreaCategory

	q := r.buildSearchByNameAndType(name, model.AreaCategoryTypeSubSubArea)

	if err := q.Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to search area by name")
	}

	return rows, nil
}

func (r *AreaCategoryQueryRepositoryImpl) FindByMetaSearchID(innAreaTypeIDs *entity.InnAreaTypeIDs) ([]*entity.AreaCategory, error) {
	var rows []*entity.AreaCategory

	if err := r.DB.Where("metasearch_area_id = ?", innAreaTypeIDs.AreaID).Or("metasearch_sub_area_id =?", innAreaTypeIDs.SubAreaID).Or("metasearch_sub_sub_area_id = ?", innAreaTypeIDs.SubSubAreaID).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed find areaCategory list by metasearch_id")
	}

	return rows, nil
}

func (r *AreaCategoryQueryRepositoryImpl) buildSearchByNameAndType(name string, areaCategoryType model.AreaCategoryType) *gorm.DB {
	return r.DB.Where("MATCH(name) AGAINST(?)", name).Where("type = ?", areaCategoryType).Limit(defaultSearchSuggestionsNumber)
}
