package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

type (
	AreaCategoryCommandRepository interface {
		Lock(c context.Context, id int) (*entity.AreaCategory, error)
		Store(c context.Context, category *entity.AreaCategory) error
		DeleteByID(id int) error
	}

	AreaCategoryQueryRepository interface {
		FindByID(id int) (*entity.AreaCategory, error)
		FindByIDAndType(id int, areaCategoryType model.AreaCategoryType) (*entity.AreaCategory, error)
		FindDetailByIDAndType(id int, areaCategoryType model.AreaCategoryType) (*entity.AreaCategoryDetail, error)
		FindBySlug(slug string) (*entity.AreaCategory, error)
		FindByIDs(ids []int) ([]*entity.AreaCategory, error)
		FindAreaListHavingPostByAreaGroup(areaGroup model.AreaGroup, limit int, excludeID []int) ([]*entity.AreaCategoryWithPostCount, error)
		FindSubAreaListHavingPostByAreaIDAndThemeID(areaID, themeID, limit int, excludeID []int) ([]*entity.AreaCategoryWithPostCount, error)
		FindSubSubAreaListHavingPostBySubAreaIDAndThemeID(subAreaID, themeID int, limit int, excludeID []int) ([]*entity.AreaCategoryWithPostCount, error)
		// name部分一致検索
		SearchByName(name string) (*entity.AreaCategories, error)
		SearchAreaByName(name string) (*entity.AreaCategories, error)
		SearchSubAreaByName(name string) (*entity.AreaCategories, error)
		SearchSubSubAreaByName(name string) (*entity.AreaCategories, error)
	}
)
