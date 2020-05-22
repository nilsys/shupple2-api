package repository

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	uuid "github.com/satori/go.uuid"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type UserCommandRepositoryImpl struct {
	DAO
	MediaUploader *s3manager.Uploader
	AWSConfig     config.AWS
	AWSSession    *session.Session
}

var UserCommandRepositorySet = wire.NewSet(
	wire.Struct(new(UserCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.UserCommandRepository), new(*UserCommandRepositoryImpl)),
)

func (r *UserCommandRepositoryImpl) Store(user *entity.User) error {
	return errors.Wrap(r.DB(context.TODO()).Save(user).Error, "failed to save user")
}

func (r *UserCommandRepositoryImpl) Update(user *entity.User) error {
	if err := r.DB(context.TODO()).Save(user).Error; err != nil {
		return errors.Wrapf(err, "failed to update user id=%d", user.ID)
	}

	return nil
}

func (r *UserCommandRepositoryImpl) StoreWithAvatar(user *entity.User, avatar io.Reader, contentType string) error {
	user.AvatarUUID = uuid.NewV4().String()
	return Transaction(r.DB(context.TODO()), func(tx *gorm.DB) error {
		if err := tx.Save(user).Error; err != nil {
			return errors.Wrap(err, "failed to save user")
		}

		_, err := r.MediaUploader.Upload(&s3manager.UploadInput{
			Bucket:      aws.String(r.AWSConfig.FilesBucket),
			Key:         aws.String(user.S3AvatarPath()),
			Body:        avatar,
			ACL:         aws.String(s3.ObjectCannedACLPublicRead),
			ContentType: aws.String(contentType),
		})

		if err != nil {
			return errors.Wrap(err, "failed to save upload avatar")
		}

		return nil

	})
}

func (r *UserCommandRepositoryImpl) UpdateWordpressID(userID, wordpressUserID int) error {
	return errors.Wrap(
		r.DB(context.TODO()).Exec("UPDATE user SET wordpress_id = ? WHERE wordpress_id = 0 AND id = ?", wordpressUserID, userID).Error,
		"failed to update user wordpress id",
	)
}

func (r *UserCommandRepositoryImpl) StoreFollow(c context.Context, following *entity.UserFollowing, followed *entity.UserFollowed) error {
	return Transaction(r.DB(c), func(tx *gorm.DB) error {
		if err := tx.Save(followed).Error; err != nil {
			return errors.Wrap(err, "failed to save user_followed")
		}

		if err := tx.Save(following).Error; err != nil {
			return errors.Wrap(err, "failed to save user_following")
		}
		return nil
	})
}

func (r *UserCommandRepositoryImpl) DeleteFollow(userID, targetID int) error {
	return Transaction(r.DB(context.TODO()), func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ? AND target_id = ?", userID, targetID).Delete(entity.UserFollowing{}).Error; err != nil {
			return errors.Wrap(err, "failed to delete user_following")
		}

		if err := tx.Where("user_id = ? AND target_id = ?", targetID, userID).Delete(entity.UserFollowed{}).Error; err != nil {
			return errors.Wrap(err, "failed to delete user_followed")
		}

		return nil
	})
}

func (r *UserCommandRepositoryImpl) PersistUserImage(user *entity.User) error {
	svc := s3.New(r.AWSSession)

	if user.AvatarUUID != "" {
		if err := r.persistImage(svc, user.AvatarUUID, user.S3AvatarPath()); err != nil {
			return errors.Wrap(err, "failed to persist avatar")
		}
	}

	if user.HeaderUUID != "" {
		if err := r.persistImage(svc, user.HeaderUUID, user.S3HeaderPath()); err != nil {
			return errors.Wrap(err, "failed to persist header")
		}
	}

	return nil
}

func (r *UserCommandRepositoryImpl) persistImage(svc *s3.S3, uuid, dest string) error {
	from := fmt.Sprint(r.AWSConfig.FilesBucket, "/", model.UploadedS3Path(uuid))

	headReq := &s3.HeadObjectInput{
		Bucket: aws.String(r.AWSConfig.FilesBucket),
		Key:    aws.String(model.UploadedS3Path(uuid)),
	}

	o, err := svc.HeadObject(headReq)
	if err != nil {
		return errors.Wrap(err, "failed to get s3 tmp object")
	}

	copyReq := &s3.CopyObjectInput{
		CopySource:  aws.String(from),
		Bucket:      aws.String(r.AWSConfig.FilesBucket),
		Key:         aws.String(dest),
		ACL:         aws.String(s3.ObjectCannedACLPublicRead),
		ContentType: o.ContentType,
	}

	_, err = svc.CopyObject(copyReq)
	if err != nil {
		return errors.Wrap(err, "failed to copy s3 object")
	}

	return nil
}
