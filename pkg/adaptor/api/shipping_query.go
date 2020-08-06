package api

import (
	"net/http"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type ShippingQueryController struct {
	service.ShippingQueryService
	converter.Converters
}

var ShippingQueryControllerSet = wire.NewSet(
	wire.Struct(new(ShippingQueryController), "*"),
)

func (c *ShippingQueryController) Show(ctx echo.Context, user entity.User) error {
	address, err := c.ShippingQueryService.ShowShippingAddressByUserID(&user)
	if err != nil {
		return errors.Wrap(err, "failed show shopping address")
	}
	return ctx.JSON(http.StatusOK, c.ConvertShippingAddressToOutput(address))
}
