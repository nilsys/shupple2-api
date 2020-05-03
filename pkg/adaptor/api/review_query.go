package api

import (
	"net/http"

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
	ReviewQueryService service.ReviewQueryService
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

	r, err := c.ReviewQueryService.ShowReviewListByParams(reviewQuery, ouser)
	if err != nil {
		return errors.Wrap(err, "Failed to show review list")
	}

	return ctx.JSON(http.StatusOK, c.ConvertQueryReviewDetailWithIsFavoriteListToOutput(r))
}

func (c *ReviewQueryController) ListFeedReview(ctx echo.Context, ouser entity.OptionalUser) error {
	p := &input.ListFeedReviewParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation feed review input")
	}

	q := c.ConvertListFeedReviewParamToQuery(p)

	reviews, err := c.ReviewQueryService.ListFeed(ouser, p.UserID, q)
	if err != nil {
		return errors.Wrap(err, "failed to show feed review list")
	}

	return ctx.JSON(http.StatusOK, c.ConvertQueryReviewDetailWithIsFavoriteListToOutput(reviews))
}

func (c *ReviewQueryController) ShowReview(ctx echo.Context, ouser entity.OptionalUser) error {
	p := &input.ShowReview{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "required review id")
	}

	review, err := c.ReviewQueryService.ShowQueryReview(p.ID, ouser)
	if err != nil {
		return errors.Wrap(err, "failed show review")
	}

	return ctx.JSON(http.StatusOK, c.ConvertQueryReviewDetailWithIsFavoriteToOutput(review))
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

	reviews, err := c.ReviewQueryService.ListFavoriteReview(ouser, p.UserID, q)
	if err != nil {
		return errors.Wrap(err, "failed list favorite review")
	}

	return ctx.JSON(http.StatusOK, c.ConvertQueryReviewDetailWithIsFavoriteListToOutput(reviews))
}

func (c *ReviewQueryController) ListReviewCommentReply(ctx echo.Context) error {
	p := &input.ListReviewCommentReply{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation list review comment reply")
	}

	replies, err := c.ReviewQueryService.ListReviewCommentReplyByReviewCommentID(p.ReviewCommentID)
	if err != nil {
		return errors.Wrap(err, "failed list review comment reply")
	}

	return ctx.JSON(http.StatusOK, c.ConvertReviewCommentReplyListToOutput(replies))
}
