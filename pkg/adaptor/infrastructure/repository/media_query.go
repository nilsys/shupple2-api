package repository

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	MediaQueryRepositoryImpl struct {
		AWSConfig  config.AWS
		AWSSession *session.Session
	}
)

var MediaQueryRepositorySet = wire.NewSet(
	wire.Struct(new(MediaQueryRepositoryImpl), "*"),
	wire.Bind(new(repository.MediaQueryRepository), new(*MediaQueryRepositoryImpl)),
)

func (r *MediaQueryRepositoryImpl) GetUploadedMedia(uuid string) (*model.MediaBody, error) {
	key := model.UploadedS3Path(uuid)
	var getObjectInput s3.GetObjectInput
	getObjectInput.SetBucket(r.AWSConfig.FilesBucket).SetKey(key)

	svc := s3.New(r.AWSSession)
	getObjectOutput, err := svc.GetObject(&getObjectInput)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == s3.ErrCodeNoSuchKey {
			return nil, serror.New(awsErr, serror.CodeNotFound, "uploaded media (%s) not found", key)
		}
		return nil, errors.Wrap(err, "failed to get uploaded media")
	}

	var contentType string
	if getObjectOutput.ContentType != nil {
		contentType = *getObjectOutput.ContentType
	}

	return &model.MediaBody{
		ContentType: contentType,
		Body:        getObjectOutput.Body,
	}, nil
}
