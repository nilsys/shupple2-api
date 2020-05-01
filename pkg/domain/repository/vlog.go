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
		UpdateViewsByID(id, views int) error
	}

	VlogQueryRepository interface {
		FindAll() ([]*entity.Vlog, error)
		FindByID(id int) (*entity.Vlog, error)
		FindListByParams(query *query.FindVlogListQuery) (*entity.VlogList, error)
		FindDetailByID(id int) (*entity.VlogDetail, error)
	}
)
