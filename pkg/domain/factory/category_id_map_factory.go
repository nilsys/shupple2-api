package factory

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	CategoryIDMapFactory struct {
		repository.AreaCategoryQueryRepository
		repository.ThemeCategoryQueryRepository
	}
)

var CategoryIDMapFactorySet = wire.NewSet(
	wire.Struct(new(CategoryIDMapFactory), "*"),
)

func (f *CategoryIDMapFactory) GenerateCategoryIDMap(areaIDs, themeIDs []int) (map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, error) {
	areaCategories, err := f.AreaCategoryQueryRepository.FindByIDs(areaIDs)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed find area_category by ids")
	}
	themeCategories, err := f.ThemeCategoryQueryRepository.FindByIDs(themeIDs)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed find theme_category by ids")
	}

	return f.newAreaCategoryIDMap(areaCategories), f.newThemeCategoryIDMap(themeCategories), nil
}

// id: AreaCategoryのマップを返す
func (f *CategoryIDMapFactory) newAreaCategoryIDMap(areaCategories []*entity.AreaCategory) map[int]*entity.AreaCategory {
	areaCategoriesMap := make(map[int]*entity.AreaCategory, len(areaCategories))

	for _, areaCategory := range areaCategories {
		areaCategoriesMap[areaCategory.ID] = areaCategory
	}

	return areaCategoriesMap
}

// id: ThemeCategoryのマップを返す
func (f *CategoryIDMapFactory) newThemeCategoryIDMap(themeCategories []*entity.ThemeCategory) map[int]*entity.ThemeCategory {
	themeCategoriesMap := make(map[int]*entity.ThemeCategory, len(themeCategories))

	for _, themeCategory := range themeCategories {
		themeCategoriesMap[themeCategory.ID] = themeCategory
	}

	return themeCategoriesMap
}
