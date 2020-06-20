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
		UndeleteByID(c context.Context, id int) error
		DeleteByID(id int) error
		IncrementFavoriteCount(c context.Context, id int) error
		DecrementFavoriteCount(c context.Context, id int) error
	}

	ComicQueryRepository interface {
		FindByID(id int) (*entity.ComicDetail, error)
		FindWithIsFavoriteByID(id, userID int) (*entity.ComicDetail, error)
		FindListOrderByCreatedAt(query *query.FindListPaginationQuery) (*entity.ComicList, error)
		FindWithIsFavoriteListOrderByCreatedAt(query *query.FindListPaginationQuery, userID int) (*entity.ComicList, error)
		IsExist(id int) (bool, error)
	}

	ComicFavoriteCommandRepository interface {
		Store(c context.Context, favorite *entity.UserFavoriteComic) error
		Delete(c context.Context, unfavorite *entity.UserFavoriteComic) error
	}

	ComicFavoriteQueryRepository interface {
		IsExist(userID, comicID int) (bool, error)
	}
)
