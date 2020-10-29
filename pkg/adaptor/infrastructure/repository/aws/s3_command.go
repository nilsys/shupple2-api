package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/uma-co82/shupple2-api/pkg/config"
	shuppleAWS "github.com/uma-co82/shupple2-api/pkg/domain/repository/aws"
)

type (
	S3CommandRepositoryImpl struct {
		AWSSession *session.Session
		AWSConfig  config.AWS
	}
)

var S3CommandRepositorySet = wire.NewSet(
	wire.Struct(new(S3CommandRepositoryImpl), "*"),
	wire.Bind(new(shuppleAWS.S3CommandRepository), new(*S3CommandRepositoryImpl)),
)

// TODO: 抽象化
func (r *S3CommandRepositoryImpl) Upload(input *s3manager.UploadInput) error {
	uploader := s3manager.NewUploader(r.AWSSession)
	_, err := uploader.Upload(input)
	return errors.Wrap(err, "failed upload to s3")
}

func (r *S3CommandRepositoryImpl) Delete(key string) error {
	deleteInut := &s3.DeleteObjectInput{
		Bucket: aws.String(r.AWSConfig.FilesBucket),
		Key:    aws.String(key),
	}

	svc := s3.New(r.AWSSession)
	_, err := svc.DeleteObject(deleteInut)

	return errors.Wrap(err, "failed delete from s3")
}
