package api

import (
	"net/http"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

	"github.com/stayway-corp/stayway-media-api/pkg/application/service"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
)

type (
	VlogFavoriteCommandController struct {
		service.VlogFavoriteCommandService
	}
)

var VlogFavoriteCommandControllerSet = wire.NewSet(
	wire.Struct(new(VlogFavoriteCommandController), "*"),
)

func (c *VlogFavoriteCommandController) Store(ctx echo.Context, user entity.User) error {
	p := input.StoreFavoriteVlogParam{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrap(err, "ID is invalid")
	}

	if err := c.VlogFavoriteCommandService.Store(&user, p.VlogID); err != nil {
		return errors.Wrap(err, "failed to storeFavorite")
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (c *VlogFavoriteCommandController) Delete(ctx echo.Context, user entity.User) error {
	p := input.DeleteFavoriteVlogParam{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrap(err, "ID is invalid")
	}

	if err := c.VlogFavoriteCommandService.Delete(&user, p.VlogID); err != nil {
		return errors.Wrap(err, "failed to deleteFavorite")
	}

	return ctx.NoContent(http.StatusNoContent)
}
