package repository

import "github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

type (
	FeatureCommandRepository interface {
		Store(feature *entity.Feature) error
	}

	FeatureQueryRepository interface {
		FindByID(id int) (*entity.Feature, error)
	}
)
