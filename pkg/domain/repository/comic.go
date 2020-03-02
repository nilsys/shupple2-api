package repository

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

type (
	ComicCommandRepository interface {
		Store(comic *entity.Comic) error
		DeleteByID(id int) error
	}

	ComicQueryRepository interface {
		FindByID(id int) (*entity.QueryComic, error)
		FindListOrderByCreatedAt(query *query.FindListPaginationQuery) ([]*entity.Comic, error)
	}
)
