package api

import (
	"net/http"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/application/scenario"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type ReviewCommandController struct {
	converter.Converters
	scenario.ReviewCommandScenario
	service.ReviewCommandService
	service.HashtagQueryService
}

var ReviewCommandControllerSet = wire.NewSet(
	wire.Struct(new(ReviewCommandController), "*"),
)

func (c *ReviewCommandController) Store(ctx echo.Context, user entity.User) error {
	p := &input.StoreReviewParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation store review input")
	}

	review, err := c.ReviewCommandScenario.Create(&user, c.ConvertCreateReviewParamToCommand(p))
	if err != nil {
		return errors.Wrap(err, "failed to store review")
	}

	return ctx.JSON(http.StatusOK, c.ConvertQueryReviewDetailWithIsFavoriteToOutput(review))
}

func (c *ReviewCommandController) Update(ctx echo.Context, user entity.User) error {
	p := &input.UpdateReviewParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation update review input")
	}

	if err := c.ReviewCommandScenario.UpdateReview(&user, c.ConvertUpdateReviewParamToCommand(p)); err != nil {
		return errors.Wrap(err, "failed to update review")
	}

	return ctx.JSON(http.StatusOK, "ok")
}

func (c *ReviewCommandController) StoreReviewComment(ctx echo.Context, user entity.User) error {
	reviewCommentParam := &input.CreateReviewComment{}
	if err := BindAndValidate(ctx, reviewCommentParam); err != nil {
		return errors.Wrap(err, "Failed to bind parameters")
	}

	comment, err := c.ReviewCommandService.CreateReviewComment(&user, reviewCommentParam.ID, reviewCommentParam.Body)
	if err != nil {
		return errors.Wrap(err, "Failed to create review comment")
	}

	return ctx.JSON(http.StatusOK, c.ConvertReviewCommentToOutput(comment))
}

func (c *ReviewCommandController) DeleteReviewComment(ctx echo.Context, user entity.User) error {
	reviewCommentParam := &input.DeleteReviewCommentParam{}
	if err := BindAndValidate(ctx, reviewCommentParam); err != nil {
		return errors.Wrap(err, "Failed to bind parameters")
	}

	if err := c.ReviewCommandService.DeleteReviewComment(&user, reviewCommentParam.ID); err != nil {
		return errors.Wrap(err, "Failed to delete review comment")
	}

	return ctx.JSON(http.StatusOK, nil)
}

func (c *ReviewCommandController) StoreReviewCommentReply(ctx echo.Context, user entity.User) error {
	p := &input.CreateReviewCommentReply{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation store review comment reply input")
	}

	reply, err := c.ReviewCommandService.CreateReviewCommentReply(&user, c.ConvertCreateReviewCommentReplyParamToCommand(p))
	if err != nil {
		return errors.Wrap(err, "failed to store review comment reply")
	}

	return ctx.JSON(http.StatusOK, c.ConvertReviewCommentReplyToOutput(reply))
}

func (c *ReviewCommandController) DeleteReviewCommentReply(ctx echo.Context, user entity.User) error {
	i := &input.ReviewReplyParam{}
	if err := BindAndValidate(ctx, i); err != nil {
		return errors.Wrap(err, "invalid review reply param")
	}

	if err := c.ReviewCommandService.DeleteReviewCommentReply(&user, i.ReplyID); err != nil {
		return errors.Wrap(err, "failed delete review reply")
	}

	return ctx.JSON(http.StatusOK, "ok")
}

func (c *ReviewCommandController) FavoriteReviewComment(ctx echo.Context, user entity.User) error {
	p := &input.FavoriteReviewComment{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation favorite review comment input")
	}

	if err := c.ReviewCommandService.FavoriteReviewComment(&user, p.ReviewCommentID); err != nil {
		return errors.Wrap(err, "failed to favorite review_comment")
	}

	return ctx.JSON(http.StatusOK, "ok")
}

func (c *ReviewCommandController) Delete(ctx echo.Context, user entity.User) error {
	p := &input.DeleteReviewParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation delete review input")
	}

	if err := c.ReviewCommandScenario.DeleteReview(p.ID, &user); err != nil {
		return errors.Wrap(err, "failed to delete review")
	}

	return ctx.JSON(http.StatusOK, "ok")
}

func (c *ReviewCommandController) UnfavoriteReviewComment(ctx echo.Context, user entity.User) error {
	p := &input.FavoriteReviewComment{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation un favorite review comment input")
	}

	if err := c.ReviewCommandService.UnfavoriteReviewComment(&user, p.ReviewCommentID); err != nil {
		return errors.Wrap(err, "failed to un favorite review_comment")
	}

	return ctx.JSON(http.StatusOK, "ok")
}
