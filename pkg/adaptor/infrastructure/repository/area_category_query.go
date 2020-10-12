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
			SELECT area_category.*, post_count FROM `area_category` JOIN(
		            SELECT area_id, count(*) AS post_count FROM post JOIN (
		                    SELECT post_id, area_id FROM post_area_category JOIN (
		                            SELECT id, area_id FROM area_category
		                            WHERE area_group = '1'
		                    ) at ON post_area_category.area_category_id = at.id
							GROUP BY post_id, area_id
		            ) pa ON post.id = pa.post_id
		            GROUP BY area_id
		    ) ac ON area_category.area_id = ac.area_id
			WHERE (type = 'Area') HAVING (post_count > 0) ORDER BY post_count DESC,sort_order ASC LIMIT 1000
	*/
	if err := q.
		Select("area_category.*, post_count").
		Joins(`JOIN(
			SELECT area_id, count(*) AS post_count FROM post JOIN (
				SELECT post_id, area_id FROM post_area_category JOIN (
					SELECT id, area_id FROM area_category
					WHERE area_group = ?
                ) at ON post_area_category.area_category_id = at.id
				GROUP BY post_id, area_id
            ) pa ON post.id = pa.post_id
			GROUP BY area_id
		) ac ON area_category.area_id = ac.area_id`, areaGroup).
		Where("type = ?", model.AreaCategoryTypeArea).
		Having("post_count > 0").
		Order(sortOrder).
		Order("post_count DESC").
		Find(&rows).Error; err != nil {
		return nil, errors.Wrapf(err, "failed to area list by areaGroup = %s", areaGroup)
	}

	return rows, nil
}

