package api

import (
	"net/http"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
)

type (
	PostFavoriteCommandController struct {
		converter.Converters
		PostFavoriteCommandService service.PostFavoriteCommandService
	}
)

var PostFavoriteCommandControllerSet = wire.NewSet(
	wire.Struct(new(PostFavoriteCommandController), "*"),
)

func (c *PostFavoriteCommandController) Store(ctx echo.Context, user entity.User) error {
	p := input.StoreFavoritePostParam{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrapf(err, "ID is invalid")
	}

	if err := c.PostFavoriteCommandService.Store(&user, p.PostID); err != nil {
		return errors.Wrap(err, "failed to storeFavorite")
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (c *PostFavoriteCommandController) Delete(ctx echo.Context, user entity.User) error {
	p := input.StoreFavoritePostParam{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrapf(err, "ID is invalid")
	}

	if err := c.PostFavoriteCommandService.Delete(&user, p.PostID); err != nil {
		return errors.Wrap(err, "failed to deleteFavorite")
	}

	return ctx.NoContent(http.StatusNoContent)
}
