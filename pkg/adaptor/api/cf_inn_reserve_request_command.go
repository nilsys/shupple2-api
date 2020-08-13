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

type (
	CfInnReserveRequestCommandController struct {
		converter.Converters
		service.CfInnReserveRequestCommandService
	}
)

var CfReserveRequestCommandControllerSet = wire.NewSet(
	wire.Struct(new(CfInnReserveRequestCommandController), "*"),
)

func (c *CfInnReserveRequestCommandController) RequestReserve(ctx echo.Context, user entity.User) error {
	i := &input.CfInnReserveRequest{}
	if err := BindAndValidate(ctx, i); err != nil {
		return errors.Wrap(err, "failed bind input")
	}

	if err := c.CfInnReserveRequestCommandService.RequestReserve(&user, i.PaymentID.ID, i.CfReturnGiftID, c.ConvertReserveRequestToEntity(i, user.ID)); err != nil {
		return errors.Wrap(err, "failed request reserve ")
	}

	return ctx.NoContent(http.StatusNoContent)
}
