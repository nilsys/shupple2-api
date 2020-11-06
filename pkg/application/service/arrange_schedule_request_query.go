package service

import (
	"github.com/google/wire"
	"github.com/uma-co82/shupple2-api/pkg/domain/entity"
	"github.com/uma-co82/shupple2-api/pkg/domain/repository"
)

type (
	ArrangeScheduleRequestQueryService interface {
		ShowReceiveList(user *entity.UserTiny) ([]*entity.ArrangeScheduleRequest, error)
		ShowSendList(user *entity.UserTiny) ([]*entity.ArrangeScheduleRequest, error)
	}

	ArrangeScheduleRequestQueryServiceImpl struct {
		repository.ArrangeScheduleRequestQueryRepository
	}
)

var ArrangeScheduleRequestQueryServiceSet = wire.NewSet(
	wire.Struct(new(ArrangeScheduleRequestQueryServiceImpl), "*"),
	wire.Bind(new(ArrangeScheduleRequestQueryService), new(*ArrangeScheduleRequestQueryServiceImpl)),
)

func (s *ArrangeScheduleRequestQueryServiceImpl) ShowReceiveList(user *entity.UserTiny) ([]*entity.ArrangeScheduleRequest, error) {
	return s.ArrangeScheduleRequestQueryRepository.FindByMatchingUserID(user.ID)
}

func (s *ArrangeScheduleRequestQueryServiceImpl) ShowSendList(user *entity.UserTiny) ([]*entity.ArrangeScheduleRequest, error) {
	return s.ArrangeScheduleRequestQueryRepository.FindByMatchingUserID(user.ID)
}
