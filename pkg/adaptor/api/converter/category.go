package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

func (c Converters) ConvertCategoryToOutput(category entity.Category, group model.AreaGroup) *output.Category {
	return output.NewCategory(category.CategoryID(), category.CategoryName(), category.CategoryType(), category.CategorySlug(), group)
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

func (c Converters) ConvertThemeCategoryWithPostCountToOutput(themeCategory *entity.ThemeCategoryWithPostCount) *output.ThemeCategoryWithPostCount {
	return &output.ThemeCategoryWithPostCount{
		ThemeCategory: output.ThemeCategory{
			ID:   themeCategory.ID,
			Name: themeCategory.Name,
			Slug: themeCategory.Slug,
			Type: themeCategory.Type,
		},
		PostCount: themeCategory.PostCount,
	}
}

func (c Converters) ConvertThemeCategoriesWithPostCountToOutput(themeCategories []*entity.ThemeCategoryWithPostCount) []*output.ThemeCategoryWithPostCount {
	var resp = make([]*output.ThemeCategoryWithPostCount, len(themeCategories))
	for i, themeCategory := range themeCategories {
		resp[i] = c.ConvertThemeCategoryWithPostCountToOutput(themeCategory)
	}

	return resp
}

func (c Converters) ConvertAreaCategoryWithPostCountToOutput(areaCategory *entity.AreaCategoryWithPostCount) *output.AreaCategoryWithPostCount {
	return &output.AreaCategoryWithPostCount{
		AreaCategory: *c.ConvertAreaCategoryToOutput(&areaCategory.AreaCategory),
		PostCount:    areaCategory.PostCount,
	}
}

func (c Converters) ConvertAreaCategoriesWithPostCountToOutput(areaCategories []*entity.AreaCategoryWithPostCount) []*output.AreaCategoryWithPostCount {
	var resp = make([]*output.AreaCategoryWithPostCount, len(areaCategories))
	for i, areaCategory := range areaCategories {
		resp[i] = c.ConvertAreaCategoryWithPostCountToOutput(areaCategory)
	}

	return resp
}
