package output

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

type (
	Category struct {
		ID        int             `json:"id"`
		Name      string          `json:"name"`
		Slug      string          `json:"slug"`
		Type      string          `json:"type"`
		AreaGroup model.AreaGroup `json:"areaGroup"`
	}

	AreaCategory struct {
		ID   int                    `json:"id"`
		Name string                 `json:"name"`
		Slug string                 `json:"slug"`
		Type model.AreaCategoryType `json:"type"`
	}

	ThemeCategory struct {
		ID   int                     `json:"id"`
		Name string                  `json:"name"`
		Slug string                  `json:"slug"`
		Type model.ThemeCategoryType `json:"type"`
	}

	ThemeCategoryWithPostCount struct {
		ThemeCategory
		PostCount int `json:"postCount"`
	}

	ThemeCategoryDetail struct {
		ID       int                     `json:"id"`
		Name     string                  `json:"name"`
		Slug     string                  `json:"slug"`
		Type     model.ThemeCategoryType `json:"type"`
		Theme    *ThemeCategory          `json:"theme"`
		SubTheme *ThemeCategory          `json:"subTheme,omitempty"`
	}

	SpotCategory struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	}

	AreaCategoryDetail struct {
		ID         int                    `json:"id"`
		Name       string                 `json:"name"`
		Slug       string                 `json:"slug"`
		Type       model.AreaCategoryType `json:"type"`
		AreaGroup  model.AreaGroup        `json:"areaGroup"`
		Area       *AreaCategory          `json:"area"`
		SubArea    *AreaCategory          `json:"subArea,omitempty"`
		SubSubArea *AreaCategory          `json:"subSubArea,omitempty"`
	}

	AreaCategoryWithPostCount struct {
		AreaCategory
		PostCount int `json:"postCount"`
	}

	AreaCategoryWithThemeCategory struct {
		AreaCategory    *AreaCategory    `json:"area"`
		ThemeCategories []*ThemeCategory `json:"theme"`
	}
)

func NewCategory(id int, name, categoryType, slug string, group model.AreaGroup) *Category {
	return &Category{
		ID:        id,
		Name:      name,
		Slug:      slug,
		Type:      categoryType,
		AreaGroup: group,
	}
}

func NewSpotCategory(id int, name, slug string) *SpotCategory {
	return &SpotCategory{
		ID:   id,
		Name: name,
		Slug: slug,
	}
}
