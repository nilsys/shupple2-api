package repository

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	MediaCommandRepositoryImpl struct {
		AWSConfig  config.AWS
		AWSSession *session.Session
	}
)

var MediaCommandRepositorySet = wire.NewSet(
	wire.Struct(new(MediaCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.MediaCommandRepository), new(*MediaCommandRepositoryImpl)),
)

func (r *MediaCommandRepositoryImpl) SavePersistRequest(request *model.PersistMediaRequest) error {
	body, err := json.Marshal(request)
	if err != nil {
		return errors.Wrap(err, "failed to marshal persist media request")
	}

	var input sqs.SendMessageInput
	input.
		SetMessageBody(string(body)).
		SetQueueUrl(r.AWSConfig.PersistMediaQueue)

	svc := sqs.New(r.AWSSession)
	_, err = svc.SendMessage(&input)
	return errors.Wrap(err, "failed to send message to sqs")
}

func (r *MediaCommandRepositoryImpl) Save(mediaBody *model.MediaBody, destination string) error {
	uploadInput := &s3manager.UploadInput{
		ACL:         aws.String(s3.ObjectCannedACLPublicRead),
		Body:        mediaBody.Body,
		Bucket:      &r.AWSConfig.FilesBucket,
		Key:         &destination,
		ContentType: &mediaBody.ContentType,
	}

	uploader := s3manager.NewUploader(r.AWSSession)
	_, err := uploader.Upload(uploadInput)

	return errors.Wrap(err, "failed to upload media")
}
