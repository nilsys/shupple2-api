package service

import (
	"github.com/google/wire"
	"github.com/uma-co82/shupple2-api/pkg/domain/repository"
)

type (
	ArrangeScheduleRequestQueryService interface {
		//Store(cmd *command.StoreArrangeScheduleRequest, user *entity.UserTiny) error
	}

	ArrangeScheduleRequestQueryServiceImpl struct {
		repository.UserQueryRepository
		repository.ArrangeScheduleRequestCommandRepository
	}
)

var ArrangeScheduleRequestCommandServiceSet = wire.NewSet(
	wire.Struct(new(ArrangeScheduleRequestCommandServiceImpl), "*"),
	wire.Bind(new(ArrangeScheduleRequestCommandService), new(*ArrangeScheduleRequestCommandServiceImpl)),
)
