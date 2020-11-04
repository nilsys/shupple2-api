package api

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/uma-co82/shupple2-api/pkg/adaptor/api/converter"
	"github.com/uma-co82/shupple2-api/pkg/application/service"
	"github.com/uma-co82/shupple2-api/pkg/domain/entity"
)

type ArrangeScheduleRequestCommandController struct {
	service.ArrangeScheduleRequestCommandService
	converter.Converters
}

var ArrangeScheduleRequestCommandControllerSet = wire.NewSet(
	wire.Struct(new(ArrangeScheduleRequestCommandController), "*"),
)

func (c *ArrangeScheduleRequestCommandController) Store(ctx echo.Context, user *entity.UserTiny) error {
}
