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
		UndeleteByID(c context.Context, id int) error
		DeleteByID(c context.Context, id int) error
		IncrementFavoriteCount(c context.Context, postID int) error
		DecrementFavoriteCount(c context.Context, postID int) error
		UpdateViewsByID(id, views int) error
		UpdateMonthlyViewsByID(id, views int) error
		UpdateWeeklyViewsByID(id, views int) error
	}

	PostQueryRepository interface {
		FindByLastID(lastID, limit int) ([]*entity.Post, error)
		FindByID(id int) (*entity.Post, error)
		FindByIDs(ids []int) (*entity.PostDetailList, error)
		FindPostDetailWithHashtagByID(id int) (*entity.PostDetailWithHashtagAndIsFavorite, error)
		FindPostDetailWithHashtagAndIsFavoriteByID(id, userID int) (*entity.PostDetailWithHashtagAndIsFavorite, error)
		FindPostDetailWithHashtagBySlug(slug string) (*entity.PostDetailWithHashtagAndIsFavorite, error)
		FindPostDetailWithHashtagAndIsFavoriteBySlug(slug string, userID int) (*entity.PostDetailWithHashtagAndIsFavorite, error)
		FindListByParams(query *query.FindPostListQuery) (*entity.PostList, error)
		FindListWithIsFavoriteByParams(query *query.FindPostListQuery, userID int) (*entity.PostList, error)
		FindFeedListByUserID(targetUserID int, query *query.FindListPaginationQuery) (*entity.PostList, error)
		FindFeedListWithIsFavoriteByUserID(userID, targetUserID int, query *query.FindListPaginationQuery) (*entity.PostList, error)
		FindFavoriteListByUserID(targetUseID int, query *query.FindListPaginationQuery) (*entity.PostList, error)
		FindFavoriteListWithIsFavoriteByUserID(userID, targetUserID int, query *query.FindListPaginationQuery) (*entity.PostList, error)
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
