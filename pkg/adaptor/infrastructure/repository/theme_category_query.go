package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type ThemeCategoryQueryRepositoryImpl struct {
	DB *gorm.DB
}

var ThemeCategoryQueryRepositorySet = wire.NewSet(
	wire.Struct(new(ThemeCategoryQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.ThemeCategoryQueryRepository), new(*ThemeCategoryQueryRepositoryImpl)),
)

func (r *ThemeCategoryQueryRepositoryImpl) FindByID(id int) (*entity.ThemeCategory, error) {
	var row entity.ThemeCategory
	if err := r.DB.First(&row, id).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "themeCategory(id=%d)", id)
	}
	return &row, nil
}

func (r *ThemeCategoryQueryRepositoryImpl) FindBySlug(slug string) (*entity.ThemeCategory, error) {
	var row entity.ThemeCategory
	if err := r.DB.Where("slug = ?", slug).First(&row).Error; err != nil {
		return nil, ErrorToFindSingleRecord(err, "themeCategory(slug=%s)", slug)
	}
	return &row, nil
}

func (r *ThemeCategoryQueryRepositoryImpl) FindByIDs(ids []int) ([]*entity.ThemeCategory, error) {
	var themeCategories []*entity.ThemeCategory

	if err := r.DB.Where("id IN (?)", ids).Find(&themeCategories).Error; err != nil {
		return nil, errors.Wrapf(err, "Failed get themeCategories")
	}

	return themeCategories, nil
}

func (r *ThemeCategoryQueryRepositoryImpl) SearchByName(name string) ([]*entity.ThemeCategory, error) {
	var rows []*entity.ThemeCategory

	if err := r.DB.Where("MATCH(name) AGAINST(?)", name).Limit(defaultSearchSuggestionsNumber).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find themeCategory list by like name")
	}

	return rows, nil
}

func (r *ThemeCategoryQueryRepositoryImpl) FindThemesByAreaCategoryID(areaID, subAreaID, subSubAreaID int, excludeIDs []int) ([]*entity.ThemeCategoryWithPostCount, error) {
	var rows []*entity.ThemeCategoryWithPostCount

	if err := r.buildFindThemesByAreaCategoryIDQuery(areaID, subAreaID, subSubAreaID, excludeIDs).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find theme list")
	}

	return rows, nil
}

func (r *ThemeCategoryQueryRepositoryImpl) FindSubThemesByAreaCategoryIDAndParentThemeID(parentThemeID, areaID, subAreaID, subSubAreaID int, excludeIDs []int) ([]*entity.ThemeCategoryWithPostCount, error) {
	var rows []*entity.ThemeCategoryWithPostCount

	if err := r.buildFindSubThemesByAreaCategoryIDAndParentThemeIDQuery(parentThemeID, areaID, subAreaID, subSubAreaID, excludeIDs).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find subTheme list")
	}

	return rows, nil
}

func (r *ThemeCategoryQueryRepositoryImpl) buildQueryByExcludeIDs(excludeIDs []int) *gorm.DB {
	if len(excludeIDs) > 0 {
		return r.DB.Not("id", excludeIDs)
	}
	return r.DB
}

