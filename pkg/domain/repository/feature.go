package repository

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

type (
	FeatureCommandRepository interface {
		Store(feature *entity.Feature) error
		DeleteByID(id int) error
	}

	FeatureQueryRepository interface {
		FindByID(id int) (*entity.Feature, error)
		FindQueryFeatureByID(id int) (*entity.QueryFeature, error)
		FindList(query *query.FindListPaginationQuery) ([]*entity.Feature, error)
	}
)
