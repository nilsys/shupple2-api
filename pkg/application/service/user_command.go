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

	user, err := entity.NewUser(cmd, firebaseID)
	if err != nil {
		return errors.Wrap(err, "failed new user")
	}

	return s.TransactionService.Do(func(ctx context.Context) error {

		if err := s.Store(ctx, &user.UserTiny); err != nil {
			return errors.Wrap(err, "failed store user")
		}

		// user.idはここで初めて取得できるので、画像のstoreの前にidをいれる
		user.InsertUserID2Images()

		if err := s.StoreUserImages(ctx, user.Images); err != nil {
			return errors.Wrap(err, "failed store user_image")
		}

		if err := s.uploadUserImage(cmd.Images, user.Images); err != nil {
			return errors.Wrap(err, "failed upload user image")
		}

		return nil
	})
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

func (s *UserCommandServiceImpl) uploadUserImage(cmd []command.StoreUserImage, images []*entity.UserImage) error {
	for i, image := range cmd {
		body, err := util.Base64StrWriteBuffer(image.ImageBase64)
		if err != nil {
			return errors.Wrap(err, "failed write buffer")
		}
		uploadInput := &s3manager.UploadInput{
			ACL:         aws.String(s3.ObjectCannedACLPublicRead),
			Body:        body,
			Bucket:      aws.String(s.AWSConfig.FilesBucket),
			Key:         aws.String(images[i].S3Path()),
			ContentType: aws.String(image.MimeType),
		}
		if err := s.S3CommandRepository.Upload(uploadInput); err != nil {
			return errors.Wrap(err, "failed upload to s3")
		}
	}
	return nil
}
