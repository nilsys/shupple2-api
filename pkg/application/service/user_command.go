package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	UserCommandService interface {
		SignUp(user *entity.User, cognitoID string) error
	}

	UserCommandServiceImpl struct {
		repository.UserCommandRepository
		repository.UserQueryRepository
	}
)

var UserCommandServiceSet = wire.NewSet(
	wire.Struct(new(UserCommandServiceImpl), "*"),
	wire.Bind(new(UserCommandService), new(*UserCommandServiceImpl)),
)

func (s *UserCommandServiceImpl) SignUp(user *entity.User, cognitoID string) error {
	isExist, err := s.UserQueryRepository.IsExistByUID(user.UID)
	if err != nil {
		return errors.Wrap(err, "failed to get user")
	}
	if isExist {
		return serror.New(nil, serror.CodeInvalidParam, "uid: %s is duplicate", user.UID)
	}

	if err := s.UserCommandRepository.Store(user); err != nil {
		return errors.Wrap(err, "failed to store user")
	}
	return nil
}
