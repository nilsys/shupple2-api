package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

type (
	// Review参照系レポジトリ
	ReviewQueryRepository interface {
		FindAll() ([]*entity.Review, error)
		FindByID(id int) (*entity.Review, error)
		ShowReviewListByParams(query *query.ShowReviewListQuery) (*entity.ReviewDetailWithIsFavoriteList, error)
		ShowReviewWithIsFavoriteListByParams(query *query.ShowReviewListQuery, userID int) (*entity.ReviewDetailWithIsFavoriteList, error)
		FindFeedReviewListByUserID(userID int, query *query.FindListPaginationQuery) (*entity.ReviewDetailWithIsFavoriteList, error)
		FindFeedReviewWithIsFavoriteListByUserID(userID, targetUserID int, query *query.FindListPaginationQuery) (*entity.ReviewDetailWithIsFavoriteList, error)
		FindFavoriteReviewListByUserID(targetUserID int, query *query.FindListPaginationQuery) (*entity.ReviewDetailWithIsFavoriteList, error)
		FindFavoriteReviewWithIsFavoriteListByUserID(userID, targetUserID int, query *query.FindListPaginationQuery) (*entity.ReviewDetailWithIsFavoriteList, error)
		FindReviewCommentListByReviewID(reviewID int, limit int) ([]*entity.ReviewCommentWithIsFavorite, error)
		FindReviewCommentWithIsFavoriteListByReviewID(reviewID int, limit int, userID int) ([]*entity.ReviewCommentWithIsFavorite, error)
		FindQueryReviewByID(id int) (*entity.ReviewDetailWithIsFavorite, error)
		FindQueryReviewWithIsFavoriteByID(id, userID int) (*entity.ReviewDetailWithIsFavorite, error)
		IsExist(id int) (bool, error)
		IsExistReviewComment(id int) (bool, error)
		IsExistReviewCommentFavorite(userID, reviewCommentID int) (bool, error)
		FindReviewCommentByID(id int) (*entity.ReviewComment, error)
		FindReviewCommentReplyByID(id int) (*entity.ReviewCommentReply, error)
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
		StoreReviewComment(c context.Context, comment *entity.ReviewComment) error
		StoreReviewCommentReply(c context.Context, reply *entity.ReviewCommentReply) error
		IncrementReviewCommentReplyCount(c context.Context, reviewCommentID int) error
		IncrementReviewCommentFavoriteCount(c context.Context, reviewCommentID int) error
		DecrementReviewCommentFavoriteCount(c context.Context, reviewCommentID int) error
		StoreReviewCommentFavorite(c context.Context, favorite *entity.UserFavoriteReviewComment) error
		DeleteReviewCommentFavoriteByID(c context.Context, userID, reviewCommentID int) error
		UpdateViewsByID(id, views int) error
		DeleteReviewByID(c context.Context, id int) error
		DeleteReviewCommentByID(c context.Context, id int) error
		DeleteReviewCommentReplyByID(c context.Context, id int) error
	}

	ReviewFavoriteCommandRepository interface {
		Store(c context.Context, favorite *entity.UserFavoriteReview) error
		Delete(c context.Context, unfavorite *entity.UserFavoriteReview) error
	}

	ReviewFavoriteQueryRepository interface {
		IsExist(userID, reviewID int) (bool, error)
	}
)
