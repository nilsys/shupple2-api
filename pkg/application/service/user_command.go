package service

import (
	"context"

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
	}

	UserCommandServiceImpl struct {
		repository.UserQueryRepository
		repository.UserCommandRepository
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

	return s.Store(context.Background(), user)
}

func (s *UserCommandServiceImpl) Matching(user *entity.UserTiny) error {
	matchingUser, err := s.UserQueryRepository.FindAvailableMatchingUser(user.Gender, user.MatchingReason, user.ID)
	if err != nil {
		return errors.Wrap(err, "failed find available matching user")
	}

	return s.Do(func(ctx context.Context) error {
		history := entity.NewUserMatchingHistory(user.ID, matchingUser.ID)
		matchingUserHistory := entity.NewUserMatchingHistory(matchingUser.ID, user.ID)

		if err := s.UserCommandRepository.StoreUserMatchingHistory(ctx, history); err != nil {
			return errors.Wrap(err, "failed store user_matching_history")
		}
		if err := s.UserCommandRepository.StoreUserMatchingHistory(ctx, matchingUserHistory); err != nil {
			return errors.Wrap(err, "failed store user_matching_history")
		}

		return nil
	})
}
