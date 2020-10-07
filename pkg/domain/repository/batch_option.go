package repository

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	BatchOptionCommandRepository interface {
		UpdateByName(name model.BatchOptionName, val string) error
	}

	BatchOptionQueryRepository interface {
		FindByName(name model.BatchOptionName) (string, error)
	}
)
