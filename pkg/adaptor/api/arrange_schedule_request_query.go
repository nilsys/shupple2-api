package api

import (
	"net/http"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
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
	resolve, err := c.ArrangeScheduleRequestQueryService.ShowReceiveList(user)
	if err != nil {
		return errors.Wrap(err, "failed show receive arrange schedule request")
	}

	return ctx.JSON(http.StatusOK, c.ConvertArrangeScheduleRequestList2Output(resolve))
}

func (c *ArrangeScheduleRequestQueryController) ShowSendList(ctx echo.Context, user *entity.UserTiny) error {
	resolve, err := c.ArrangeScheduleRequestQueryService.ShowSendList(user)
	if err != nil {
		return errors.Wrap(err, "failed show send arrange schedule request")
	}

	return ctx.JSON(http.StatusOK, c.ConvertArrangeScheduleRequestList2Output(resolve))
}
