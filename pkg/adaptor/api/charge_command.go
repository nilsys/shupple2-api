package api

import (
	"net/http"

	"github.com/stayway-corp/stayway-media-api/pkg/application/scenario"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type ChargeCommandController struct {
	scenario.ChargeCommandScenario
	service.ChargeCommandService
	converter.Converters
}

var ChargeCommandControllerSet = wire.NewSet(
	wire.Struct(new(ChargeCommandController), "*"),
)

func (c *ChargeCommandController) Create(ctx echo.Context, user entity.User) error {
	i := &input.CreateCharge{}
	if err := BindAndValidate(ctx, i); err != nil {
		return errors.Wrap(err, "invalid request body")
	}

	resolve, err := c.ChargeCommandScenario.Create(&user, c.ConvertCaptureChargeToCmd(i))
	if err != nil {
		return errors.Wrap(err, "failed capture charge")
	}

	return ctx.JSON(http.StatusOK, resolve)
}

func (c *ChargeCommandController) InstantCreate(ctx echo.Context) error {
	i := &input.InstantCreateCharge{}
	if err := BindAndValidate(ctx, i); err != nil {
		return errors.Wrap(err, "invalid input")
	}

	resolve, err := c.ChargeCommandScenario.InstantCreate(c.ConvertCaptureChargeToCmd(&i.CreateCharge), i.CardToken, c.ConvertShippingAddressInputToEntity(0, i.ShippingAddress))
	if err != nil {
		return errors.Wrap(err, "failed capture charge")
	}

	return ctx.JSON(http.StatusOK, resolve)
}

func (c *ChargeCommandController) Refund(ctx echo.Context, user entity.User) error {
	i := &input.RefundCharge{}
	if err := BindAndValidate(ctx, i); err != nil {
		return errors.Wrap(err, "failed bind input")
	}

	if err := c.ChargeCommandService.Refund(&user, i.ID, i.CfReturnGiftID); err != nil {
		return errors.Wrap(err, "failed refund charge")
	}

	return ctx.NoContent(http.StatusNoContent)
}
