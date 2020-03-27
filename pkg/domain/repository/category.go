package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

type (
	CategoryCommandRepository interface {
		Lock(c context.Context, id int) (*entity.Category, error)
		Store(c context.Context, category *entity.Category) error
	}

	CategoryQueryRepository interface {
		FindTypeByID(id int) (*model.CategoryType, error)
		FindByID(id int) (*entity.Category, error)
		FindBySlug(slug string) (*entity.Category, error)
		FindListByParentID(parentID int, limit int, excludeID []int) ([]*entity.Category, error)
		FindByIDs(ids []int) ([]*entity.Category, error)
		// name部分一致検索
		SearchByName(name string) ([]*entity.Category, error)
		FindByTouristSpotID(touristSpotID int) ([]*entity.Category, error)
		FindByMetaSearchID(innAreaTypeIDs *entity.InnAreaTypeIDs) ([]*entity.Category, error)
		SearchAreaByName(name string) ([]*entity.Category, error)
		SearchSubAreaByName(name string) ([]*entity.Category, error)
		SearchSubSubAreaByName(name string) ([]*entity.Category, error)
	}
)
