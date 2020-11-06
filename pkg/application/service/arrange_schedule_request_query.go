package service

import (
	"github.com/google/wire"
	"github.com/uma-co82/shupple2-api/pkg/domain/repository"
)

type (
	ArrangeScheduleRequestQueryService interface {
	}

	ArrangeScheduleRequestQueryServiceImpl struct {
		repository.UserQueryRepository
		repository.ArrangeScheduleRequestCommandRepository
	}
)

var ArrangeScheduleRequestQueryServiceSet = wire.NewSet(
	wire.Struct(new(ArrangeScheduleRequestQueryServiceImpl), "*"),
	wire.Bind(new(ArrangeScheduleRequestQueryService), new(*ArrangeScheduleRequestQueryServiceImpl)),
)
