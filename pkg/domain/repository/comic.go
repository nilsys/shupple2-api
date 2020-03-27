package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

type (
	ComicCommandRepository interface {
		Lock(c context.Context, id int) (*entity.Comic, error)
		Store(c context.Context, comic *entity.Comic) error
		DeleteByID(id int) error
	}

	ComicQueryRepository interface {
		FindByID(id int) (*entity.QueryComic, error)
		FindListOrderByCreatedAt(query *query.FindListPaginationQuery) (*entity.ComicList, error)
	}
)