func (r *ThemeCategoryQueryRepositoryImpl) buildFindThemesByAreaCategoryIDQuery(areaID, subAreaID, subSubAreaID int, excludeIDs []int) *gorm.DB {
	q := r.buildQueryByExcludeIDs(excludeIDs).Limit(defaultAcquisitionNumber)

	if areaID != 0 {
		/*
				theme_category AS theme_category, tc
				area_category AS ac
				post_area_category AS pac
				post_theme_category AS ptc

			SELECT theme_category.*, post_count FROM `theme_category` JOIN (
				SELECT tc.theme_id, count(tc.theme_id) AS post_count FROM theme_category tc JOIN(
					SELECT ptc.theme_category_id FROM post_theme_category ptc JOIN(
						SELECT pac.post_id FROM post_area_category pac
						JOIN area_category ac ON ac.id = pac.area_category_id
						WHERE ac.area_id = '1'
						GROUP BY pac.post_id
					)a ON a.post_id = ptc.post_id
				)b ON b.theme_category_id = tc.id
				GROUP BY tc.theme_id
			)c ON c.theme_id = theme_category.id
			ORDER BY post_count DESC
			LIMIT 1000
		*/
		return q.
			Select("theme_category.*, post_count").
			Joins(`JOIN (
				SELECT tc.theme_id, count(tc.theme_id) AS post_count FROM theme_category tc JOIN(
					SELECT ptc.theme_category_id FROM post_theme_category ptc JOIN(
						SELECT pac.post_id FROM post_area_category pac
						JOIN area_category ac ON ac.id = pac.area_category_id
						WHERE ac.area_id = ?
						GROUP BY pac.post_id
					)a ON a.post_id = ptc.post_id
				)b ON b.theme_category_id = tc.id
				GROUP BY tc.theme_id
			)c ON c.theme_id = theme_category.id`, areaID).
			Order("post_count DESC")
	}

	if subAreaID != 0 {
		/*
			theme_category AS theme_category, tc
			area_category AS ac
			post_area_category AS pac
			post_theme_category AS ptc

			SELECT theme_category.*, post_count FROM theme_category JOIN (
				SELECT tc.theme_id, count(tc.theme_id) AS post_count FROM theme_category tc JOIN(
					SELECT ptc.theme_category_id FROM post_theme_category ptc JOIN(
						SELECT pac.post_id FROM post_area_category pac
						JOIN area_category ac ON ac.id = pac.area_category_id
						WHERE ac.sub_area_id = '10'
						GROUP BY pac.post_id
					)a ON a.post_id = ptc.post_id
				)b ON b.theme_category_id = tc.id
				GROUP BY tc.theme_id
			)c ON c.theme_id = theme_category.id
			ORDER BY post_count DESC
			LIMIT 1000
		*/
		return q.
			Select("theme_category.*, post_count").
			Joins(`JOIN (
					SELECT tc.theme_id, count(tc.theme_id) AS post_count FROM theme_category tc JOIN(
						SELECT ptc.theme_category_id FROM post_theme_category ptc JOIN(
							SELECT pac.post_id FROM post_area_category pac
							JOIN area_category ac ON ac.id = pac.area_category_id
							WHERE ac.sub_area_id = ?
							GROUP BY pac.post_id
						)a ON a.post_id = ptc.post_id
					)b ON b.theme_category_id = tc.id
					GROUP BY tc.theme_id
				)c ON c.theme_id = theme_category.id`, subAreaID).
			Order("post_count DESC")
	}

	if subSubAreaID != 0 {
		/*
			theme_category AS theme_category, tc
			area_category AS ac
			post_area_category AS pac
			post_theme_category AS ptc

			SELECT theme_category.*, post_count FROM theme_category JOIN (
				SELECT tc.theme_id, count(tc.theme_id) AS post_count FROM theme_category tc JOIN(
					SELECT ptc.theme_category_id FROM post_theme_category ptc JOIN(
						SELECT pac.post_id FROM post_area_category pac
						JOIN area_category ac ON ac.id = pac.area_category_id
						WHERE ac.sub_sub_area_id = '12'
						GROUP BY pac.post_id
					)a ON a.post_id = ptc.post_id
				)b ON b.theme_category_id = tc.id
				GROUP BY tc.theme_id
			)c ON c.theme_id = theme_category.id
			ORDER BY post_count DESC
			LIMIT 1000
		*/
		return q.
			Select("theme_category.*, post_count").
			Joins(`JOIN (
					SELECT tc.theme_id, count(tc.theme_id) AS post_count FROM theme_category tc JOIN(
						SELECT ptc.theme_category_id FROM post_theme_category ptc JOIN(
							SELECT pac.post_id FROM post_area_category pac
							JOIN area_category ac ON ac.id = pac.area_category_id
							WHERE ac.sub_sub_area_id = ?
							GROUP BY pac.post_id
						)a ON a.post_id = ptc.post_id
					)b ON b.theme_category_id = tc.id
					GROUP BY tc.theme_id
				)c ON c.theme_id = theme_category.id`, subSubAreaID).
			Order("post_count DESC")
	}

	/*
		theme_category AS theme_category,tc
		post_theme_category AS ptc

		SELECT theme_category.*, post_count FROM `theme_category` JOIN (
			SELECT tc.theme_id AS id, count(tc.theme_id) AS post_count
			FROM theme_category tc
			JOIN post_theme_category ptc ON tc.id = ptc.theme_category_id
			GROUP BY tc.theme_id
		) tid ON theme_category.id = tid.id
		ORDER BY post_count DESC
		LIMIT 1000
	*/
	return q.
		Select("theme_category.*, post_count").
		Joins(`JOIN (
				SELECT tc.theme_id AS id, count(tc.theme_id) AS post_count 
				FROM theme_category tc 
				JOIN post_theme_category ptc ON tc.id = ptc.theme_category_id 
				GROUP BY tc.theme_id
			) tid ON theme_category.id = tid.id`).
		Order("post_count DESC")
}

