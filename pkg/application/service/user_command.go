package service

import (
	"context"
	"time"

	"github.com/uma-co82/shupple2-api/pkg/config"

	"github.com/uma-co82/shupple2-api/pkg/util"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	shuppleAWS "github.com/uma-co82/shupple2-api/pkg/domain/repository/aws"

	"github.com/pkg/errors"

	"github.com/google/wire"

	"github.com/uma-co82/shupple2-api/pkg/domain/entity"
	"github.com/uma-co82/shupple2-api/pkg/domain/model/command"
	"github.com/uma-co82/shupple2-api/pkg/domain/model/serror"
	"github.com/uma-co82/shupple2-api/pkg/domain/repository"
)

type (
	UserCommandService interface {
		SignUp(cmd command.StoreUser, firebaseToken string) error
		Matching(user *entity.UserTiny) error
		ApproveMainMatching(user *entity.User, matchingUserID int, isApprove bool) error
		StoreUserImage(cmd command.StoreUserImage, user *entity.UserTiny) error
		DeleteUserImage(imageUUID string, user *entity.UserTiny) error
	}

	UserCommandServiceImpl struct {
		repository.UserQueryRepository
		repository.UserCommandRepository
		shuppleAWS.S3CommandRepository
		AWSConfig config.AWS
		AuthService
		TransactionService
	}
)

var UserCommandServiceSet = wire.NewSet(
	wire.Struct(new(UserCommandServiceImpl), "*"),
	wire.Bind(new(UserCommandService), new(*UserCommandServiceImpl)),
)

func (s *UserCommandServiceImpl) SignUp(cmd command.StoreUser, firebaseToken string) error {
	firebaseID, err := s.AuthService.Authorize(firebaseToken)
	if err != nil {
		return serror.New(err, serror.CodeUnauthorized, "unauthorized")
	}

	user := entity.NewUserTinyFromCmd(cmd, firebaseID)

	return s.TransactionService.Do(func(ctx context.Context) error {
		if err := s.UserCommandRepository.Store(ctx, user); err != nil {
			return errors.Wrap(err, "failed store user")
		}
		return nil
	})
}

func (s *UserCommandServiceImpl) StoreUserImage(cmd command.StoreUserImage, user *entity.UserTiny) error {
	image, err := entity.NewUserImage(cmd, user.ID)
	if err != nil {
		return errors.Wrap(err, "failed new user images")
	}

	return s.TransactionService.Do(func(ctx context.Context) error {
		if err := s.UserCommandRepository.StoreUserImages(ctx, image); err != nil {
			return errors.Wrap(err, "failed store user_image")
		}

		if err := s.uploadUserImage(cmd, image); err != nil {
			return errors.Wrap(err, "failed upload user image")
		}

		return nil
	})
}

func (s *UserCommandServiceImpl) DeleteUserImage(imageUUID string, user *entity.UserTiny) error {
	image, err := s.UserQueryRepository.FindImageByUUID(imageUUID)
	if err != nil {
		return errors.Wrap(err, "failed find user_image")
	}

	if image.UserID != user.ID {
		return serror.New(nil, serror.CodeForbidden, "forbidden")
	}

	return s.S3CommandRepository.Delete(image.S3Path())
}

func (s *UserCommandServiceImpl) Matching(user *entity.UserTiny) error {
	matchingUser, err := s.UserQueryRepository.FindAvailableMatchingUser(user.Gender, user.MatchingReason, user.ID)
	if err != nil {
		return errors.Wrap(err, "failed find available matching user")
	}

	return s.Do(func(ctx context.Context) error {
		matchedAt := time.Now()
		history := entity.NewUserMatchingHistory(user.ID, matchingUser.ID, matchedAt)
		matchingUserHistory := entity.NewUserMatchingHistory(matchingUser.ID, user.ID, matchedAt)

		if err := s.UserCommandRepository.StoreUserMatchingHistory(ctx, history); err != nil {
			return errors.Wrap(err, "failed store user_matching_history")
		}
		if err := s.UserCommandRepository.StoreUserMatchingHistory(ctx, matchingUserHistory); err != nil {
			return errors.Wrap(err, "failed store user_matching_history")
		}

		if err := s.UserCommandRepository.UpdateIsMatchingToTrueByIDs(ctx, []int{user.ID, matchingUser.ID}); err != nil {
			return errors.Wrap(err, "failed update user is_matching")
		}

		return nil
	})
}

func (s *UserCommandServiceImpl) ApproveMainMatching(user *entity.User, matchingUserID int, isApprove bool) error {
	history, err := s.UserQueryRepository.FindMatchingHistoryByUserIDAndMatchingUserID(user.ID, matchingUserID)
	if err != nil {
		return errors.Wrap(err, "failed find user_matching_history")
	}

	if !history.IsExpired() {
		return serror.New(nil, serror.CodeMatchingNotExpired, "matching is not expired")
	}

	// 既に評価済みの場合
	if history.UserMainMatchingApprove.Valid {
		return serror.New(nil, serror.CodeInvalidParam, "duplicate confirm")
	}

	return s.TransactionService.Do(func(ctx context.Context) error {
		if err := s.UserCommandRepository.UpdateUserMatchingHistoryUserMainMatchingApprove(ctx, user.ID, matchingUserID, isApprove); err != nil {
			return errors.Wrap(err, "failed update user_matching_history.user_main_matching_approve")
		}
		if err := s.UserCommandRepository.UpdateUserMatchingHistoryMatchingUserMainMatchingApprove(ctx, matchingUserID, user.ID, isApprove); err != nil {
			return errors.Wrap(err, "failed update user_matching_history.matching_user_main_matching_approve")
		}
		return nil
	})
}

func (s *UserCommandServiceImpl) uploadUserImage(cmd command.StoreUserImage, image *entity.UserImage) error {
	body, err := util.Base64StrWriteBuffer(cmd.ImageBase64)
	if err != nil {
		return errors.Wrap(err, "failed write buffer")
	}
	uploadInput := &s3manager.UploadInput{
		ACL:         aws.String(s3.ObjectCannedACLPublicRead),
		Body:        body,
		Bucket:      aws.String(s.AWSConfig.FilesBucket),
		Key:         aws.String(image.S3Path()),
		ContentType: aws.String(image.MimeType),
	}
	if err := s.S3CommandRepository.Upload(uploadInput); err != nil {
		return errors.Wrap(err, "failed upload to s3")
	}
	return nil
}
