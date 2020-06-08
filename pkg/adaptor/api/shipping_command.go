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

type ShippingCommandController struct {
	service.ShippingCommandService
	converter.Converters
}

var ShippingCommandControllerSet = wire.NewSet(
	wire.Struct(new(ShippingCommandController), "*"),
)

func (c *ShippingCommandController) StoreShippingAddress(ctx echo.Context, user entity.User) error {
	i := &input.ShippingAddress{}
	if err := BindAndValidate(ctx, i); err != nil {
		return errors.Wrap(err, "validation shipping address request body")
	}

	if err := c.ShippingCommandService.StoreShippingAddress(&user, c.ConvertShippingAddressInputToEntity(user.ID, i)); err != nil {
		return errors.Wrap(err, "failed store shipping address")
	}

	return ctx.JSON(http.StatusOK, "ok")
}
