package service

import (
	"github.com/uma-co82/shupple2-api/pkg/domain/entity"
	"github.com/uma-co82/shupple2-api/pkg/domain/model/command"
	"github.com/uma-co82/shupple2-api/pkg/domain/repository"
)

type (
	ArrangeScheduleRequestCommandService interface {
		Store(cmd *command.StoreArrangeScheduleRequest, user *entity.UserTiny) error
	}

	ArrangeScheduleRequestCommandServiceImpl struct {
		repository.ArrangeScheduleRequestCommandRepository
	}
)

func (s *ArrangeScheduleRequestCommandServiceImpl) Store(cmd *command.StoreArrangeScheduleRequest, user *entity.UserTiny) error {
	req := entity.ArrangeScheduleRequest{}
}
