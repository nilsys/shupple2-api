package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

type (
	FeatureCommandRepository interface {
		Lock(c context.Context, id int) (*entity.Feature, error)
		Store(c context.Context, feature *entity.Feature) error
		UndeleteByID(c context.Context, id int) error
		DeleteByID(id int) error
		UpdateViewsByID(id, views int) error
		UpdateMonthlyViewsByID(id, views int) error
		UpdateWeeklyViewsByID(id, views int) error
	}

	FeatureQueryRepository interface {
		FindAll() ([]*entity.Feature, error)
		FindByID(id int) (*entity.Feature, error)
		FindQueryFeatureByID(id int) (*entity.FeatureDetailWithPosts, error)
		FindList(query *query.FindListPaginationQuery) (*entity.FeatureList, error)
	}
)
