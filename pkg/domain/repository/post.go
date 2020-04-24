package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

type (
	PostCommandRepository interface {
		Lock(c context.Context, id int) (*entity.Post, error)
		Store(c context.Context, post *entity.Post) error
		DeleteByID(c context.Context, id int) error
		IncrementFavoriteCount(c context.Context, postID int) error
		DecrementFavoriteCount(c context.Context, postID int) error
	}

	PostQueryRepository interface {
		FindByID(id int) (*entity.Post, error)
		FindPostDetailWithHashtagByID(id int) (*entity.PostDetailWithHashtagAndIsFavorite, error)
		FindPostDetailWithHashtagAndIsFavoriteByID(id, userID int) (*entity.PostDetailWithHashtagAndIsFavorite, error)
		FindPostDetailWithHashtagBySlug(slug string) (*entity.PostDetailWithHashtag, error)
		FindListByParams(query *query.FindPostListQuery) (*entity.PostList, error)
		FindListWithIsFavoriteByParams(query *query.FindPostListQuery, userID int) (*entity.PostList, error)
		FindFeedListByUserID(userID int, query *query.FindListPaginationQuery) ([]*entity.PostDetail, error)
		FindFavoriteListByUserID(userID int, query *query.FindListPaginationQuery) ([]*entity.PostDetail, error)
		IsExist(id int) (bool, error)
	}

	PostFavoriteCommandRepository interface {
		Store(c context.Context, favorite *entity.UserFavoritePost) error
		Delete(c context.Context, favorite *entity.UserFavoritePost) error
	}

	PostFavoriteQueryRepository interface {
		IsExist(userID, postID int) (bool, error)
	}
)
