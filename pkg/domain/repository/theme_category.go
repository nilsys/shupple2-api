package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type (
	ThemeCategoryCommandRepository interface {
		Lock(c context.Context, id int) (*entity.ThemeCategory, error)
		Store(c context.Context, category *entity.ThemeCategory) error
		DeleteByID(id int) error
	}

	ThemeCategoryQueryRepository interface {
		FindByID(id int) (*entity.ThemeCategory, error)
		FindBySlug(slug string) (*entity.ThemeCategory, error)
		FindByIDs(ids []int) ([]*entity.ThemeCategory, error)

		// name部分一致検索
		SearchByName(name string) ([]*entity.ThemeCategory, error)

		FindThemesByAreaCategoryID(areaID, subAreaID, subSubAreaID int, excludeIDs []int) ([]*entity.ThemeCategoryWithPostCount, error)
		FindSubThemesByAreaCategoryIDAndParentThemeID(parentThemeID, areaID, subAreaID, subSubAreaID int, excludeIDs []int) ([]*entity.ThemeCategoryWithPostCount, error)
	}
)
