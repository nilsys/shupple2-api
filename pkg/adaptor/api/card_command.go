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

type CardCommandController struct {
	service.CardCommandService
	converter.Converters
}

var CardCommandControllerSet = wire.NewSet(
	wire.Struct(new(CardCommandController), "*"),
)

func (c *CardCommandController) Register(ctx echo.Context, user entity.User) error {
	i := &input.StoreCard{}
	if err := BindAndValidate(ctx, i); err != nil {
		return errors.Wrap(err, "validation card request body")
	}
	if err := c.CardCommandService.Register(&user, i.CardToken); err != nil {
		return errors.Wrap(err, "failed register card")
	}

	return ctx.JSON(http.StatusOK, "ok")
}

func (c *CardCommandController) Delete(ctx echo.Context, user entity.User) error {
	i := &input.IDParam{}
	if err := BindAndValidate(ctx, i); err != nil {
		return errors.Wrap(err, "failed bind input")
	}

	if err := c.CardCommandService.Delete(&user, i.ID); err != nil {
		return errors.Wrap(err, "failed delete card")
	}

	return ctx.NoContent(http.StatusNoContent)
}
