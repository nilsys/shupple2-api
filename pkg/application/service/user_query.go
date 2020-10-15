package service

import (
	"github.com/google/wire"
	"github.com/uma-co82/shupple2-api/pkg/domain/entity"
	"github.com/uma-co82/shupple2-api/pkg/domain/repository"
)

type (
	UserQueryService interface {
		Show(id int) (*entity.UserTiny, error)
	}

	UserQueryServiceImpl struct {
		repository.UserQueryRepository
	}
)

var UserQueryServiceSet = wire.NewSet(
	wire.Struct(new(UserQueryServiceImpl), "*"),
	wire.Bind(new(UserQueryService), new(*UserQueryServiceImpl)),
)

func (s *UserQueryServiceImpl) Show(id int) (*entity.UserTiny, error) {
	return s.UserQueryRepository.FindByID(id)
}
