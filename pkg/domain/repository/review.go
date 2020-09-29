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
		FindFeedReviewWithIsFavoriteListByUserID(userID int, query *query.FindListPaginationQuery) (*entity.ReviewDetailWithIsFavoriteList, error)
		FindFavoriteReviewListByUserID(targetUserID int, query *query.FindListPaginationQuery) (*entity.ReviewDetailWithIsFavoriteList, error)
		FindFavoriteReviewWithIsFavoriteListByUserID(userID, targetUserID int, query *query.FindListPaginationQuery) (*entity.ReviewDetailWithIsFavoriteList, error)
		FindReviewCommentListByReviewID(reviewID int, limit int) ([]*entity.ReviewCommentWithIsFavorite, error)
		FindReviewCommentWithIsFavoriteListByReviewID(reviewID int, limit int, userID int) ([]*entity.ReviewCommentWithIsFavorite, error)
		FindReviewCommentReplyListByReviewCommentID(reviewCommentID int) ([]*entity.ReviewCommentReplyWithIsFavorite, error)
		FindReviewCommentReplyWithIsFavoriteListByReviewCommentID(reviewCommentID int, userID int) ([]*entity.ReviewCommentReplyWithIsFavorite, error)
		FindQueryReviewByID(id int) (*entity.ReviewDetailWithIsFavorite, error)
		FindQueryReviewWithIsFavoriteByID(id, userID int) (*entity.ReviewDetailWithIsFavorite, error)
		IsExist(id int) (bool, error)
		IsExistReviewComment(id int) (bool, error)
		IsExistReviewCommentReply(id int) (bool, error)
		IsExistReviewCommentFavorite(userID, reviewCommentID int) (bool, error)
		FindReviewCommentByID(id int) (*entity.ReviewCommentTiny, error)
		FindReviewCommentDetailByID(id int) (*entity.ReviewCommentDetail, error)
		FindReviewCommentReplyByID(id int) (*entity.ReviewCommentReplyTiny, error)
		FindReviewCommentReplyDetailByID(id int) (*entity.ReviewCommentReplyDetail, error)
		FindLatestHasMediaReviewByTouristSpotIDs(touristSpotID []int) (*entity.ReviewList, error)
		FindRelationLocationReview(query *query.FindListPaginationQuery, userID int) (*entity.ReviewDetailWithIsFavoriteList, error)
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
		ShowReviewComment(c context.Context, commentID int) (*entity.ReviewCommentTiny, error)
		StoreReviewComment(c context.Context, comment *entity.ReviewCommentTiny) error
		StoreReviewCommentReply(c context.Context, reply *entity.ReviewCommentReplyTiny) error
		IncrementReviewCommentReplyCount(c context.Context, reviewCommentID int) error
		DecrementReviewCommentReplyCount(c context.Context, reviewCommentID int) error
		IncrementReviewCommentFavoriteCount(c context.Context, reviewCommentID int) error
		DecrementReviewCommentFavoriteCount(c context.Context, reviewCommentID int) error
		StoreReviewCommentFavorite(c context.Context, favorite *entity.UserFavoriteReviewComment) error
		DeleteReviewCommentFavoriteByID(c context.Context, userID, reviewCommentID int) error
		IncrementReviewCommentReplyFavoriteCount(c context.Context, reviewCommentReplyID int) error
		DecrementReviewCommentReplyFavoriteCount(c context.Context, reviewCommentReplyID int) error
		UpdateViewsByID(id, views int) error
		UpdateMonthlyViewsByID(id, views int) error
		UpdateWeeklyViewsByID(id, views int) error
		DeleteReviewByID(c context.Context, id int) error
		DeleteReviewCommentByID(c context.Context, id int) error
		DeleteReviewCommentReplyByID(c context.Context, id int) error
	}

	ReviewFavoriteCommandRepository interface {
		Store(c context.Context, favorite *entity.UserFavoriteReview) error
		Delete(c context.Context, unfavorite *entity.UserFavoriteReview) error
		StoreReviewCommentReply(c context.Context, favorite *entity.UserFavoriteReviewCommentReply) error
		DeleteReviewCommentReply(c context.Context, favorite *entity.UserFavoriteReviewCommentReply) error
	}

	ReviewFavoriteQueryRepository interface {
		IsExist(userID, reviewID int) (bool, error)
		IsExistReviewCommentReply(userID, reviewCommentReplyID int) (bool, error)
	}
)
