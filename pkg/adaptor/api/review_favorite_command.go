package api

import (
	"net/http"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
)

type (
	ReviewFavoriteCommandController struct {
		ReviewFavoriteCommandService service.ReviewFavoriteCommandService
	}
)

var ReviewFavoriteCommandControllerSet = wire.NewSet(
	wire.Struct(new(ReviewFavoriteCommandController), "*"),
)

func (c *ReviewFavoriteCommandController) Store(ctx echo.Context, user entity.User) error {
	p := input.StoreFavoriteReviewParam{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrapf(err, "ID is invalid")
	}

	if err := c.ReviewFavoriteCommandService.Store(&user, p.ReviewID); err != nil {
		return errors.Wrap(err, "failed to store")
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (c *ReviewFavoriteCommandController) Delete(ctx echo.Context, user entity.User) error {
	p := input.StoreFavoriteReviewParam{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrapf(err, "ID is invalid")
	}

	if err := c.ReviewFavoriteCommandService.Delete(&user, p.ReviewID); err != nil {
		return errors.Wrap(err, "failed to delete")
	}

	return ctx.NoContent(http.StatusNoContent)
}
