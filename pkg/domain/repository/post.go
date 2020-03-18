package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

type (
	PostCommandRepository interface {
		Store(c context.Context, post *entity.Post) error
		DeleteByID(c context.Context, id int) error
		IncrementFavoriteCount(c context.Context, postID int) error
		DecrementFavoriteCount(c context.Context, postID int) error
	}

	PostQueryRepository interface {
		FindByID(id int) (*entity.Post, error)
		FindQueryShowByID(id int) (*entity.PostDetailWithHashtag, error)
		FindListByParams(query *query.FindPostListQuery) (*entity.PostDetailList, error)
		FindFeedListByUserID(userID int, query *query.FindListPaginationQuery) ([]*entity.PostDetail, error)
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
