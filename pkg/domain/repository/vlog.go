package repository

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

type (
	VlogCommandRepository interface {
		Store(vlog *entity.Vlog) error
		DeleteByID(id int) error
	}

	VlogQueryRepository interface {
		FindByID(id int) (*entity.Vlog, error)
		FindListByParams(query *query.FindVlogListQuery) ([]*entity.QueryVlog, error)
	}
)
