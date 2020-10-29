package aws

import (
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type (
	S3CommandRepository interface {
		Upload(input *s3manager.UploadInput) error
		Delete(key string) error
	}
)
