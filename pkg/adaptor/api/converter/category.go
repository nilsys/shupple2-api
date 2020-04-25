package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

func (c Converters) ConvertCategoryToOutput(category entity.Category) *output.Category {
	return output.NewCategory(category.CategoryID(), category.CategoryName(), category.CategoryType(), category.CategorySlug())
}

func (c Converters) ConvertAreaCategoryToOutput(areaCategory *entity.AreaCategory) *output.AreaCategory {
	return &output.AreaCategory{
		ID:   areaCategory.ID,
		Name: areaCategory.Name,
		Slug: areaCategory.Slug,
		Type: areaCategory.Type,
	}
}

func (c Converters) ConvertAreaCategoriesToOutput(areaCategories []*entity.AreaCategory) []*output.AreaCategory {
	var resp = make([]*output.AreaCategory, len(areaCategories))
	for i, areaCategory := range areaCategories {
		resp[i] = c.ConvertAreaCategoryToOutput(areaCategory)
	}

	return resp
}

func (c Converters) ConvertThemeCategoryToOutput(themeCategory *entity.ThemeCategory) *output.ThemeCategory {
	return &output.ThemeCategory{
		ID:   themeCategory.ID,
		Name: themeCategory.Name,
		Slug: themeCategory.Slug,
		Type: themeCategory.Type,
	}
}

func (c Converters) ConvertThemeCategoriesToOutput(themeCategories []*entity.ThemeCategory) []*output.ThemeCategory {
	var resp = make([]*output.ThemeCategory, len(themeCategories))
	for i, themeCategory := range themeCategories {
		resp[i] = c.ConvertThemeCategoryToOutput(themeCategory)
	}

	return resp
}

func (c Converters) ConvertAreaCategoryDetailToOutput(areaCategoryDetail *entity.AreaCategoryDetail) *output.AreaCategoryDetail {
	var subArea *output.AreaCategory
	var subSubArea *output.AreaCategory
	if areaCategoryDetail.SubAreaID.Valid {
		subArea = c.ConvertAreaCategoryToOutput(areaCategoryDetail.SubArea)
	}
	if areaCategoryDetail.SubSubAreaID.Valid {
		subSubArea = c.ConvertAreaCategoryToOutput(areaCategoryDetail.SubSubArea)
	}
	return &output.AreaCategoryDetail{
		ID:         areaCategoryDetail.ID,
		Name:       areaCategoryDetail.Name,
		Slug:       areaCategoryDetail.Slug,
		Type:       areaCategoryDetail.Type,
		Area:       c.ConvertAreaCategoryToOutput(areaCategoryDetail.Area),
		SubArea:    subArea,
		SubSubArea: subSubArea,
	}
}

func (c Converters) ConvertThemeCategoryWithPostCountToOutput(themeCategory *entity.ThemeCategoryWithPostCount) *output.ThemeCategory {
	return &output.ThemeCategory{
		ID:        themeCategory.ID,
		Name:      themeCategory.Name,
		Slug:      themeCategory.Slug,
		Type:      themeCategory.Type,
		PostCount: themeCategory.PostCount,
	}
}

func (c Converters) ConvertThemeCategoriesWithPostCountToOutput(themeCategories []*entity.ThemeCategoryWithPostCount) []*output.ThemeCategory {
	var resp = make([]*output.ThemeCategory, len(themeCategories))
	for i, themeCategory := range themeCategories {
		resp[i] = c.ConvertThemeCategoryWithPostCountToOutput(themeCategory)
	}

	return resp
}
