package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/wire"
	"github.com/pkg/errors"
	shuppleAWS "github.com/uma-co82/shupple2-api/pkg/domain/repository/aws"
)

type (
	S3CommandRepositoryImpl struct {
		AWSSession *session.Session
	}
)

var S3CommandRepositorySet = wire.NewSet(
	wire.Struct(new(S3CommandRepositoryImpl), "*"),
	wire.Bind(new(shuppleAWS.S3CommandRepository), new(*S3CommandRepositoryImpl)),
)

func (r *S3CommandRepositoryImpl) Upload(input *s3manager.UploadInput) error {
	uploader := s3manager.NewUploader(r.AWSSession)
	_, err := uploader.Upload(input)
	return errors.Wrap(err, "failed upload to s3")
}