func (r *AreaCategoryQueryRepositoryImpl) FindSubAreaListHavingPostByAreaIDAndThemeID(areaID, themeID, limit int, excludeID []int) ([]*entity.AreaCategoryWithPostCount, error) {
	var rows []*entity.AreaCategoryWithPostCount
	q := r.buildQueryByLimitAndExcludeID(limit, excludeID)

	if themeID != 0 {
		/*
				SELECT area_category.*, post_count FROM `area_category` JOIN(
			            SELECT sub_area_id, count(*) AS post_count FROM post JOIN (
			                    SELECT post_id, sub_area_id FROM post_area_category JOIN (
			                            SELECT id, sub_area_id FROM area_category
			                            WHERE area_id = '418'
			                    ) at ON post_area_category.area_category_id = at.id
								GROUP BY post_id, sub_area_id
			            ) ps ON post.id = ps.post_id
						WHERE post.id IN (
								SELECT post_id FROM post_theme_category WHERE theme_category_id IN (
									SELECT id FROM theme_category
									WHERE theme_id = '1699'
								)
						)
			            GROUP BY sub_area_id
			    ) sc ON area_category.sub_area_id = sc.sub_area_id
				WHERE (type = '2') HAVING (post_count > 0) ORDER BY post_count DESC,sort_order ASC LIMIT 1000
		*/
		if err := q.
			Select("area_category.*, post_count").
			Joins(`JOIN(
               SELECT sub_area_id, count(*) AS post_count FROM post JOIN (
                       SELECT post_id, sub_area_id FROM post_area_category JOIN (
                               SELECT id, sub_area_id FROM area_category
                               WHERE area_id = ?
			           ) at ON post_area_category.area_category_id = at.id
		    			GROUP BY post_id, sub_area_id
				) ps ON post.id = ps.post_id
		    	WHERE post.id IN (
		    			SELECT post_id FROM post_theme_category WHERE theme_category_id IN (
		    				SELECT id FROM theme_category
		    				WHERE theme_id = ?
		    			)
		    	)
               GROUP BY sub_area_id
			) sc ON area_category.sub_area_id = sc.sub_area_id`, areaID, themeID).
			Where("type = ?", model.AreaCategoryTypeSubArea).
			Having("post_count > 0").
			Order(sortOrder).
			Order("post_count DESC").
			Find(&rows).Error; err != nil {
			return nil, errors.Wrapf(err, "failed to sub_area list by area_id = %d", areaID)
		}
		return rows, nil
	}
	/*
			SELECT area_category.*, post_count FROM `area_category` JOIN(
		            SELECT sub_area_id, count(*) AS post_count FROM post JOIN (
		                    SELECT post_id, sub_area_id FROM post_area_category JOIN (
		                            SELECT id, sub_area_id FROM area_category
		                            WHERE area_id = '418'
		                    ) at ON post_area_category.area_category_id = at.id
							GROUP BY post_id, sub_area_id
		            ) ps ON post.id = ps.post_id
		            GROUP BY sub_area_id
		    ) sc ON area_category.sub_area_id = sc.sub_area_id
			WHERE (type = 'SubArea') HAVING (post_count > 0) ORDER BY post_count DESC,sort_order ASC LIMIT 1000
	*/
	if err := q.
		Select("area_category.*, post_count").
		Joins(`JOIN(
			SELECT sub_area_id, count(*) AS post_count FROM post JOIN (
				SELECT post_id, sub_area_id FROM post_area_category JOIN (
					SELECT id, sub_area_id FROM area_category
					WHERE area_id = ?
				) at ON post_area_category.area_category_id = at.id
				GROUP BY post_id, sub_area_id
			) ps ON post.id = ps.post_id
			GROUP BY sub_area_id
		) sc ON area_category.sub_area_id = sc.sub_area_id`, areaID).
		Where("type = ?", model.AreaCategoryTypeSubArea).
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
				SELECT area_category.*, post_count FROM `area_category` JOIN(
			            SELECT sub_sub_area_id, count(*) AS post_count FROM post INNER JOIN (
			                    SELECT post_id, sub_sub_area_id FROM post_area_category JOIN (
			                            SELECT id, sub_sub_area_id FROM area_category
			                            WHERE sub_area_id = '416' AND type = '3'
			                    ) at ON post_area_category.area_category_id = at.id
								GROUP BY post_id, sub_sub_area_id
			            ) pss ON post.id = pss.post_id
						WHERE id IN (
								SELECT post_id FROM post_theme_category WHERE theme_category_id IN (
									SELECT id FROM theme_category
									WHERE theme_id = '1696'
								)
						)
			            GROUP BY sub_sub_area_id
			    ) ssc ON area_category.sub_sub_area_id = ssc.sub_sub_area_id
				HAVING (post_count > 0) ORDER BY post_count DESC,sort_order ASC LIMIT 1000
		*/
		if err := q.
			Select("area_category.*, post_count").
			Joins(`JOIN(
                SELECT sub_sub_area_id, count(*) AS post_count FROM post JOIN (
                        SELECT post_id, sub_sub_area_id FROM post_area_category JOIN (
                                SELECT id, sub_sub_area_id FROM area_category
                                WHERE sub_area_id = ? AND type = ?
                        ) at ON post_area_category.area_category_id = at.id
			    		GROUP BY post_id, sub_sub_area_id
                ) pss ON post.id = pss.post_id
			    WHERE id IN (
			    		SELECT post_id FROM post_theme_category WHERE theme_category_id IN (
			    			SELECT id FROM theme_category
			    			WHERE theme_id = ?
			    		)
			    )
                GROUP BY sub_sub_area_id
			) ssc ON area_category.sub_sub_area_id = ssc.sub_sub_area_id`, subAreaID, model.AreaCategoryTypeSubSubArea, themeID).
			Having("post_count > 0").
			Order(sortOrder).
			Order("post_count DESC").
			Find(&rows).Error; err != nil {
			return nil, errors.Wrapf(err, "failed to sub_sub_area list by sub_area_id = %d", subAreaID)
		}
		return rows, nil
	}
	/*
			SELECT area_category.*, post_count FROM `area_category` JOIN(
		            SELECT sub_sub_area_id, count(*) AS post_count FROM post JOIN (
		                    SELECT post_id, sub_sub_area_id FROM post_area_category JOIN (
		                            SELECT id, sub_sub_area_id FROM area_category
		                            WHERE sub_area_id = '441'
		                    ) at ON post_area_category.area_category_id = at.id
							GROUP BY post_id, sub_sub_area_id
		            ) pss ON post.id = pss.post_id
		            GROUP BY sub_sub_area_id
		    ) ssc ON area_category.sub_sub_area_id = ssc.sub_sub_area_id
			WHERE (type = 'SubSubArea') HAVING (post_count > 0) ORDER BY post_count DESC,sort_order ASC LIMIT 1000
	*/
	if err := q.
		Select("area_category.*, post_count").
		Joins(`JOIN (
           SELECT sub_sub_area_id, count(*) AS post_count FROM post JOIN (
                   SELECT post_id, sub_sub_area_id FROM post_area_category JOIN (
                           SELECT id, sub_sub_area_id FROM area_category
                           WHERE sub_area_id = ?
                   ) at ON post_area_category.area_category_id = at.id
					GROUP BY post_id, sub_sub_area_id
           ) pss ON post.id = pss.post_id
           GROUP BY sub_sub_area_id
		) ssc ON area_category.sub_sub_area_id = ssc.sub_sub_area_id`, subAreaID).
		Where("type = ?", model.AreaCategoryTypeSubSubArea).
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

func (r *AreaCategoryQueryRepositoryImpl) SearchByName(name string) (*entity.AreaCategories, error) {
	var rows entity.AreaCategories

	if err := r.DB.Where("MATCH(name) AGAINST(?)", name).Limit(defaultSearchSuggestionsNumber).Order(sortOrder).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find areaCategory list by like name")
	}

	return &rows, nil
}

func (r *AreaCategoryQueryRepositoryImpl) SearchAreaByName(name string) (*entity.AreaCategories, error) {
	var rows entity.AreaCategories

	q := r.buildSearchByNameAndType(name, model.AreaCategoryTypeArea)

	if err := q.Order(sortOrder).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to search area by name")
	}

	return &rows, nil
}

func (r *AreaCategoryQueryRepositoryImpl) SearchSubAreaByName(name string) (*entity.AreaCategories, error) {
	var rows entity.AreaCategories

	q := r.buildSearchByNameAndType(name, model.AreaCategoryTypeSubArea)

	if err := q.Order(sortOrder).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to search area by name")
	}

	return &rows, nil
}

func (r *AreaCategoryQueryRepositoryImpl) SearchSubSubAreaByName(name string) (*entity.AreaCategories, error) {
	var rows entity.AreaCategories

	q := r.buildSearchByNameAndType(name, model.AreaCategoryTypeSubSubArea)

	if err := q.Order(sortOrder).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to search area by name")
	}

	return &rows, nil
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
