package api

import (
	"net/http"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
)

type (
	ReviewFavoriteCommandController struct {
		ReviewQueryService           service.ReviewQueryService
		ReviewFavoriteCommandService service.ReviewFavoriteCommandService
	}
)

var ReviewFavoriteCommandControllerSet = wire.NewSet(
	wire.Struct(new(ReviewFavoriteCommandController), "*"),
)

func (c *ReviewFavoriteCommandController) Store(ctx echo.Context) error {
	// userIDは認証過程で取得するので、暫定的に固定値で実装
	var user entity.User
	user.ID = 1

	q := param.StoreFavoriteReviewParam{}
	if err := BindAndValidate(ctx, &q); err != nil {
		return errors.Wrapf(err, "ID is invalid")
	}

	if err := c.ReviewFavoriteCommandService.Store(&user, q.ReviewID); err != nil {
		return errors.Wrap(err, "failed to store")
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (c *ReviewFavoriteCommandController) Delete(ctx echo.Context) error {
	// userIDは認証過程で取得するので、暫定的に固定値で実装
	var user entity.User
	user.ID = 1

	q := param.StoreFavoriteReviewParam{}
	if err := BindAndValidate(ctx, &q); err != nil {
		return errors.Wrapf(err, "ID is invalid")
	}

	if err := c.ReviewFavoriteCommandService.Delete(&user, q.ReviewID); err != nil {
		return errors.Wrap(err, "failed to store")
	}

	return ctx.NoContent(http.StatusNoContent)
}
