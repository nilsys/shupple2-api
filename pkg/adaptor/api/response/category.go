package response

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	Category struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
		Type string `json:"type"`
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

	Lcategory struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	}
)

func NewCategory(id int, name, categoryType, slug string) *Category {
	return &Category{
		ID:   id,
		Name: name,
		Slug: slug,
		Type: categoryType,
	}
}

func NewLcategory(id int, name, slug string) *Lcategory {
	return &Lcategory{
		ID:   id,
		Name: name,
		Slug: slug,
	}
}
