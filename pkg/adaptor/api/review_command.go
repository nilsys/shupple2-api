package api

import (
	"net/http"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/application/scenario"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type ReviewCommandController struct {
	scenario.ReviewCommandScenario
	service.ReviewCommandService
}

var ReviewCommandControllerSet = wire.NewSet(
	wire.Struct(new(ReviewCommandController), "*"),
)

func (c *ReviewCommandController) Store(ctx echo.Context, user entity.User) error {
	p := &param.StoreReviewParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation store review param")
	}

	if err := c.ReviewCommandScenario.Create(&user, converter.ConvertCreateReviewParamToCommand(p)); err != nil {
		return errors.Wrap(err, "failed to store review")
	}

	return ctx.JSON(http.StatusOK, "ok")
}

func (c *ReviewCommandController) Update(ctx echo.Context, user entity.User) error {
	p := &param.UpdateReviewParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation update review param")
	}

	if err := c.ReviewCommandScenario.UpdateReview(&user, converter.ConvertUpdateReviewParamToCommand(p)); err != nil {
		return errors.Wrap(err, "failed to update review")
	}

	return ctx.JSON(http.StatusOK, "ok")
}

func (c *ReviewCommandController) StoreReviewComment(ctx echo.Context, user entity.User) error {
	reviewCommentParam := &param.CreateReviewComment{}
	if err := BindAndValidate(ctx, reviewCommentParam); err != nil {
		return errors.Wrap(err, "Failed to bind parameters")
	}

	if err := c.ReviewCommandService.CreateReviewComment(&user, reviewCommentParam.ID, reviewCommentParam.Body); err != nil {
		return errors.Wrap(err, "Failed to create review comment")
	}

	return ctx.JSON(http.StatusOK, "ok")
}

func (c *ReviewCommandController) DeleteReviewComment(ctx echo.Context, user entity.User) error {
	reviewCommentParam := &param.DeleteReviewCommentParam{}
	if err := BindAndValidate(ctx, reviewCommentParam); err != nil {
		return errors.Wrap(err, "Failed to bind parameters")
	}

	if err := c.ReviewCommandService.DeleteReviewComment(&user, reviewCommentParam.ID); err != nil {
		return errors.Wrap(err, "Failed to delete review comment")
	}

	return ctx.JSON(http.StatusOK, nil)
}

func (c *ReviewCommandController) StoreReviewCommentReply(ctx echo.Context, user entity.User) error {
	p := &param.CreateReviewCommentReply{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation store review comment reply param")
	}

	if err := c.ReviewCommandService.CreateReviewCommentReply(&user, converter.ConvertCreateReviewCommentReplyParamToCommand(p)); err != nil {
		return errors.Wrap(err, "failed to store review comment reply")
	}

	return ctx.JSON(http.StatusOK, "ok")
}

func (c *ReviewCommandController) FavoriteReviewComment(ctx echo.Context, user entity.User) error {
	p := &param.FavoriteReviewComment{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation favorite review comment param")
	}

	if err := c.ReviewCommandService.FavoriteReviewComment(&user, p.ReviewCommentID); err != nil {
		return errors.Wrap(err, "failed to favorite review_comment")
	}

	return ctx.JSON(http.StatusOK, "ok")
}

func (c *ReviewCommandController) Delete(ctx echo.Context, user entity.User) error {
	p := &param.DeleteReviewParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation delete review param")
	}

	if err := c.ReviewCommandScenario.DeleteReview(p.ID, &user); err != nil {
		return errors.Wrap(err, "failed to delete review")
	}

	return ctx.JSON(http.StatusOK, "ok")
}

func (c *ReviewCommandController) UnfavoriteReviewComment(ctx echo.Context, user entity.User) error {
	p := &param.FavoriteReviewComment{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation un favorite review comment param")
	}

	if err := c.ReviewCommandService.UnfavoriteReviewComment(&user, p.ReviewCommentID); err != nil {
		return errors.Wrap(err, "failed to un favorite review_comment")
	}

	return ctx.JSON(http.StatusOK, "ok")
}
