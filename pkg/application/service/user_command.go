package service

import (
	"context"
	"strconv"
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository/payjp"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/logger"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/command"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/service"
	"go.uber.org/zap"
	"gopkg.in/guregu/null.v3"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity/wordpress"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	UserCommandService interface {
		SignUp(user *entity.User, cognitoToken string, migrationCode *string) error
		Update(user *entity.User, cmd *command.UpdateUser) error
		ImportFromWordpressByID(wordpressUserID int) error
		Follow(user *entity.User, targetID int) error
		Unfollow(user *entity.User, targetID int) error
	}

	UserCommandServiceImpl struct {
		repository.UserCommandRepository
		repository.UserQueryRepository
		repository.WordpressQueryRepository
		service.UserValidatorDomainService
		payjp.CustomerCommandRepository
		payjp.CustomerQueryRepository
		AuthService
		service.NoticeDomainService
		TransactionService
	}
)

var UserCommandServiceSet = wire.NewSet(
	wire.Struct(new(UserCommandServiceImpl), "*"),
	wire.Bind(new(UserCommandService), new(*UserCommandServiceImpl)),
)

func (s *UserCommandServiceImpl) SignUp(user *entity.User, cognitoToken string, migrationCode *string) error {

	if err := s.UserValidatorDomainService.Do(user); err != nil {
		return errors.Wrap(err, "invalid")
	}

	// token検証
	cognitoID, err := s.AuthService.Authorize(cognitoToken)
	if err != nil {
		return serror.New(err, serror.CodeUnauthorized, "unauthorized")
	}
	user.CognitoID = null.StringFrom(cognitoID)
	user.AddAttribute(model.UserAttributeCommon)

	if migrationCode != nil && *migrationCode != "" {
		existingUser, err := s.UserQueryRepository.FindByMigrationCode(*migrationCode)
		if err != nil {
			return errors.Wrap(err, "failed to get user by migration code")
		}
		user.ID = existingUser.ID
		if !existingUser.IsNonLogin {
			// 属性付与
			user.AddAttribute(model.UserAttributeWP)
		}
		user.UID = existingUser.UID
		user.WordpressID = existingUser.WordpressID
		user.AvatarUUID = existingUser.AvatarUUID
		user.HeaderUUID = existingUser.HeaderUUID
	}

	return s.TransactionService.Do(func(ctx context.Context) error {
		if err := s.UserCommandRepository.Store(ctx, user); err != nil {
			return errors.Wrap(err, "failed to store user")
		}

		_, err = s.CustomerQueryRepository.FindCustomer(user.PayjpCustomerID())
		if err != nil {
			if !serror.IsErrorCode(err, serror.CodeNotFound) {
				return errors.Wrap(err, "failed register to payjp")
			}
			if err := s.CustomerCommandRepository.StoreCustomer(user.PayjpCustomerID(), user.Email); err != nil {
				return errors.Wrap(err, "failed store customer to pay.jp")
			}
		}

		return nil
	})
}

// TODO: エラー時はslackに通知飛ばしたほうが良さそう
func (s *UserCommandServiceImpl) ImportFromWordpressByID(wordpressUserID int) error {
	wpUser, err := s.WordpressQueryRepository.FindUserByID(wordpressUserID)
	if err != nil {
		return errors.Wrapf(err, "failed to get wordpress user(id=%d)", wordpressUserID)
	}

	user, err := s.UserQueryRepository.FindByWordpressID(wordpressUserID)
	if err != nil {
		if !serror.IsErrorCode(err, serror.CodeNotFound) {
			return errors.Wrapf(err, "failed to import wordpress user(id=%d)", wordpressUserID)
		}

		// 新規登録かつメディア側で登録済みの場合
		if wpUser.Attributes.MediaUserID != "" {
			return s.updateMapping(wpUser)
		}

		// 新規登録でメディア側で登録がない場合
		return s.storeWithAvatar(entity.NewUserByWordpressUser(wpUser), wpUser)
	}

	// 更新の場合
	// すでにログイン済みの場合は無視
	if user.CognitoID.Valid {
		const msg = "tried to import user already logged in"
		logger.Info(msg, zap.Int("wordpress_user_id", wordpressUserID))
		return serror.New(nil, serror.CodeInvalidParam, msg)
	}

	// TODO: lock取ったほうがいいかも？
	user.PatchByWordpressUser(wpUser)
	return s.storeWithAvatar(user, wpUser)
}

