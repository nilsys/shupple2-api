package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

func ConvertCategoryToOutput(category entity.Category) *output.Category {
	return output.NewCategory(category.CategoryID(), category.CategoryName(), category.CategoryType(), category.CategorySlug())
}

func ConvertAreaCategoryToOutput(areaCategory *entity.AreaCategory) *output.AreaCategory {
	return &output.AreaCategory{
		ID:   areaCategory.ID,
		Name: areaCategory.Name,
		Slug: areaCategory.Slug,
		Type: areaCategory.Type,
	}
}

func ConvertAreaCategoriesToOutput(areaCategories []*entity.AreaCategory) []*output.AreaCategory {
	var resp = make([]*output.AreaCategory, len(areaCategories))
	for i, areaCategory := range areaCategories {
		resp[i] = ConvertAreaCategoryToOutput(areaCategory)
	}

	return resp
}

func ConvertThemeCategoryToOutput(themeCategory *entity.ThemeCategory) *output.ThemeCategory {
	return &output.ThemeCategory{
		ID:   themeCategory.ID,
		Name: themeCategory.Name,
		Slug: themeCategory.Slug,
		Type: themeCategory.Type,
	}
}

func ConvertThemeCategoriesToOutput(themeCategories []*entity.ThemeCategory) []*output.ThemeCategory {
	var resp = make([]*output.ThemeCategory, len(themeCategories))
	for i, themeCategory := range themeCategories {
		resp[i] = ConvertThemeCategoryToOutput(themeCategory)
	}

	return resp
}

func ConvertAreaCategoryDetailToOutput(areaCategoryDetail *entity.AreaCategoryDetail) *output.AreaCategoryDetail {
	var subArea *output.AreaCategory
	var subSubArea *output.AreaCategory
	if areaCategoryDetail.SubAreaID.Valid {
		subArea = ConvertAreaCategoryToOutput(areaCategoryDetail.SubArea)
	}
	if areaCategoryDetail.SubSubAreaID.Valid {
		subSubArea = ConvertAreaCategoryToOutput(areaCategoryDetail.SubSubArea)
	}
	return &output.AreaCategoryDetail{
		ID:         areaCategoryDetail.ID,
		Name:       areaCategoryDetail.Name,
		Slug:       areaCategoryDetail.Slug,
		Type:       areaCategoryDetail.Type,
		Area:       ConvertAreaCategoryToOutput(areaCategoryDetail.Area),
		SubArea:    subArea,
		SubSubArea: subSubArea,
	}
}

func ConvertThemeCategoryWithPostCountToOutput(themeCategory *entity.ThemeCategoryWithPostCount) *output.ThemeCategory {
	return &output.ThemeCategory{
		ID:        themeCategory.ID,
		Name:      themeCategory.Name,
		Slug:      themeCategory.Slug,
		Type:      themeCategory.Type,
		PostCount: themeCategory.PostCount,
	}
}

func ConvertThemeCategoriesWithPostCountToOutput(themeCategories []*entity.ThemeCategoryWithPostCount) []*output.ThemeCategory {
	var resp = make([]*output.ThemeCategory, len(themeCategories))
	for i, themeCategory := range themeCategories {
		resp[i] = ConvertThemeCategoryWithPostCountToOutput(themeCategory)
	}

	return resp
}
