package repository

import (
	"context"
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

func (r *UserCommandRepositoryImpl) Store(ctx context.Context, user *entity.User) error {
	return errors.Wrap(r.DB(ctx).Save(user).Error, "failed to save user")
}

// StoreはUpdateにも対応している、その為論理削除されたものにStore()を使うと、新たに作成(INSERT)しようとするので、PKで不整合が起きる
// 論理削除された物を更新等する場合はStore()では無く、この関数を使う事
func (r *UserCommandRepositoryImpl) UnscopedStore(ctx context.Context, user *entity.User) error {
	return errors.Wrap(r.DB(ctx).Unscoped().Save(user).Error, "failed to save user")
}

func (r *UserCommandRepositoryImpl) Update(user *entity.User) error {
	if err := r.DB(context.TODO()).Save(user).Error; err != nil {
		return errors.Wrapf(err, "failed to update user id=%d", user.ID)
	}

	return nil
}

func (r *UserCommandRepositoryImpl) UpdateDeviceTokenByID(id int, deviceToken string) error {
	if err := r.DB(context.TODO()).Exec("UPDATE user SET device_token = ? WHERE id = ?", deviceToken, id).Error; err != nil {
		return errors.Wrap(err, "failed update user.device_token")
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
			Key:         aws.String(model.UserS3Path(user.AvatarUUID)),
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
		r.DB(context.TODO()).Exec("UPDATE user SET wordpress_id = ? WHERE wordpress_id IS NULL AND id = ?", wordpressUserID, userID).Error,
		"failed to update user wordpress id",
	)
}

func (r *UserCommandRepositoryImpl) StoreFollow(c context.Context, following *entity.UserFollowing, followed *entity.UserFollowed) error {
	if err := r.DB(c).Save(followed).Error; err != nil {
		return errors.Wrap(err, "failed to save user_followed")
	}

	if err := r.DB(c).Save(following).Error; err != nil {
		return errors.Wrap(err, "failed to save user_following")
	}
	return nil
}

func (r *UserCommandRepositoryImpl) DeleteFollow(userID, targetID int) error {
	if err := r.DB(context.Background()).Where("user_id = ? AND target_id = ?", userID, targetID).Delete(entity.UserFollowing{}).Error; err != nil {
		return errors.Wrap(err, "failed to delete user_following")
	}

	if err := r.DB(context.Background()).Where("user_id = ? AND target_id = ?", targetID, userID).Delete(entity.UserFollowed{}).Error; err != nil {
		return errors.Wrap(err, "failed to delete user_followed")
	}

	return nil
}

func (r *UserCommandRepositoryImpl) StoreUserBlock(userBlock *entity.UserBlockUser) error {
	if err := r.DB(context.Background()).Save(userBlock).Error; err != nil {
		return errors.Wrap(err, "failed store user_block_user")
	}
	return nil
}

func (r *UserCommandRepositoryImpl) DeleteUserBlock(userBlock *entity.UserBlockUser) error {
	if err := r.DB(context.Background()).Delete(userBlock).Error; err != nil {
		return errors.Wrap(err, "failed delete user_block_user")
	}
	return nil
}

func (r *UserCommandRepositoryImpl) UpdateHeaderUUIDToBlank(ctx context.Context, userID int) error {
	if err := r.DB(ctx).Exec("UPDATE user SET header_uuid = NULL WHERE id = ?", userID).Error; err != nil {
		return errors.Wrap(err, "failed update avatar_uuid")
	}
	return nil
}

func (r *UserCommandRepositoryImpl) UpdateIconUUIDToNull(ctx context.Context, userID int) error {
	if err := r.DB(ctx).Exec("UPDATE user SET avatar_uuid = '' WHERE id = ?", userID).Error; err != nil {
		return errors.Wrap(err, "failed update header_uuid")
	}
	return nil
}
