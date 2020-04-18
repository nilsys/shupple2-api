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

func (r *ThemeCategoryQueryRepositoryImpl) FindAll(excludeIDs []int) ([]*entity.ThemeCategoryWithPostCount, error) {
	var rows []*entity.ThemeCategoryWithPostCount

	/*
		theme_category AS theme_category,tc
		post_theme_category AS ptc

		SELECT theme_category.*, post_count FROM `theme_category` JOIN (
			SELECT tc.theme_id AS id, count(tc.theme_id) AS post_count
			FROM theme_category tc
			JOIN post_theme_category ptc ON tc.id = ptc.theme_category_id
			GROUP BY tc.theme_id
		) tid ON theme_category.id = tid.id
		WHERE (theme_category.type = 'Theme')
		ORDER BY post_count DESC
		LIMIT 1000
	*/
	if err := r.buildQueryByExcludeIDs(excludeIDs).Limit(defaultAcquisitionNumber).
		Table("theme_category").
		Select("theme_category.*, post_count").
		Joins(`JOIN (
			SELECT tc.theme_id AS id, count(tc.theme_id) AS post_count 
			FROM theme_category tc 
			JOIN post_theme_category ptc ON tc.id = ptc.theme_category_id 
			GROUP BY tc.theme_id
			) tid ON theme_category.id = tid.id`).
		Where("theme_category.type = ?", model.ThemeCategoryTypeTheme).
		Order("post_count DESC").
		Find(&rows).Error; err != nil {

		return nil, errors.Wrap(err, "failed to find theme_category")
	}

	return rows, nil
}

func (r *ThemeCategoryQueryRepositoryImpl) FindThemesByAreaCategoryID(excludeIDs []int, categoryID int) ([]*entity.ThemeCategoryWithPostCount, error) {
	var rows []*entity.ThemeCategoryWithPostCount

	/*
		theme_category AS theme_category, tc
		post_area_category AS pac
		post_theme_category AS ptc

		SELECT theme_category.*, post_count FROM `theme_category` JOIN (
			SELECT tc.theme_id as id, count(tc.id) AS post_count FROM post_theme_category ptc
			JOIN post_area_category pac ON ptc.post_id = pac.post_id
			JOIN theme_category tc ON ptc.theme_category_id = tc.id
			WHERE pac.area_category_id = '3'
			GROUP BY tc.theme_id
		) t ON theme_category.id = t.id
		WHERE (theme_category.type = 'Theme')
		ORDER BY post_count DESC
		LIMIT 1000
	*/
	if err := r.buildQueryByExcludeIDs(excludeIDs).Limit(defaultAcquisitionNumber).
		Table("theme_category").
		Select("theme_category.*, post_count").
		Joins(`JOIN (
			SELECT tc.theme_id as id, count(tc.id) AS post_count FROM post_theme_category ptc 
			JOIN post_area_category pac ON ptc.post_id = pac.post_id 
			JOIN theme_category tc ON ptc.theme_category_id = tc.id 
			WHERE pac.area_category_id = ? 
			GROUP BY tc.theme_id
			) t ON theme_category.id = t.id`, categoryID).
		Where("theme_category.type = ?", model.ThemeCategoryTypeTheme).
		Order("post_count DESC").
		Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find theme_category")
	}

	return rows, nil
}

func (r *ThemeCategoryQueryRepositoryImpl) buildQueryByExcludeIDs(excludeIDs []int) *gorm.DB {
	if len(excludeIDs) > 0 {
		return r.DB.Not("id", excludeIDs)
	}
	return r.DB
}
