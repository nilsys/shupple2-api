package response

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	Category struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
	}

	AreaCategory struct {
		ID   int                    `json:"id"`
		Name string                 `json:"name"`
		Type model.AreaCategoryType `json:"type"`
	}

	ThemeCategory struct {
		ID   int                     `json:"id"`
		Name string                  `json:"name"`
		Type model.ThemeCategoryType `json:"type"`
	}

	Lcategory struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

func NewCategory(id int, name string, categoryType string) *Category {
	return &Category{
		ID:   id,
		Name: name,
		Type: categoryType,
	}
}

func NewLcategory(id int, name string) *Lcategory {
	return &Lcategory{
		ID:   id,
		Name: name,
	}
}
