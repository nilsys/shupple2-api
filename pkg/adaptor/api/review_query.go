package api

import (
	"net/http"

	"github.com/stayway-corp/stayway-media-api/pkg/application/scenario"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"

	"github.com/google/wire"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
)

type ReviewQueryController struct {
	converter.Converters
	service.ReviewQueryService
	scenario.ReviewQueryScenario
}

var ReviewQueryControllerSet = wire.NewSet(
	wire.Struct(new(ReviewQueryController), "*"),
)

func (c *ReviewQueryController) LisReview(ctx echo.Context, ouser entity.OptionalUser) error {
	reviewParam := &input.ListReviewParams{}

	if err := BindAndValidate(ctx, reviewParam); err != nil {
		return errors.Wrap(err, "Failed to bind parameters")
	}

	reviewQuery := c.ConvertFindReviewListParamToQuery(reviewParam)

	list, idIsFollowMap, err := c.ReviewQueryScenario.ListByParams(reviewQuery, ouser)
	if err != nil {
		return errors.Wrap(err, "Failed to show review list")
	}

	return ctx.JSON(http.StatusOK, c.ConvertQueryReviewDetailWithIsFavoriteListToOutput(list, idIsFollowMap))
}

func (c *ReviewQueryController) ListFeedReview(ctx echo.Context, ouser entity.OptionalUser) error {
	p := &input.ListFeedReviewParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation feed review input")
	}

	q := c.ConvertListFeedReviewParamToQuery(p)

	reviews, idIsFollowMap, err := c.ReviewQueryScenario.ListFeed(p.UserID, q, ouser)
	if err != nil {
		return errors.Wrap(err, "failed to show feed review list")
	}

	return ctx.JSON(http.StatusOK, c.ConvertQueryReviewDetailWithIsFavoriteListToOutput(reviews, idIsFollowMap))
}

func (c *ReviewQueryController) ShowReview(ctx echo.Context, ouser entity.OptionalUser) error {
	p := &input.ShowReview{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "required review id")
	}

	review, idIsFollowMap, err := c.ReviewQueryScenario.Show(p.ID, ouser)
	if err != nil {
		return errors.Wrap(err, "failed show review")
	}

	return ctx.JSON(http.StatusOK, c.ConvertQueryReviewDetailWithIsFavoriteToOutput(review, idIsFollowMap))
}

func (c *ReviewQueryController) ListReviewCommentByReviewID(ctx echo.Context, ouser entity.OptionalUser) error {
	p := &input.ListReviewCommentParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "Failed to bind parameters")
	}

	reviewComments, err := c.ReviewQueryService.ListReviewCommentByReviewID(p.ID, p.GetLimit(), ouser)
	if err != nil {
		return errors.Wrap(err, "failed to show review comment list")
	}

	return ctx.JSON(http.StatusOK, c.ConvertReviewCommentWithIsFavoriteListToOutput(reviewComments))
}

func (c *ReviewQueryController) ListFavoriteReview(ctx echo.Context, ouser entity.OptionalUser) error {
	p := &input.ListFeedReviewParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation list favorite review")
	}

	q := c.ConvertListFeedReviewParamToQuery(p)

	reviews, idIsFollowMap, err := c.ReviewQueryScenario.ListFavorite(p.UserID, q, ouser)
	if err != nil {
		return errors.Wrap(err, "failed list favorite review")
	}

	return ctx.JSON(http.StatusOK, c.ConvertQueryReviewDetailWithIsFavoriteListToOutput(reviews, idIsFollowMap))
}

func (c *ReviewQueryController) ListReviewCommentReply(ctx echo.Context, ouser entity.OptionalUser) error {
	p := &input.ListReviewCommentReply{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation list review comment reply")
	}

	replies, err := c.ReviewQueryService.ListReviewCommentReplyByReviewCommentID(p.ReviewCommentID, &ouser)
	if err != nil {
		return errors.Wrap(err, "failed list review comment reply")
	}

	return ctx.JSON(http.StatusOK, c.ConvertReviewCommentReplyListToOutput(replies))
}
