package api

import (
	"net/http"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
)

type (
	ComicFavoriteCommandController struct {
		service.ComicFavoriteCommandService
	}
)

var ComicFavoriteCommandControllerSet = wire.NewSet(
	wire.Struct(new(ComicFavoriteCommandController), "*"),
)

func (c *ComicFavoriteCommandController) Store(ctx echo.Context, user entity.User) error {
	p := input.FavoriteComicParam{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrap(err, "ID is invalid")
	}

	if err := c.ComicFavoriteCommandService.Store(&user, p.ComicID); err != nil {
		return errors.Wrap(err, "failed to storeFavorite")
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (c *ComicFavoriteCommandController) Delete(ctx echo.Context, user entity.User) error {
	p := input.FavoriteComicParam{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrap(err, "ID is invalid")
	}

	if err := c.ComicFavoriteCommandService.Delete(&user, p.ComicID); err != nil {
		return errors.Wrap(err, "failed to deleteFavorite")
	}

	return ctx.NoContent(http.StatusNoContent)
}
