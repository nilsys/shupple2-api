package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/response"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

func ConvertCategoryToOutput(category entity.Category) *response.Category {
	return response.NewCategory(category.CategoryID(), category.CategoryName(), category.CategoryType())
}

func ConvertAreaCategoryToOutput(areaCategory *entity.AreaCategory) *response.AreaCategory {
	return &response.AreaCategory{
		ID:   areaCategory.ID,
		Name: areaCategory.Name,
		Type: areaCategory.Type,
	}
}

func ConvertAreaCategoriesToOutput(areaCategories []*entity.AreaCategory) []*response.AreaCategory {
	var resp = make([]*response.AreaCategory, len(areaCategories))
	for i, areaCategory := range areaCategories {
		resp[i] = ConvertAreaCategoryToOutput(areaCategory)
	}

	return resp
}

func ConvertThemeCategoryToOutput(themeCategory *entity.ThemeCategory) *response.ThemeCategory {
	return &response.ThemeCategory{
		ID:   themeCategory.ID,
		Name: themeCategory.Name,
		Type: themeCategory.Type,
	}
}

func ConvertThemeCategoriesToOutput(themeCategories []*entity.ThemeCategory) []*response.ThemeCategory {
	var resp = make([]*response.ThemeCategory, len(themeCategories))
	for i, themeCategory := range themeCategories {
		resp[i] = ConvertThemeCategoryToOutput(themeCategory)
	}

	return resp
}
