package service

import (
	"context"

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

	user := &entity.UserTiny{
		FirebaseID: firebaseID,
		Name:       cmd.Name,
		Email:      cmd.Email,
		Birthdate:  cmd.Birthdate,
		Profile:    cmd.Profile,
		Gender:     cmd.Gender,
	}

	return s.Store(context.Background(), user)
}
