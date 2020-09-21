package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/logger"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"go.uber.org/zap"
)

const (
	bulkSize = 100
	debug    = false
)

type Script struct {
	service.MediaCommandService
	AWSSession *session.Session
	DB         *gorm.DB
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	script, err := InitializeScript("./config.yaml")
	if err != nil {
		return errors.Wrap(err, "failed to initialize script")
	}

	return script.Run()
}

func (s Script) Run() error {
	if err := s.resizeReviewMedia(); err != nil {
		return errors.WithStack(err)
	}

	if err := s.resizeUserMedia(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s Script) resizeReviewMedia() error {
	offset := 0
	for {
		var rows []*entity.ReviewMedia
		q := s.DB.Limit(bulkSize).Offset(offset)
		if err := q.Find(&rows).Error; err != nil {
			return errors.WithStack(err)
		}

		if len(rows) == 0 {
			break
		}

		for _, row := range rows {
			if row.ID == "73c91fd4-29cb-4996-942c-75f1ba0bc630" {
				// 異常データなので弾く
				continue
			}

			req := &model.PersistMediaRequest{
				UUID:        row.ID,
				Destination: row.S3Path(),
				MediaType:   s.mediaTypeForReview(row),
			}
			if err := s.MediaCommandService.Persist(req); err != nil {
				// return errors.Wrapf(err, "failed to resize review media: %s", row.ID)
				logger.Error("failed to resize review media", zap.String("uuid", row.ID), zap.Error(err))
			}

			if debug {
				return nil
			}
		}
		offset += len(rows)
	}

	return nil
}

func (s Script) resizeUserMedia() error {
	offset := 0
	for {
		var rows []*entity.UserTiny
		q := s.DB.Limit(bulkSize).Offset(offset)
		if err := q.Find(&rows).Error; err != nil {
			return errors.WithStack(err)
		}

		if len(rows) == 0 {
			break
		}

		for _, row := range rows {
			if row.AvatarUUID != "" {
				req := &model.PersistMediaRequest{
					UUID:        row.AvatarUUID,
					Destination: model.UserS3Path(row.AvatarUUID),
					MediaType:   model.MediaTypeUserIcon,
				}
				if err := s.MediaCommandService.Persist(req); err != nil {
					if !serror.IsErrorCode(err, serror.CodeNotFound) {
						// return errors.Wrapf(err, "failed to resize avatar: %s", row.AvatarUUID)
						logger.Error("failed to resize avatar", zap.String("uuid", row.AvatarUUID), zap.Error(err))
					}
				}
			}
			if row.HeaderUUID != "" {
				req := &model.PersistMediaRequest{
					UUID:        row.HeaderUUID,
					Destination: model.UserS3Path(row.HeaderUUID),
					MediaType:   model.MediaTypeUserHeader,
				}
				if err := s.MediaCommandService.Persist(req); err != nil {
					if !serror.IsErrorCode(err, serror.CodeNotFound) {
						// return errors.Wrapf(err, "failed to resize header: %s", row.HeaderUUID)
						logger.Error("failed to resize header", zap.String("uuid", row.HeaderUUID), zap.Error(err))
					}
				}
			}

			if debug {
				return nil
			}
		}
		offset += len(rows)
	}

	return nil
}

func (s Script) mediaTypeForReview(media *entity.ReviewMedia) model.MediaType {
	switch {
	case strings.HasPrefix(media.MimeType, "video/"):
		return model.MediaTypeReviewVideo
	case strings.HasPrefix(media.MimeType, "image/"):
		return model.MediaTypeReviewImage
	default:
		logger.Warn(fmt.Sprintf("unkown media type: %s", media.MimeType))
		return model.MediaType(0)
	}
}
