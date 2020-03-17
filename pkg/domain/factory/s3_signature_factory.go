package factory

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"

	"github.com/google/wire"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type (
	S3SignatureFactory struct {
		Session   *session.Session
		AWSConfig config.AWS
	}
)

var S3SignatureFactorySet = wire.NewSet(
	wire.Struct(new(S3SignatureFactory), "*"),
)

func (f *S3SignatureFactory) GenerateS3SignatureEntity(contentType string) (*entity.S3Signature, error) {
	uuidStr, err := model.NewRandUUID()
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate uuid")
	}

	svc := s3.New(f.Session)
	resp, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket:      aws.String(f.AWSConfig.FilesBucket),
		Key:         aws.String(f.generateUploadURL(uuidStr)),
		ContentType: aws.String(contentType),
	})

	url, err := resp.Presign(f.AWSConfig.UploadExpire)
	if err != nil {
		return nil, errors.Wrap(err, "failed to presign url")
	}

	return &entity.S3Signature{
		UUID: uuidStr,
		URL:  url,
	}, nil
}

// TODO: 仮置き
func (f *S3SignatureFactory) generateUploadURL(uuidStr string) string {
	return fmt.Sprintf("tmp/%s", uuidStr)
}
