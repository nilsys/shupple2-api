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
	PaymentQueryController struct {
		converter.Converters
		service.PaymentQueryService
	}
)

var PaymentQueryControllerSet = wire.NewSet(
	wire.Struct(new(PaymentQueryController), "*"),
)

func (c *PaymentQueryController) List(ctx echo.Context, user entity.User) error {
	i := &input.ListPayment{}
	if err := BindAndValidate(ctx, i); err != nil {
		return errors.Wrap(err, "failed bind input")
	}

	payments, err := c.PaymentQueryService.ListByUser(&user, i.CfProjectID, c.ConvertListPaymentToQuery(i))
	if err != nil {
		return errors.Wrap(err, "failed list payment")
	}

	return ctx.JSON(http.StatusOK, c.ConvertPaymentListToOutput(payments))
}
