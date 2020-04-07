package repository

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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

const (
	avatarKeyFormat = "avatars/%d"
)

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

func (r *UserCommandRepositoryImpl) StoreWithAvatar(user *entity.User, avatar []byte) error {
	return Transaction(r.DB(context.TODO()), func(tx *gorm.DB) error {
		if err := tx.Save(user).Error; err != nil {
			return errors.Wrap(err, "failed to save user")
		}

		key := fmt.Sprintf(avatarKeyFormat, user.ID)
		_, err := r.MediaUploader.Upload(&s3manager.UploadInput{
			Bucket: &r.AWSConfig.FilesBucket,
			Key:    &key,
			Body:   bytes.NewReader(avatar),
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

// TODO: 消す事考える
func (r *UserCommandRepositoryImpl) PersistUserImage(user *entity.User) error {
	if user.AvatarUUID != "" {
		avatarFrom := fmt.Sprint(r.AWSConfig.FilesBucket, "/", model.UploadedS3Path(user.AvatarUUID))
		svc := s3.New(r.AWSSession)

		avatarGetReq := &s3.GetObjectInput{
			Bucket: aws.String(r.AWSConfig.FilesBucket),
			Key:    aws.String(avatarFrom),
		}

		o, err := svc.GetObject(avatarGetReq)
		if err != nil {
			return errors.Wrap(err, "failed to get s3 tmp object")
		}

		avatarCopyReq := &s3.CopyObjectInput{
			CopySource:  aws.String(avatarFrom),
			Bucket:      aws.String(r.AWSConfig.FilesBucket),
			Key:         aws.String(user.S3AvatarPath()),
			ContentType: o.ContentType,
		}

		_, err = svc.CopyObject(avatarCopyReq)
		if err != nil {
			return errors.Wrap(err, "failed to copy s3 object")
		}
	}

	if user.HeaderUUID != "" {
		headerFrom := fmt.Sprint(r.AWSConfig.FilesBucket, "/", model.UploadedS3Path(user.HeaderUUID))
		svc := s3.New(r.AWSSession)

		headerGetReq := &s3.GetObjectInput{
			Bucket: aws.String(r.AWSConfig.FilesBucket),
			Key:    aws.String(headerFrom),
		}

		o, err := svc.GetObject(headerGetReq)
		if err != nil {
			return errors.Wrap(err, "failed to get s3 tmp object")
		}

		headerCopyReq := &s3.CopyObjectInput{
			CopySource:  aws.String(headerFrom),
			Bucket:      aws.String(r.AWSConfig.FilesBucket),
			Key:         aws.String(user.S3HeaderPath()),
			ContentType: o.ContentType,
		}

		_, err = svc.CopyObject(headerCopyReq)
		if err != nil {
			return errors.Wrap(err, "failed to copy s3 object")
		}
	}

	return nil
}
