package repository

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type UserCommandRepositoryImpl struct {
	DB            *gorm.DB
	MediaUploader *s3manager.Uploader
	AWSConfig     config.AWS
}

const (
	avatarKeyFormat = "avatars/%d"
)

var UserCommandRepositorySet = wire.NewSet(
	wire.Struct(new(UserCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.UserCommandRepository), new(*UserCommandRepositoryImpl)),
)

func (r *UserCommandRepositoryImpl) Store(user *entity.User) error {
	return errors.Wrap(r.DB.Save(user).Error, "failed to save user")
}

func (r *UserCommandRepositoryImpl) StoreWithAvatar(user *entity.User, avatar []byte) error {
	return Transaction(r.DB, func(tx *gorm.DB) error {
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
		r.DB.Exec("UPDATE user SET wordpress_id = ? WHERE wordpress_id = 0 AND id = ?", wordpressUserID, userID).Error,
		"failed to update user wordpress id",
	)
}

func (r *UserCommandRepositoryImpl) StoreFollow(following *entity.UserFollowing, followed *entity.UserFollowed) error {
	return Transaction(r.DB, func(tx *gorm.DB) error {
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
	return Transaction(r.DB, func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ? AND target_id = ?", userID, targetID).Delete(entity.UserFollowing{}).Error; err != nil {
			return errors.Wrap(err, "failed to delete user_following")
		}

		if err := tx.Where("user_id = ? AND target_id = ?", targetID, userID).Delete(entity.UserFollowed{}).Error; err != nil {
			return errors.Wrap(err, "failed to delete user_followed")
		}

		return nil
	})
}
