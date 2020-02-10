package repository

import "github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

type (
	VlogCommandRepository interface {
		Store(vlog *entity.Vlog) error
	}

	VlogQueryRepository interface {
		FindByID(id int) (*entity.Vlog, error)
	}
)
