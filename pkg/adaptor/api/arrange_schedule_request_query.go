package api

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/uma-co82/shupple2-api/pkg/adaptor/api/converter"
	"github.com/uma-co82/shupple2-api/pkg/application/service"
	"github.com/uma-co82/shupple2-api/pkg/domain/entity"
)

type ArrangeScheduleRequestQueryController struct {
	service.ArrangeScheduleRequestQueryService
	converter.Converters
}

var ArrangeScheduleRequestQueryControllerSet = wire.NewSet(
	wire.Struct(new(ArrangeScheduleRequestQueryController), "*"),
)

func (c *ArrangeScheduleRequestQueryController) ShowReceiveList(ctx echo.Context, user *entity.UserTiny) error {
}
