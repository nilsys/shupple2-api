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
	ReviewFavoriteCommandController struct {
		converter.Converters
		ReviewFavoriteCommandService service.ReviewFavoriteCommandService
	}
)

var ReviewFavoriteCommandControllerSet = wire.NewSet(
	wire.Struct(new(ReviewFavoriteCommandController), "*"),
)

func (c *ReviewFavoriteCommandController) Store(ctx echo.Context, user entity.User) error {
	p := input.StoreFavoriteReviewParam{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrap(err, "ID is invalid")
	}

	if err := c.ReviewFavoriteCommandService.Store(&user, p.ReviewID); err != nil {
		return errors.Wrap(err, "failed to store")
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (c *ReviewFavoriteCommandController) Delete(ctx echo.Context, user entity.User) error {
	p := input.StoreFavoriteReviewParam{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrap(err, "ID is invalid")
	}

	if err := c.ReviewFavoriteCommandService.Delete(&user, p.ReviewID); err != nil {
		return errors.Wrap(err, "failed to delete")
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (c *ReviewFavoriteCommandController) FavoriteReviewCommentReply(ctx echo.Context, user entity.User) error {
	p := input.StoreFavoriteReviewReplyParam{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrap(err, "ID is invalid")
	}

	if err := c.ReviewFavoriteCommandService.FavoriteReviewCommentReply(&user, p.ReplyID); err != nil {
		return errors.Wrap(err, "failed to store")
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (c *ReviewFavoriteCommandController) UnFavoriteReviewCommentReply(ctx echo.Context, user entity.User) error {
	p := input.DeleteFavoriteReviewReplyParam{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrap(err, "ID is invalid")
	}

	if err := c.ReviewFavoriteCommandService.UnFavoriteReviewCommentReply(&user, p.ReplyID); err != nil {
		return errors.Wrap(err, "failed to delete")
	}

	return ctx.NoContent(http.StatusNoContent)
}
