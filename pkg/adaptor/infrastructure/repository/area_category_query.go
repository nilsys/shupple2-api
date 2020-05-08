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

const sortOrder = "sort_order ASC"

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

	if err := r.DB.Where("id IN (?)", ids).Order(sortOrder).Find(&areaCategories).Error; err != nil {
		return nil, errors.Wrapf(err, "Failed get areaCategories")
	}

	return areaCategories, nil
}

func (r *AreaCategoryQueryRepositoryImpl) FindAreaListHavingPostByAreaGroup(areaGroup model.AreaGroup, limit int, excludeID []int) ([]*entity.AreaCategoryWithPostCount, error) {
	var rows []*entity.AreaCategoryWithPostCount
	q := r.buildQueryByLimitAndExcludeID(limit, excludeID)

	/*
		area_category AS area_category, ac
		post_area_category AS pac

		SELECT area_category.*, post_count FROM `area_category` JOIN(
				SELECT ac.area_id, count(ac.id) AS post_count FROM post_area_category pac
				JOIN area_category ac ON pac.area_category_id = ac.id
				GROUP BY ac.area_id
		) a ON a.area_id = area_category.id
		WHERE (area_category.area_group = 'Japan' AND area_category.type = 'Area')
		HAVING (post_count > 0)
		ORDER BY post_count DESC LIMIT 1000
	*/
	if err := q.
		Select("area_category.*, post_count").
		Joins(`JOIN(
			SELECT ac.area_id, count(ac.id) AS post_count FROM post_area_category pac
			JOIN area_category ac ON pac.area_category_id = ac.id
			GROUP BY ac.area_id) a ON a.area_id = area_category.id`).
		Where("area_category.area_group = ? AND area_category.type = ?", areaGroup, model.AreaCategoryTypeArea).
		Having("post_count > 0").
		Order(sortOrder).
		Order("post_count DESC").
		Find(&rows).Error; err != nil {
		return nil, errors.Wrapf(err, "failed to area list by areaGroup= %s", areaGroup)
	}

	return rows, nil
}

func (r *AreaCategoryQueryRepositoryImpl) FindSubAreaListHavingPostByAreaIDAndThemeID(areaID, themeID, limit int, excludeID []int) ([]*entity.AreaCategoryWithPostCount, error) {
	var rows []*entity.AreaCategoryWithPostCount
	q := r.buildQueryByLimitAndExcludeID(limit, excludeID)

	if themeID != 0 {
		/*
			area_category AS area_category, ac
			post_area_category AS pac
			post_theme_category AS ptc

			SELECT area_category.*, post_count FROM `area_category` JOIN(
				SELECT ac.sub_area_id, count(ac.id) AS post_count FROM area_category ac JOIN (
					SELECT pac.* FROM post_theme_category ptc
					JOIN post_area_category pac ON pac.post_id = ptc.post_id
					WHERE ptc.theme_category_id = '1'
				) a ON ac.id = a.area_category_id
				WHERE (ac.type = 'SubArea' OR ac.type = 'SubSubArea')
				GROUP BY ac.sub_area_id
			)b ON area_category.id = b.sub_area_id
			WHERE (area_category.area_id = '1' AND area_category.type = 'SubArea')
			HAVING (post_count > 0)
			ORDER BY post_count DESC LIMIT 1000
		*/
		if err := q.
			Select("area_category.*, post_count").
			Joins(`JOIN(
				SELECT ac.sub_area_id, count(ac.id) AS post_count FROM area_category ac JOIN (
					SELECT pac.* FROM post_theme_category ptc
					JOIN post_area_category pac ON pac.post_id = ptc.post_id
					WHERE ptc.theme_category_id = ?
				) a ON ac.id = a.area_category_id
				WHERE (ac.type = ? OR ac.type = ?)
				GROUP BY ac.sub_area_id)b ON area_category.id = b.sub_area_id`, themeID, model.AreaCategoryTypeSubArea, model.AreaCategoryTypeSubSubArea).
			Where("area_category.area_id = ? AND area_category.type = ?", areaID, model.AreaCategoryTypeSubArea).
			Having("post_count > 0").
			Order(sortOrder).
			Order("post_count DESC").
			Find(&rows).Error; err != nil {
			return nil, errors.Wrapf(err, "failed to sub_area list by area_id = %d", areaID)
		}
		return rows, nil
	}
	/*
		area_category AS area_category, ac
		post_area_category AS pac

		SELECT area_category.*, post_count FROM `area_category` JOIN(
			SELECT ac.sub_area_id, count(ac.id) AS post_count FROM post_area_category pac
			JOIN area_category ac ON pac.area_category_id = ac.id
			WHERE (ac.type = 'SubArea' OR ac.type = 'SubSubArea')
			GROUP BY ac.sub_area_id
		)a ON area_category.id = a.sub_area_id
		WHERE (area_category.area_id = '7' AND area_category.type = 'SubArea')
		HAVING (post_count > 0)
		ORDER BY post_count DESC LIMIT 1000
	*/
	if err := q.
		Select("area_category.*, post_count").
		Joins(`JOIN(
			SELECT ac.sub_area_id, count(ac.id) AS post_count FROM post_area_category pac
			JOIN area_category ac ON pac.area_category_id = ac.id
			WHERE (ac.type = ? OR ac.type = ?)
			GROUP BY ac.sub_area_id)a ON area_category.id = a.sub_area_id`, model.AreaCategoryTypeSubArea, model.AreaCategoryTypeSubSubArea).
		Where("area_category.area_id = ? AND area_category.type = ?", areaID, model.AreaCategoryTypeSubArea).
		Having("post_count > 0").
		Order(sortOrder).
		Order("post_count DESC").
		Find(&rows).Error; err != nil {
		return nil, errors.Wrapf(err, "failed to sub_area list by area_id = %d", areaID)
	}
	return rows, nil
}

