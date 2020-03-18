package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

type (
	// Review参照系レポジトリ
	ReviewQueryRepository interface {
		ShowReviewListByParams(query *query.ShowReviewListQuery) ([]*entity.QueryReview, error)
		FindFeedReviewListByUserID(userID int, query *query.FindListPaginationQuery) ([]*entity.QueryReview, error)
		FindReviewCommentListByReviewID(reviewID int, limit int) ([]*entity.ReviewComment, error)
		FindQueryReviewByID(id int) (*entity.QueryReview, error)
		FindByID(reviewID int) (*entity.Review, error)
		IsExist(id int) (bool, error)
	}

	ReviewFavoriteCommandRepository interface {
		Store(c context.Context, favorite *entity.UserFavoriteReview) error
		Delete(c context.Context, unfavorite *entity.UserFavoriteReview) error
	}

	ReviewFavoriteQueryRepository interface {
		IsExist(userID, reviewID int) (bool, error)
	}

	// Reviewコマンド系レポジトリ
	ReviewCommandRepository interface {
		StoreReview(c context.Context, review *entity.Review) error
		IncrementFavoriteCount(c context.Context, reviewID int) error
		DecrementFavoriteCount(c context.Context, reviewID int) error
		PersistReviewMedia(reviewMedia *entity.ReviewMedia) error
	}
)
