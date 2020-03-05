package response

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	Category struct {
		ID   int                `json:"id"`
		Name string             `json:"name"`
		Type model.CategoryType `json:"type"`
	}
)

func NewCategory(id int, name string, categoryType model.CategoryType) Category {
	return Category{
		ID:   id,
		Name: name,
		Type: categoryType,
	}
}
