package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

type (
	// Review参照系レポジトリ
	ReviewQueryRepository interface {
		FindByID(id int) (*entity.Review, error)
		ShowReviewListByParams(query *query.ShowReviewListQuery) ([]*entity.QueryReview, error)
		FindFeedReviewListByUserID(userID int, query *query.FindListPaginationQuery) ([]*entity.QueryReview, error)
		FindReviewCommentListByReviewID(reviewID int, limit int) ([]*entity.ReviewComment, error)
		FindQueryReviewByID(id int) (*entity.QueryReview, error)
		FindFavoriteListByUserID(userID int, query *query.FindListPaginationQuery) ([]*entity.QueryReview, error)
		IsExist(id int) (bool, error)
		IsExistReviewComment(id int) (bool, error)
		IsExistReviewCommentFavorite(userID, reviewCommentID int) (bool, error)
		FindReviewCommentByID(id int) (*entity.ReviewComment, error)
		FindReviewCommentReplyListByReviewCommentID(reviewCommentID int) ([]*entity.ReviewCommentReply, error)
	}

	// Review更新系レポジトリ
	ReviewCommandRepository interface {
		StoreReview(c context.Context, review *entity.Review) error
		DeleteReview(c context.Context, review *entity.Review) error
		IncrementReviewCommentCount(c context.Context, reviewID int) error
		DecrementReviewCommentCount(c context.Context, reviewID int) error
		IncrementFavoriteCount(c context.Context, reviewID int) error
		DecrementFavoriteCount(c context.Context, reviewID int) error
		PersistReviewMedia(reviewMedia *entity.ReviewMedia) error
		ShowReviewComment(c context.Context, commentID int) (*entity.ReviewComment, error)
		DeleteReviewCommentByID(c context.Context, reviewComment *entity.ReviewComment) error
		CreateReviewComment(c context.Context, comment *entity.ReviewComment) error
		StoreReviewCommentReply(c context.Context, reply *entity.ReviewCommentReply) error
		IncrementReviewCommentReplyCount(c context.Context, reviewCommentID int) error
		IncrementReviewCommentFavoriteCount(c context.Context, reviewCommentID int) error
		DecrementReviewCommentFavoriteCount(c context.Context, reviewCommentID int) error
		StoreReviewCommentFavorite(c context.Context, favorite *entity.UserFavoriteReviewComment) error
		DeleteReviewCommentFavoriteByID(c context.Context, userID, reviewCommentID int) error
	}

	ReviewFavoriteCommandRepository interface {
		Store(c context.Context, favorite *entity.UserFavoriteReview) error
		Delete(c context.Context, unfavorite *entity.UserFavoriteReview) error
	}

	ReviewFavoriteQueryRepository interface {
		IsExist(userID, reviewID int) (bool, error)
	}
)
