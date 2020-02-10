package repository

import "github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

type (
	LcategoryCommandRepository interface {
		Store(lcategory *entity.Lcategory) error
	}

	LcategoryQueryRepository interface {
		FindByID(id int) (*entity.Lcategory, error)
	}
)
