package repository

import "github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

type (
	CategoryCommandRepository interface {
		Store(category *entity.Category) error
	}

	CategoryQueryRepository interface {
		FindByID(id int) (*entity.Category, error)
		FindByIDs(ids []int) ([]*entity.Category, error)
	}
)
