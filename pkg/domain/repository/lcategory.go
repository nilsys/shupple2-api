package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type (
	LcategoryCommandRepository interface {
		Lock(c context.Context, id int) (*entity.Lcategory, error)
		Store(c context.Context, lcategory *entity.Lcategory) error
	}

	LcategoryQueryRepository interface {
		FindByID(id int) (*entity.Lcategory, error)
	}
)
