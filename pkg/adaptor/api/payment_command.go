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
	PaymentCommandController struct {
		converter.Converters
		service.PaymentCommandService
	}
)

var PaymentCommandControllerSet = wire.NewSet(
	wire.Struct(new(PaymentCommandController), "*"),
)

func (c *PaymentCommandController) ReservePaymentCfReturnGiftReservedTicket(ctx echo.Context, user entity.User) error {
	i := &input.CfReserveRequest{}
	if err := BindAndValidate(ctx, i); err != nil {
		return errors.Wrap(err, "failed bind input")
	}

	if err := c.PaymentCommandService.ReservePaymentCfReturnGift(&user, i.PaymentID.ID, i.CfReturnGiftID, c.ConvertReserveRequestToEntity(i)); err != nil {
		return errors.Wrap(err, "failed reserve payment_cf_return_gift")
	}

	return ctx.NoContent(http.StatusNoContent)
}
