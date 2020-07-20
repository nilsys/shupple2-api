package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type (
	SpotCategoryCommandRepository interface {
		Lock(c context.Context, id int) (*entity.SpotCategory, error)
		Store(c context.Context, spotCategory *entity.SpotCategory) error
		DeleteByID(id int) error
	}

	SpotCategoryQueryRepository interface {
		FindByID(id int) (*entity.SpotCategory, error)
	}
)