func (s *UserCommandServiceImpl) updateMapping(wpUser *wordpress.User) error {
	mediaUserID, err := strconv.Atoi(wpUser.Attributes.MediaUserID)
	if err != nil {
		return errors.Wrap(err, "invalid media_user_id")
	}

	targetUser, err := s.UserQueryRepository.FindByID(mediaUserID)
	if err != nil {
		return errors.Wrap(err, "failed to find target user")
	}
	if targetUser.WordpressID.Valid {
		return serror.New(nil, serror.CodeInvalidParam, "already mapped user; wordpress_user_id=%d, target_user_id=%d", wpUser.ID, mediaUserID)
	}

	return s.UserCommandRepository.UpdateWordpressID(targetUser.ID, wpUser.ID)
}

func (s *UserCommandServiceImpl) storeWithAvatar(user *entity.User, wpUser *wordpress.User) error {
	var (
		avatar *wordpress.MediaBody
		err    error
	)

	if wpUser.Meta.WPUserAvatar != 0 {
		avatar, err = s.WordpressQueryRepository.FetchMediaBodyByID(wpUser.Meta.WPUserAvatar)
	} else {
		avatar, err = s.WordpressQueryRepository.FetchResource(wpUser.AvatarURLs.Num96)
	}
	if err != nil {
		return errors.Wrap(err, "failed to download avatar")
	}
	defer avatar.Body.Close()

	if err := s.UserCommandRepository.StoreWithAvatar(user, avatar.Body, avatar.ContentType); err != nil {
		return errors.Wrap(err, "faield to store user avatar")
	}

	return nil
}

func (s *UserCommandServiceImpl) Follow(user *entity.User, targetID int) error {
	if user.IsSelfID(targetID) {
		return serror.New(nil, serror.CodeInvalidParam, "can not follow self")
	}
	following := entity.NewUserFollowing(user.ID, targetID)
	followed := entity.NewUserFollowed(targetID, user.ID)

	return s.TransactionService.Do(func(c context.Context) error {
		if err := s.UserCommandRepository.StoreFollow(c, following, followed); err != nil {
			return errors.Wrap(err, "failed to store follow")
		}

		return s.NoticeDomainService.FollowUser(c, following)
	})
}

func (s *UserCommandServiceImpl) Unfollow(user *entity.User, targetID int) error {
	if user.IsSelfID(targetID) {
		return serror.New(nil, serror.CodeInvalidParam, "can not un follow self")
	}
	return s.UserCommandRepository.DeleteFollow(user.ID, targetID)
}

func (s *UserCommandServiceImpl) Update(user *entity.User, cmd *command.UpdateUser) error {
	s.updateUserCmd(user, cmd)

	if err := s.UserCommandRepository.Update(user); err != nil {
		return errors.Wrap(err, "failed to update user")
	}

	if err := s.persistUserImage(cmd); err != nil {
		return errors.Wrap(err, "failed to persist user image")
	}

	return nil
}

func (s *UserCommandServiceImpl) persistUserImage(cmd *command.UpdateUser) error {
	if err := s.UserCommandRepository.PersistUserImage(cmd); err != nil {
		return errors.Wrapf(err, "failed to persist user image")
	}
	return nil
}

func (s *UserCommandServiceImpl) updateUserCmd(user *entity.User, cmd *command.UpdateUser) {
	user.Name = cmd.Name
	user.Email = cmd.Email
	user.Birthdate = time.Time(cmd.BirthDate)
	user.Gender = cmd.Gender
	user.Profile = cmd.Profile
	// 画像だけは更新が無い場合はputしない
	if cmd.IconUUID != "" {
		user.AvatarUUID = cmd.IconUUID
	}
	if cmd.HeaderUUID != "" {
		user.HeaderUUID = cmd.HeaderUUID
	}
	user.URL = cmd.URL
	user.FacebookURL = cmd.FacebookURL
	user.InstagramURL = cmd.InstagramURL
	user.TwitterURL = cmd.TwitterURL
	user.YoutubeURL = cmd.YoutubeURL
	user.LivingArea = cmd.LivingArea
	user.UserInterests = cmd.Interests
}
