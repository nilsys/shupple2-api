package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
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