func (r *ThemeCategoryQueryRepositoryImpl) buildFindSubThemesByAreaCategoryIDAndParentThemeIDQuery(parentThemeID, areaID, subAreaID, subSubAreaID int, excludeIDs []int) *gorm.DB {
	q := r.buildQueryByExcludeIDs(excludeIDs).Limit(defaultAcquisitionNumber)

	if areaID != 0 {
		/*
			theme_category AS theme_category, tc
			area_category AS ac
			post_area_category AS pac
			post_theme_category AS ptc

			SELECT theme_category.*, post_count FROM `theme_category` JOIN (
				SELECT tc.sub_theme_id, count(tc.sub_theme_id) AS post_count FROM theme_category tc JOIN(
					SELECT ptc.theme_category_id FROM post_theme_category ptc JOIN(
						SELECT pac.post_id FROM post_area_category pac
						JOIN area_category ac ON ac.id = pac.area_category_id
						WHERE ac.area_id = '1'
						GROUP BY pac.post_id
					)a ON a.post_id = ptc.post_id
				)b ON b.theme_category_id = tc.id
				GROUP BY tc.sub_theme_id
			)c ON c.sub_theme_id = theme_category.id
			WHERE (theme_category.theme_id = '1')
			ORDER BY post_count DESC
			LIMIT 1000
		*/
		return q.
			Select("theme_category.*, post_count").
			Joins(`JOIN (
					SELECT tc.sub_theme_id, count(tc.sub_theme_id) AS post_count FROM theme_category tc JOIN(
						SELECT ptc.theme_category_id FROM post_theme_category ptc JOIN(
							SELECT pac.post_id FROM post_area_category pac
							JOIN area_category ac ON ac.id = pac.area_category_id
							WHERE ac.area_id = ?
							GROUP BY pac.post_id
						)a ON a.post_id = ptc.post_id
					)b ON b.theme_category_id = tc.id
					GROUP BY tc.sub_theme_id
				)c ON c.sub_theme_id = theme_category.id`, areaID).
			Where("theme_category.theme_id = ?", parentThemeID).
			Order("post_count DESC")
	}

	if subAreaID != 0 {
		/*
			theme_category AS theme_category, tc
			area_category AS ac
			post_area_category AS pac
			post_theme_category AS ptc

			SELECT theme_category.*, post_count FROM theme_category JOIN (
				SELECT tc.sub_theme_id, count(tc.sub_theme_id) AS post_count FROM theme_category tc JOIN(
					SELECT ptc.theme_category_id FROM post_theme_category ptc JOIN(
						SELECT pac.post_id FROM post_area_category pac
						JOIN area_category ac ON ac.id = pac.area_category_id
						WHERE ac.sub_area_id = '10'
						GROUP BY pac.post_id
					)a ON a.post_id = ptc.post_id
				)b ON b.theme_category_id = tc.id
				GROUP BY tc.sub_theme_id
			)c ON c.sub_theme_id = theme_category.id
			WHERE theme_category.theme_id = '1'
			ORDER BY post_count DESC
			LIMIT 1000
		*/
		return q.
			Select("theme_category.*, post_count").
			Joins(`JOIN (
					SELECT tc.sub_theme_id, count(tc.sub_theme_id) AS post_count FROM theme_category tc JOIN(
						SELECT ptc.theme_category_id FROM post_theme_category ptc JOIN(
							SELECT pac.post_id FROM post_area_category pac
							JOIN area_category ac ON ac.id = pac.area_category_id
							WHERE ac.sub_area_id = ?
							GROUP BY pac.post_id
						)a ON a.post_id = ptc.post_id
					)b ON b.theme_category_id = tc.id
					GROUP BY tc.sub_theme_id
				)c ON c.sub_theme_id = theme_category.id`, subAreaID).
			Where("theme_category.theme_id = ?", parentThemeID).
			Order("post_count DESC")
	}

	if subSubAreaID != 0 {
		/*
			theme_category AS theme_category, tc
			area_category AS ac
			post_area_category AS pac
			post_theme_category AS ptc

			SELECT theme_category.*, post_count FROM theme_category JOIN (
				SELECT tc.sub_theme_id, count(tc.sub_theme_id) AS post_count FROM theme_category tc JOIN(
					SELECT ptc.theme_category_id FROM post_theme_category ptc JOIN(
						SELECT pac.post_id FROM post_area_category pac
						JOIN area_category ac ON ac.id = pac.area_category_id
						WHERE ac.sub_sub_area_id = '12'
						GROUP BY pac.post_id
					)a ON a.post_id = ptc.post_id
				)b ON b.theme_category_id = tc.id
				GROUP BY tc.sub_theme_id
			)c ON c.sub_theme_id = theme_category.id
			WHERE theme_category.theme_id = '1'
			ORDER BY post_count DESC
			LIMIT 1000
		*/
		return q.
			Select("theme_category.*, post_count").
			Joins(`JOIN (
					SELECT tc.sub_theme_id, count(tc.sub_theme_id) AS post_count FROM theme_category tc JOIN(
						SELECT ptc.theme_category_id FROM post_theme_category ptc JOIN(
							SELECT pac.post_id FROM post_area_category pac
							JOIN area_category ac ON ac.id = pac.area_category_id
							WHERE ac.sub_sub_area_id = ?
							GROUP BY pac.post_id
						)a ON a.post_id = ptc.post_id
					)b ON b.theme_category_id = tc.id
					GROUP BY tc.sub_theme_id
				)c ON c.sub_theme_id = theme_category.id`, subSubAreaID).
			Where("theme_category.theme_id = ?", parentThemeID).
			Order("post_count DESC")
	}

	/*
		theme_category AS theme_category
		post_theme_category AS ptc

		SELECT `theme_category`.*, post_count FROM `theme_category` JOIN (
		       SELECT ptc.theme_category_id AS id, count(ptc.theme_category_id) AS post_count FROM post_theme_category ptc
		       GROUP BY ptc.theme_category_id
		) tid ON theme_category.id = tid.id
		WHERE (theme_category.type = 'SubTheme' AND theme_category.theme_id = '1')
		ORDER BY post_count DESC LIMIT 1000
	*/
	// TODO:パフォーマンス調べる
	return q.
		Select("theme_category.*, post_count").
		Joins(`JOIN (
			SELECT ptc.theme_category_id AS id, count(ptc.theme_category_id) AS post_count 
			FROM post_theme_category ptc 
			GROUP BY ptc.theme_category_id
			) tid ON theme_category.id = tid.id`).
		Where("theme_category.type = ? AND theme_category.theme_id = ?", model.ThemeCategoryTypeSubTheme, parentThemeID).
		Order("post_count DESC")
}
