package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type CategoryQueryRepositoryImpl struct {
	DB *gorm.DB
}

var CategoryQueryRepositorySet = wire.NewSet(
	wire.Struct(new(CategoryQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.CategoryQueryRepository), new(*CategoryQueryRepositoryImpl)),
)

func (r *CategoryQueryRepositoryImpl) FindTypeByID(id int) (*model.CategoryType, error) {
	var category entity.Category

	if err := r.DB.First(&category, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "failed to find category_type by parentCategoryID")
	}

	return &category.Type, nil
}

func (r *CategoryQueryRepositoryImpl) FindListByParentCategoryID(parentCategoryID int, limit int, excludeID []int) ([]*entity.Category, error) {
	var rows []*entity.Category
	q := r.buildFindListByParentCategoryID(limit, excludeID)

	if err := q.Where("parent_id = ?", parentCategoryID).Find(&rows).Error; err != nil {
		return nil, errors.Wrapf(err, "category not found")
	}

	return rows, nil
}

// FindListByParentCategoryID関数内でのqueryBuilder
func (r *CategoryQueryRepositoryImpl) buildFindListByParentCategoryID(limit int, excludeID []int) *gorm.DB {
	if limit == 0 {
		limit = defaultAcquisitionNumber
	}

	query := r.DB.Limit(limit)

	if len(excludeID) > 0 {
		query = query.Not("id", excludeID)
	}

	return query
}

func (r *CategoryQueryRepositoryImpl) FindByIDs(ids []int) ([]*entity.Category, error) {
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

func (r *CategoryQueryRepositoryImpl) SearchByName(name string) ([]*entity.Category, error) {
	var rows []*entity.Category

	if err := r.DB.Where("MATCH(name) AGAINST(?)", name).Not("type = ?", model.CategoryTypeAreaGroup).Limit(defaultSearchSuggestionsNumber).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find category list by like name")
	}

	return rows, nil
}

func (r *CategoryQueryRepositoryImpl) FindByTouristSpotID(touristSpotID int) ([]*entity.Category, error) {
	var rows []*entity.Category

	if err := r.DB.Where("id IN (SELECT category_id FROM tourist_spot_category WHERE tourist_spot_id = ?)", touristSpotID).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed find category list by tourist_spot_id")
	}

	return rows, nil
}

func (r *CategoryQueryRepositoryImpl) SearchAreaByName(name string) ([]*entity.Category, error) {
	var rows []*entity.Category

	q := r.buildSearchByNameAndType(name, model.CategoryTypeArea)

	if err := q.Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to search area by name")
	}

	return rows, nil
}

func (r *CategoryQueryRepositoryImpl) SearchSubAreaByName(name string) ([]*entity.Category, error) {
	var rows []*entity.Category

	q := r.buildSearchByNameAndType(name, model.CategoryTypeSubArea)

	if err := q.Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to search area by name")
	}

	return rows, nil
}

func (r *CategoryQueryRepositoryImpl) FindByMetaSearchID(innAreaTypeIDs *entity.InnAreaTypeIDs) ([]*entity.Category, error) {
	var rows []*entity.Category

	if err := r.DB.Where("metasearch_area_id = ?", innAreaTypeIDs.AreaID).Or("metasearch_sub_area_id =?", innAreaTypeIDs.SubAreaID).Or("metasearch_sub_sub_area_id = ?", innAreaTypeIDs.SubSubAreaID).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed find category list by metasearch_id")
	}

	return rows, nil
}

func (r *CategoryQueryRepositoryImpl) SearchSubSubAreaByName(name string) ([]*entity.Category, error) {
	var rows []*entity.Category

	q := r.buildSearchByNameAndType(name, model.CategoryTypeSubSubArea)

	if err := q.Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to search area by name")
	}

	return rows, nil
}

func (r *CategoryQueryRepositoryImpl) buildSearchByNameAndType(name string, categoryType model.CategoryType) *gorm.DB {
	return r.DB.Where("MATCH(name) AGAINST(?)", name).Where("type = ?", categoryType).Limit(defaultSearchSuggestionsNumber)
}
