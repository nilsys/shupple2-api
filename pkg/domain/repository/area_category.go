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
	}

	AreaCategoryQueryRepository interface {
		FindByID(id int) (*entity.AreaCategory, error)
		FindByIDAndType(id int, areaCategoryType model.AreaCategoryType) (*entity.AreaCategory, error)
		FindBySlug(slug string) (*entity.AreaCategory, error)
		FindByIDs(ids []int) ([]*entity.AreaCategory, error)

		FindAreaListByAreaGroup(areaGroup model.AreaGroup, limit int, excludeID []int) ([]*entity.AreaCategory, error)
		FindSubAreaListByAreaID(areaID, limit int, excludeID []int) ([]*entity.AreaCategory, error)
		FindSubSubAreaListBySubAreaID(subAreaID, limit int, excludeID []int) ([]*entity.AreaCategory, error)

		// FindByTouristSpotID(touristSpotID int) ([]*entity.AreaCategory, error)
		// FindByMetaSearchID(innAreaTypeIDs *entity.InnAreaTypeIDs) ([]*entity.AreaCategory, error)

		// name部分一致検索
		SearchByName(name string) ([]*entity.AreaCategory, error)
		SearchAreaByName(name string) ([]*entity.AreaCategory, error)
		SearchSubAreaByName(name string) ([]*entity.AreaCategory, error)
		SearchSubSubAreaByName(name string) ([]*entity.AreaCategory, error)
	}
)
