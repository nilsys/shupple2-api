package api

import (
	"net/http"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/uma-co82/shupple2-api/pkg/adaptor/api/converter"
	"github.com/uma-co82/shupple2-api/pkg/adaptor/api/input"
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
	i := &input.StoreArrangeScheduleRequest{}
	if err := BindAndValidate(ctx, i); err != nil {
		return errors.Wrap(err, "failed bind input")
	}

	if err := c.ArrangeScheduleRequestCommandService.Store(c.ConvertStoreArrangeScheduleRequestInputToCmd(i), user); err != nil {
		return errors.Wrap(err, "failed store arrange schedule request")
	}

	return ctx.NoContent(http.StatusNoContent)
}
