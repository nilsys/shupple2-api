package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

type (
	VlogCommandRepository interface {
		Lock(c context.Context, id int) (*entity.Vlog, error)
		Store(c context.Context, vlog *entity.Vlog) error
		DeleteByID(id int) error
	}

	VlogQueryRepository interface {
		FindByID(id int) (*entity.Vlog, error)
		FindListByParams(query *query.FindVlogListQuery) (*entity.VlogDetailList, error)
		FindWithTouristSpotsByID(id int) (*entity.VlogDetailWithTouristSpots, error)
	}
)
