package repository

import "github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

type (
	ComicCommandRepository interface {
		Store(comic *entity.Comic) error
	}

	ComicQueryRepository interface {
		FindByID(id int) (*entity.Comic, error)
	}
)
