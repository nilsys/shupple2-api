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
	CfProjectCommandController struct {
		converter.Converters
		service.CfProjectCommandService
	}
)

var CfProjectCommandControllerSet = wire.NewSet(
	wire.Struct(new(CfProjectCommandController), "*"),
)

func (c *CfProjectCommandController) Favorite(ctx echo.Context, user entity.User) error {
	i := &input.IDParam{}
	if err := BindAndValidate(ctx, i); err != nil {
		return errors.Wrap(err, "failed bind input")
	}

	if err := c.CfProjectCommandService.Favorite(&user, i.ID); err != nil {
		return errors.Wrap(err, "failed favorite")
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (c *CfProjectCommandController) Unfavorite(ctx echo.Context, user entity.User) error {
	i := &input.IDParam{}
	if err := BindAndValidate(ctx, i); err != nil {
		return errors.Wrap(err, "failed bind input")
	}

	if err := c.CfProjectCommandService.Unfavorite(&user, i.ID); err != nil {
		return errors.Wrap(err, "failed unfavorite")
	}

	return ctx.NoContent(http.StatusNoContent)
}
