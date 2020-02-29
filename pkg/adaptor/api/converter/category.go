package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/response"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

func ConvertCategoryToOutput(category *entity.Category) *response.Category {
	return &response.Category{
		ID:   category.ID,
		Name: category.Name,
	}
}

func ConvertCategoriesToOutput(categories []*entity.Category) []*response.Category {
	var resp = make([]*response.Category, len(categories))
	for i, category := range categories {
		resp[i] = ConvertCategoryToOutput(category)
	}

	return resp
}
