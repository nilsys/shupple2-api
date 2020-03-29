package service

import (
	"strconv"
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/command"

	"github.com/google/wire"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity/wordpress"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
	"gopkg.in/guregu/null.v3"
)

type (
	UserCommandService interface {
		SignUp(user *entity.User, cognitoToken string, migrationCode *string) error
		Update(user *entity.User, cmd *command.UpdateUser) error
		RegisterWordpressUser(wordpressUserID int) error
		Follow(user *entity.User, targetID int) error
		Unfollow(user *entity.User, targetID int) error
	}

	UserCommandServiceImpl struct {
		repository.UserCommandRepository
		repository.UserQueryRepository
		repository.WordpressQueryRepository
		AuthService
	}
)

var UserCommandServiceSet = wire.NewSet(
	wire.Struct(new(UserCommandServiceImpl), "*"),
	wire.Bind(new(UserCommandService), new(*UserCommandServiceImpl)),
)

func (s *UserCommandServiceImpl) SignUp(user *entity.User, cognitoToken string, migrationCode *string) error {
	isExist, err := s.UserQueryRepository.IsExistByUID(user.UID)
	if err != nil {
		return errors.Wrap(err, "failed to get user")
	}
	if isExist {
		return serror.New(nil, serror.CodeInvalidParam, "uid: %s is duplicate", user.UID)
	}

	cognitoID, err := s.AuthService.Authorize(cognitoToken)
	if err != nil {
		return serror.New(err, serror.CodeUnauthorized, "unauthorized")
	}
	user.CognitoID = cognitoID

	if migrationCode != nil && *migrationCode != "" {
		existingUser, err := s.UserQueryRepository.FindByMigrationCode(*migrationCode)
		if err != nil {
			return errors.Wrap(err, "failed to get user by migration code")
		}
		user.ID = existingUser.ID
	}

	if err := s.UserCommandRepository.Store(user); err != nil {
		return errors.Wrap(err, "failed to store user")
	}
	return nil
}

// TODO: エラー時はslackに通知飛ばしたほうが良さそう
func (s *UserCommandServiceImpl) RegisterWordpressUser(wordpressUserID int) error {
	// すでに紐づけされているユーザーが居る場合はエラー
	// Wordpress側のユーザー登録時にしか叩かれないので、ありえないはずだが一応チェック
	_, err := s.UserQueryRepository.FindByWordpressID(wordpressUserID)
	if !serror.IsErrorCode(err, serror.CodeNotFound) {
		if err == nil {
			err = serror.New(nil, serror.CodeInvalidParam, "existing wordpress user; id=%d", wordpressUserID)
		}
		return errors.Wrap(err, "failed to register wordpress user")
	}

	wpUsers, err := s.WordpressQueryRepository.FindUsersByIDs([]int{wordpressUserID})
	if err != nil || len(wpUsers) == 0 {
		return serror.NewResourcesNotFoundError(err, "wordpress user(id=%d)", wordpressUserID)
	}
	wpUser := wpUsers[0]

	if wpUser.Attributes.MediaUserID != "" {
		return s.updateMapping(wpUser)
	}

	return s.registerDummyUserForWordpress(wpUser)
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

func (s *UserCommandServiceImpl) registerDummyUserForWordpress(wpUser *wordpress.User) error {
	avatar, err := s.WordpressQueryRepository.DownloadAvatar(wpUser.AvatarURLs.Num96)
	if err != nil {
		return errors.Wrap(err, "failed to download avatar")
	}

	user := &entity.User{
		UID:           wpUser.Slug,
		Name:          wpUser.Name,
		MigrationCode: null.StringFrom(uuid.NewV4().String()),
		WordpressID:   null.IntFrom(int64(wpUser.ID)),
		Profile:       wpUser.Description,
		Birthdate:     time.Date(1900, 1, 1, 0, 0, 0, 0, time.Local),
	}

	return errors.Wrap(s.UserCommandRepository.StoreWithAvatar(user, avatar), "failed to register dummy user")
}

func (s *UserCommandServiceImpl) Follow(user *entity.User, targetID int) error {
	if user.IsSelfID(targetID) {
		return serror.New(nil, serror.CodeInvalidParam, "can not follow self")
	}
	following := entity.NewUserFollowing(user.ID, targetID)
	followed := entity.NewUserFollowed(targetID, user.ID)

	return s.UserCommandRepository.StoreFollow(following, followed)
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

	if err := s.persistUserImage(user); err != nil {
		return errors.Wrap(err, "failed to persist user image")
	}

	return nil
}

func (s *UserCommandServiceImpl) persistUserImage(user *entity.User) error {
	if err := s.UserCommandRepository.PersistUserImage(user); err != nil {
		return errors.Wrapf(err, "failed to persist user(id=%s) image", user.ID)
	}
	return nil
}

func (s *UserCommandServiceImpl) updateUserCmd(user *entity.User, cmd *command.UpdateUser) {
	user.Name = cmd.Name
	user.Email = cmd.Email
	user.Birthdate = time.Time(cmd.BirthDate)
	user.Gender = cmd.Gender
	user.Profile = cmd.Profile
	user.AvatarUUID = cmd.IconUUID
	user.HeaderUUID = cmd.HeaderUUID
	user.URL = cmd.URL
	user.FacebookURL = cmd.FacebookURL
	user.InstagramURL = cmd.InstagramURL
	user.TwitterURL = cmd.TwitterURL
	user.LivingArea = cmd.LivingArea
	user.Interests = cmd.Interests
}
