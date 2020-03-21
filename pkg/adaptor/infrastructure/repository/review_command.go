package repository

import (
	"context"

	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

// Review更新系レポジトリ実装
type ReviewCommandRepositoryImpl struct {
	DAO
	AWSSession *session.Session
	AWSConfig  config.AWS
}

var ReviewCommandRepositorySet = wire.NewSet(
	wire.Struct(new(ReviewCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.ReviewCommandRepository), new(*ReviewCommandRepositoryImpl)),
)

// TODO: updateの時にSaveの挙動確認
func (r *ReviewCommandRepositoryImpl) StoreReview(c context.Context, review *entity.Review) error {
	if err := r.DB(c).Save(review).Error; err != nil {
		return errors.Wrap(err, "failed store review")
	}
	return nil
}

func (r *ReviewCommandRepositoryImpl) CreateReviewComment(c context.Context, reviewComment *entity.ReviewComment) error {
	return errors.Wrap(r.DB(c).Save(reviewComment).Error, "failed to save review comment")
}

func (r *ReviewCommandRepositoryImpl) IncrementReviewCommentCount(c context.Context, reviewID int) error {
	if err := r.DB(c).
		Exec("UPDATE review SET comment_count=comment_count+1 WHERE id = ?", reviewID).
		Error; err != nil {
		return errors.Wrapf(err, "Failed insert reviews")
	}

	return nil
}

func (r *ReviewCommandRepositoryImpl) IncrementFavoriteCount(c context.Context, reviewID int) error {
	return errors.Wrapf(r.DB(c).Exec("UPDATE review SET favorite_count = favorite_count + 1 WHERE id = ?", reviewID).Error, "failed to update")
}

func (r *ReviewCommandRepositoryImpl) DecrementFavoriteCount(c context.Context, reviewID int) error {
	return errors.Wrapf(r.DB(c).Exec("UPDATE review SET favorite_count = favorite_count - 1 WHERE id = ?", reviewID).Error, "failed to update")
}

func (r *ReviewCommandRepositoryImpl) PersistReviewMedia(reviewMedia *entity.ReviewMedia) error {
	from := fmt.Sprint(r.AWSConfig.FilesBucket, "/", model.UploadedS3Path(reviewMedia.ID))
	req := &s3.CopyObjectInput{
		CopySource:  aws.String(from),
		Bucket:      aws.String(r.AWSConfig.FilesBucket),
		Key:         aws.String(reviewMedia.S3Path()),
		ContentType: aws.String(reviewMedia.MimeType),
	}
	_, err := s3.New(r.AWSSession).CopyObject(req)
	if err != nil {
		return errors.Wrap(err, "failed to copy s3 object")
	}

	return nil
}

func (r *ReviewCommandRepositoryImpl) StoreReviewCommentReply(c context.Context, reply *entity.ReviewCommentReply) error {
	if err := r.DB(c).Save(reply).Error; err != nil {
		return errors.Wrap(err, "failed to store review comment reply")
	}
	return nil
}

func (r *ReviewCommandRepositoryImpl) IncrementReviewCommentReplyCount(c context.Context, reviewCommentID int) error {
	if err := r.DB(c).Exec("UPDATE review_comment SET reply_count=reply_count+1 WHERE id = ?", reviewCommentID).Error; err != nil {
		return errors.Wrap(err, "failed to increment review_comment.reply_count")
	}
	return nil
}

func (r *ReviewCommandRepositoryImpl) IncrementReviewCommentFavoriteCount(c context.Context, reviewCommentID int) error {
	if err := r.DB(c).Exec("UPDATE review_comment SET favorite_count=favorite_count+1 WHERE id = ?", reviewCommentID).Error; err != nil {
		return errors.Wrap(err, "failed to increment review_comment.favorite_count")
	}
	return nil
}

func (r *ReviewCommandRepositoryImpl) DecrementReviewCommentFavoriteCount(c context.Context, reviewCommentID int) error {
	if err := r.DB(c).Exec("UPDATE review_comment SET favorite_count=favorite_count-1 WHERE id = ?", reviewCommentID).Error; err != nil {
		return errors.Wrap(err, "failed to decrement review_comment.favorite_count")
	}
	return nil
}

func (r *ReviewCommandRepositoryImpl) StoreReviewCommentFavorite(c context.Context, favorite *entity.UserFavoriteReviewComment) error {
	if err := r.DB(c).Save(favorite).Error; err != nil {
		return errors.Wrap(err, "failed to save review_comment_favorite")
	}
	return nil
}

func (r *ReviewCommandRepositoryImpl) DeleteReviewCommentFavoriteByID(c context.Context, userID, reviewCommentID int) error {
	if err := r.DB(c).Where("user_id = ? AND review_comment_id = ?", userID, reviewCommentID).Delete(entity.UserFavoriteReviewComment{}).Error; err != nil {
		return errors.Wrap(err, "failed to delete review_comment_favorite")
	}
	return nil
}
