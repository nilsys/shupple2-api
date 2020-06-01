package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type ReviewFavoriteCommandRepositoryImpl struct {
	DAO
}

var ReviewFavoriteCommandRepositorySet = wire.NewSet(
	wire.Struct(new(ReviewFavoriteCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.ReviewFavoriteCommandRepository), new(*ReviewFavoriteCommandRepositoryImpl)),
)

func (r *ReviewFavoriteCommandRepositoryImpl) Store(c context.Context, favorite *entity.UserFavoriteReview) error {
	return errors.Wrap(r.DB(c).Save(favorite).Error, "failed to save favorite")
}

func (r *ReviewFavoriteCommandRepositoryImpl) Delete(c context.Context, unfavorite *entity.UserFavoriteReview) error {
	return errors.Wrap(r.DB(c).Where("user_id = ? AND review_id = ?", unfavorite.UserID, unfavorite.ReviewID).Delete(&unfavorite).Error, "failed to delete")
}

func (r *ReviewFavoriteCommandRepositoryImpl) StoreReviewCommentReply(c context.Context, favorite *entity.UserFavoriteReviewCommentReply) error {
	return errors.Wrap(r.DB(c).Save(favorite).Error, "failed to save favorite")
}

func (r *ReviewFavoriteCommandRepositoryImpl) DeleteReviewCommentReply(c context.Context, unfavorite *entity.UserFavoriteReviewCommentReply) error {
	return errors.Wrap(r.DB(c).Where("user_id = ? AND review_comment_reply_id = ?", unfavorite.UserID, unfavorite.ReviewCommentReplyID).Delete(unfavorite).Error, "failed to delete")
}
