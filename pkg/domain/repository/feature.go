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
		DeleteByID(id int) error
	}

	FeatureQueryRepository interface {
		FindByID(id int) (*entity.Feature, error)
		FindQueryFeatureByID(id int) (*entity.FeatureDetailWithPosts, error)
		FindList(query *query.FindListPaginationQuery) (*entity.FeatureList, error)
	}
)
