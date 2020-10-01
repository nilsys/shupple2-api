package api

import (
	"net/http"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
)

type (
	CfReturnGiftQueryController struct {
		converter.Converters
		service.CfReturnGiftQueryService
	}
)

var CfReturnGiftQueryControllerSet = wire.NewSet(
	wire.Struct(new(CfReturnGiftQueryController), "*"),
)

func (c *CfReturnGiftQueryController) List(ctx echo.Context) error {
	i := &input.ListCfReturnGift{}
	if err := BindAndValidate(ctx, i); err != nil {
		return errors.Wrap(err, "failed bind input")
	}

	gifts, err := c.CfReturnGiftQueryService.ListByCfProjectID(c.ConvertListCfReturnGiftInputToQuery(i))
	if err != nil {
		return errors.Wrap(err, "failed list cf_return_gift")
	}

	return ctx.JSON(http.StatusOK, c.ConvertCfReturnGiftWithCountListToOutput(gifts))
}
