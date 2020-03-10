package repository

import "github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

type (
	InterestQueryRepository interface {
		FindAll() ([]*entity.Interest, error)
	}
)
