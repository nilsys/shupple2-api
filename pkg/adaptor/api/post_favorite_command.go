package api

import (
	"net/http"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type (
	PostFavoriteCommandController struct {
		PostQueryService           service.PostQueryService
		PostFavoriteCommandService service.PostFavoriteCommandService
	}
)

var PostFavoriteCommandControllerSet = wire.NewSet(
	wire.Struct(new(PostFavoriteCommandController), "*"),
)

func (c *PostFavoriteCommandController) Store(ctx echo.Context) error {
	// userIDは認証過程で取得するので、暫定的に固定値で実装
	var user entity.User
	user.ID = 1

	q := param.StoreFavoritePostParam{}
	if err := BindAndValidate(ctx, &q); err != nil {
		return errors.Wrapf(err, "ID is invalid")
	}

	if err := c.PostFavoriteCommandService.Store(&user, q.PostID); err != nil {
		return errors.Wrap(err, "failed to storeFavorite")
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (c *PostFavoriteCommandController) Delete(ctx echo.Context) error {
	// userIDは認証過程で取得するので、暫定的に固定値で実装
	var user entity.User
	user.ID = 1

	q := param.StoreFavoritePostParam{}
	if err := BindAndValidate(ctx, &q); err != nil {
		return errors.Wrapf(err, "ID is invalid")
	}

	if err := c.PostFavoriteCommandService.Delete(&user, q.PostID); err != nil {
		return errors.Wrap(err, "failed to deleteFavorite")
	}

	return ctx.NoContent(http.StatusNoContent)
}
