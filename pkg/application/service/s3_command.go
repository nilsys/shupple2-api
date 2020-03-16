package service

import (
	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/factory"
)

type (
	S3CommandService interface {
		GenerateS3Signature(contentType string) (*entity.S3Signature, error)
	}

	S3CommandServiceImpl struct {
		factory.S3SignatureFactory
	}
)

var S3CommandServiceSet = wire.NewSet(
	wire.Struct(new(S3CommandServiceImpl), "*"),
	wire.Bind(new(S3CommandService), new(*S3CommandServiceImpl)),
)

func (s *S3CommandServiceImpl) GenerateS3Signature(contentType string) (*entity.S3Signature, error) {
	return s.GenerateS3SignatureEntity(contentType)
}
