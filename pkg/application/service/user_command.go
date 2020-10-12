package service

import (
	"context"
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
		UpdateDeviceToken(user *entity.User, deviceToken string) error
		ImportFromWordpressByID(wordpressUserID int) error
		Follow(user *entity.User, targetID int) error
		Unfollow(user *entity.User, targetID int) error
		Block(user *entity.User, blockedUserID int) error
		Unblock(user *entity.User, blockedUserID int) error
		DeleteUserIcon(user *entity.User) error
		DeleteUserHeader(user *entity.User) error
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
		MediaCommandService
		repository.MediaCommandRepository
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

		// 非ログインユーザーは基本的なクエリに掛からない様、論理削除されている
		// その為、論理削除も更新対象であるUnscopedStore()を使う
		if err := s.UserCommandRepository.UnscopedStore(ctx, user); err != nil {
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

	existingUser, err := s.UserQueryRepository.FindByWordpressID(wordpressUserID)
	if err != nil {
		if !serror.IsErrorCode(err, serror.CodeNotFound) {
			return errors.Wrapf(err, "failed to import wordpress user(id=%d)", wordpressUserID)
		}

		mappingTargetUser, err := s.UserQueryRepository.FindByUID(string(wpUser.Slug))
		if err != nil {
			if !serror.IsErrorCode(err, serror.CodeNotFound) {
				return errors.Wrapf(err, "failed to find user by wordpress user slug(%s)", string(wpUser.Slug))
			}

			// 新規登録でメディア側で登録がない場合
			return s.storeWithAvatar(entity.NewUserByWordpressUser(wpUser), wpUser)
		}

		// 新規登録かつメディア側で登録済みの場合
		// ターゲットとして指定したユーザーが既にマッピングを持っていないかをチェックして大丈夫ならマッピングをセット
		if mappingTargetUser.WordpressID.Valid {
			return serror.New(nil, serror.CodeInvalidParam, "already mapped user; wordpress_user_id=%d, target_user_id=%d", wpUser.ID, mappingTargetUser.ID)
		}
		return s.UserCommandRepository.UpdateWordpressID(mappingTargetUser.ID, wpUser.ID)
	}

	// 更新の場合
	// すでにログイン済みの場合は無視
	if existingUser.CognitoID.Valid {
		const msg = "tried to import user already logged in"
		logger.Info(msg, zap.Int("wordpress_user_id", wordpressUserID))
		return serror.New(nil, serror.CodeInvalidParam, msg)
	}

	// TODO: lock取ったほうがいいかも？
	existingUser.PatchByWordpressUser(wpUser)
	return s.storeWithAvatar(existingUser, wpUser)
}

func (s *UserCommandServiceImpl) storeWithAvatar(user *entity.User, wpUser *wordpress.User) error {
	var (
		avatar *model.MediaBody
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

		return s.NoticeDomainService.FollowUser(c, following, user)
	})
}

func (s *UserCommandServiceImpl) Unfollow(user *entity.User, targetID int) error {
	if user.IsSelfID(targetID) {
		return serror.New(nil, serror.CodeInvalidParam, "can not un follow self")
	}
	return s.UserCommandRepository.DeleteFollow(user.ID, targetID)
}

func (s *UserCommandServiceImpl) Block(user *entity.User, blockedUserID int) error {
	if user.IsSelfID(blockedUserID) {
		return serror.New(nil, serror.CodeInvalidParam, "can't block self")
	}

	isFollowMap, err := s.UserQueryRepository.IsFollowing(user.ID, []int{blockedUserID})
	if err != nil {
		return errors.Wrap(err, "failed ref user_blocking")
	}

	// フォローしている場合はフォロー解除する
	if isFollowMap[blockedUserID] {
		if err := s.UserCommandRepository.DeleteFollow(user.ID, blockedUserID); err != nil {
			return errors.Wrap(err, "failed del user_following")
		}
	}

	return s.UserCommandRepository.StoreUserBlock(entity.NewUserBlock(user.ID, blockedUserID))
}

func (s *UserCommandServiceImpl) Unblock(user *entity.User, blockedUserID int) error {
	if user.IsSelfID(blockedUserID) {
		return serror.New(nil, serror.CodeInvalidParam, "can't unblock self")
	}
	return s.UserCommandRepository.DeleteUserBlock(entity.NewUserBlock(user.ID, blockedUserID))
}

func (s *UserCommandServiceImpl) Update(user *entity.User, cmd *command.UpdateUser) error {
	s.updateUserByCmd(user, cmd)

	if err := s.UserCommandRepository.Update(user); err != nil {
		return errors.Wrap(err, "failed to update user")
	}

	// 画像だけは更新が無い場合はputされない
	if err := s.persistUserImage(cmd); err != nil {
		return errors.Wrap(err, "failed to persist user image")
	}

	return nil
}

func (s *UserCommandServiceImpl) UpdateDeviceToken(user *entity.User, deviceToken string) error {
	if err := s.UserCommandRepository.UpdateDeviceTokenByID(user.ID, deviceToken); err != nil {
		return errors.Wrap(err, "failed find update device_token")
	}

	return nil
}

func (s *UserCommandServiceImpl) DeleteUserIcon(user *entity.User) error {
	return s.TransactionService.Do(func(ctx context.Context) error {
		if !user.HasIcon() {
			return serror.New(nil, serror.CodeInvalidParam, "not have icon")
		}

		if err := s.UserCommandRepository.UpdateIconUUIDToNull(ctx, user.ID); err != nil {
			return errors.Wrap(err, "failed update icon uuid")
		}

		if err := s.MediaCommandRepository.Delete(model.UserS3Path(user.AvatarUUID)); err != nil {
			return errors.Wrap(err, "failed delete from s3")
		}

		return nil
	})
}

func (s *UserCommandServiceImpl) DeleteUserHeader(user *entity.User) error {
	return s.TransactionService.Do(func(ctx context.Context) error {
		if !user.HasHeader() {
			return serror.New(nil, serror.CodeInvalidParam, "not have header")
		}

		if err := s.UserCommandRepository.UpdateHeaderUUIDToBlank(ctx, user.ID); err != nil {
			return errors.Wrap(err, "failed update header uuid")
		}

		if err := s.MediaCommandRepository.Delete(model.UserS3Path(user.HeaderUUID)); err != nil {
			return errors.Wrap(err, "failed delete from s3")
		}

		return nil
	})
}

func (s *UserCommandServiceImpl) persistUserImage(cmd *command.UpdateUser) error {
	if cmd.IconUUID != "" {
		if err := s.MediaCommandService.PreparePersist(cmd.IconUUID, model.UserS3Path(cmd.IconUUID), model.MediaTypeUserIcon); err != nil {
			return errors.Wrap(err, "failed to persist avatar")
		}
	}

	if cmd.HeaderUUID != "" {
		if err := s.MediaCommandService.PreparePersist(cmd.HeaderUUID, model.UserS3Path(cmd.HeaderUUID), model.MediaTypeUserHeader); err != nil {
			return errors.Wrap(err, "failed to persist header")
		}
	}

	return nil
}

func (s *UserCommandServiceImpl) updateUserByCmd(user *entity.User, cmd *command.UpdateUser) {
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
