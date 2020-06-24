package api

import (
	"net/http"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type ChargeCommandController struct {
	service.ChargeCommandService
	converter.Converters
}

var ChargeCommandControllerSet = wire.NewSet(
	wire.Struct(new(ChargeCommandController), "*"),
)

func (c *ChargeCommandController) CaptureCharge(ctx echo.Context, user entity.User) error {
	i := &input.CaptureCharge{}
	if err := BindAndValidate(ctx, i); err != nil {
		return errors.Wrap(err, "invalid request body")
	}

	if err := c.ChargeCommandService.CaptureCharge(&user, c.ConvertPaymentsToCmd(i)); err != nil {
		return errors.Wrap(err, "failed capture charge")
	}

	return ctx.JSON(http.StatusOK, "ok")
}