func (r *AreaCategoryQueryRepositoryImpl) FindSubSubAreaListHavingPostBySubAreaIDAndThemeID(subAreaID, themeID int, limit int, excludeID []int) ([]*entity.AreaCategoryWithPostCount, error) {
	var rows []*entity.AreaCategoryWithPostCount
	q := r.buildQueryByLimitAndExcludeID(limit, excludeID)

	if themeID != 0 {
		/*
			area_category AS area_category
			post_area_category AS pac
			post_theme_category AS ptc

			SELECT area_category.*, count(area_category.id) AS post_count FROM `area_category` JOIN(
				SELECT pac.* FROM post_theme_category ptc
				JOIN post_area_category pac ON pac.post_id = ptc.post_id
				WHERE ptc.theme_category_id = '1'
			)a ON area_category.sub_sub_area_id = a.area_category_id
			WHERE (sub_area_id = '10' AND type = 'SubSubArea')
			GROUP BY area_category.id
			HAVING (post_count > 0)
			ORDER BY post_count DESC LIMIT 1000
		*/
		if err := q.
			Select("area_category.*, count(area_category.id) AS post_count").
			Joins(`JOIN(
					SELECT pac.* FROM post_theme_category ptc
					JOIN post_area_category pac ON pac.post_id = ptc.post_id
					WHERE ptc.theme_category_id = ?
				)a ON area_category.sub_sub_area_id = a.area_category_id`, themeID).
			Where("sub_area_id = ? AND type = ?", subAreaID, model.AreaCategoryTypeSubSubArea).
			Group("area_category.id").
			Having("post_count > 0").
			Order(sortOrder).
			Order("post_count DESC").
			Find(&rows).Error; err != nil {
			return nil, errors.Wrapf(err, "failed to sub_sub_area list by sub_area_id = %d", subAreaID)
		}
		return rows, nil
	}
	/*
		area_category AS area_category
		post_area_category AS pac

		SELECT area_category.*, count(pac.post_id) as post_count FROM `area_category`
		JOIN post_area_category pac ON area_category.id = pac.area_category_id
		WHERE (sub_area_id = '1' AND type = 'SubSubArea')
		GROUP BY area_category.id
		HAVING post_count > 0
		ORDER BY post_count DESC LIMIT 1000
	*/
	if err := q.
		Select("area_category.*, count(pac.post_id) as post_count").
		Joins("JOIN post_area_category pac ON area_category.id = pac.area_category_id").
		Where("sub_area_id = ? AND type = ?", subAreaID, model.AreaCategoryTypeSubSubArea).
		Group("area_category.id").
		Having("post_count > 0").
		Order(sortOrder).
		Order("post_count DESC").
		Find(&rows).Error; err != nil {
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

	if err := r.DB.Where("MATCH(name) AGAINST(?)", name).Limit(defaultSearchSuggestionsNumber).Order(sortOrder).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find areaCategory list by like name")
	}

	return rows, nil
}

func (r *AreaCategoryQueryRepositoryImpl) SearchAreaByName(name string) ([]*entity.AreaCategory, error) {
	var rows []*entity.AreaCategory

	q := r.buildSearchByNameAndType(name, model.AreaCategoryTypeArea)

	if err := q.Order(sortOrder).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to search area by name")
	}

	return rows, nil
}

func (r *AreaCategoryQueryRepositoryImpl) SearchSubAreaByName(name string) ([]*entity.AreaCategory, error) {
	var rows []*entity.AreaCategory

	q := r.buildSearchByNameAndType(name, model.AreaCategoryTypeSubArea)

	if err := q.Order(sortOrder).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to search area by name")
	}

	return rows, nil
}

func (r *AreaCategoryQueryRepositoryImpl) SearchSubSubAreaByName(name string) ([]*entity.AreaCategory, error) {
	var rows []*entity.AreaCategory

	q := r.buildSearchByNameAndType(name, model.AreaCategoryTypeSubSubArea)

	if err := q.Order(sortOrder).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to search area by name")
	}

	return rows, nil
}

func (r *AreaCategoryQueryRepositoryImpl) FindByMetaSearchID(innAreaTypeIDs *entity.InnAreaTypeIDs) ([]*entity.AreaCategory, error) {
	var rows []*entity.AreaCategory

	if err := r.DB.Order(sortOrder).Where("metasearch_area_id = ?", innAreaTypeIDs.AreaID).Or("metasearch_sub_area_id =?", innAreaTypeIDs.SubAreaID).Or("metasearch_sub_sub_area_id = ?", innAreaTypeIDs.SubSubAreaID).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed find areaCategory list by metasearch_id")
	}

	return rows, nil
}

func (r *AreaCategoryQueryRepositoryImpl) buildSearchByNameAndType(name string, areaCategoryType model.AreaCategoryType) *gorm.DB {
	return r.DB.Where("MATCH(name) AGAINST(?)", name).Where("type = ?", areaCategoryType).Limit(defaultSearchSuggestionsNumber)
}
