package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/response"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

func ConvertCategoryToOutput(category entity.Category) *response.Category {
	return response.NewCategory(category.CategoryID(), category.CategoryName(), category.CategoryType(), category.CategorySlug())
}

func ConvertAreaCategoryToOutput(areaCategory *entity.AreaCategory) *response.AreaCategory {
	return &response.AreaCategory{
		ID:   areaCategory.ID,
		Name: areaCategory.Name,
		Slug: areaCategory.Slug,
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
		Slug: themeCategory.Slug,
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

func ConvertAreaCategoryDetailToOutput(areaCategoryDetail *entity.AreaCategoryDetail) *response.AreaCategoryDetail {
	var subArea *response.AreaCategory
	var subSubArea *response.AreaCategory
	if areaCategoryDetail.SubAreaID.Valid {
		subArea = ConvertAreaCategoryToOutput(areaCategoryDetail.SubArea)
	}
	if areaCategoryDetail.SubSubAreaID.Valid {
		subSubArea = ConvertAreaCategoryToOutput(areaCategoryDetail.SubSubArea)
	}
	return &response.AreaCategoryDetail{
		ID:         areaCategoryDetail.ID,
		Name:       areaCategoryDetail.Name,
		Slug:       areaCategoryDetail.Slug,
		Type:       areaCategoryDetail.Type,
		Area:       ConvertAreaCategoryToOutput(areaCategoryDetail.Area),
		SubArea:    subArea,
		SubSubArea: subSubArea,
	}
}
