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

// Review参照系レポジトリ実装
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
