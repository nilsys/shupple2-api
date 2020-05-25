package repository

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type ReviewFavoriteQueryRepositoryImpl struct {
	DB *gorm.DB
}

var ReviewFavoriteQueryRepositorySet = wire.NewSet(
	wire.Struct(new(ReviewFavoriteQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.ReviewFavoriteQueryRepository), new(*ReviewFavoriteQueryRepositoryImpl)),
)

func (r *ReviewFavoriteQueryRepositoryImpl) IsExist(userID, reviewID int) (bool, error) {
	var row entity.UserFavoriteReview

	err := r.DB.Where("user_id = ? AND review_id = ?", userID, reviewID).First(&row).Error

	return ErrorToIsExist(err, "user_favorite_post(user_id=%d,review_id=%d)", userID, reviewID)
}

func (r *ReviewFavoriteQueryRepositoryImpl) IsExistReviewCommentReply(userID, reviewCommentReplyID int) (bool, error) {
	var row entity.UserFavoriteReviewCommentReply

	err := r.DB.Where("user_id = ? AND review_comment_reply_id = ?", userID, reviewCommentReplyID).First(&row).Error

	return ErrorToIsExist(err, "user_favorite_review_comment_reply(user_id=%d,review_comment_reply_id=%d)", userID, reviewCommentReplyID)
}
