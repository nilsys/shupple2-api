package api

import (
	"net/http"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"

	"github.com/google/wire"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
)

type ReviewQueryController struct {
	ReviewQueryService service.ReviewQueryService
}

var ReviewQueryControllerSet = wire.NewSet(
	wire.Struct(new(ReviewQueryController), "*"),
)

func (c *ReviewQueryController) LisReview(ctx echo.Context) error {
	reviewParam := &param.ListReviewParams{}

	if err := BindAndValidate(ctx, reviewParam); err != nil {
		return errors.Wrap(err, "Failed to bind parameters")
	}

	reviewQuery := converter.ConvertFindReviewListParamToQuery(reviewParam)

	r, err := c.ReviewQueryService.ShowReviewListByParams(reviewQuery)
	if err != nil {
		return errors.Wrap(err, "Failed to show review list")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertQueryReviewListToOutput(r))
}

func (c *ReviewQueryController) ListFeedReview(ctx echo.Context) error {
	p := &param.ListFeedReviewParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation feed review param")
	}

	q := converter.ConvertListFeedReviewParamToQuery(p)

	reviews, err := c.ReviewQueryService.ShowListFeed(p.UserID, q)
	if err != nil {
		return errors.Wrap(err, "failed to show feed review list")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertQueryReviewListToOutput(reviews))
}

func (c *ReviewQueryController) ShowReview(ctx echo.Context) error {
	p := &param.ShowReview{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "required review id")
	}

	review, err := c.ReviewQueryService.ShowQueryReview(p.ID)
	if err != nil {
		return errors.Wrap(err, "failed show review")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertQueryReviewShowToOutput(review))
}

func (c *ReviewQueryController) ListReviewCommentByReviewID(ctx echo.Context) error {
	p := &param.ListReviewCommentParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "Failed to bind parameters")
	}

	reviewComments, err := c.ReviewQueryService.ListReviewCommentByReviewID(p.ID, p.GetLimit())
	if err != nil {
		return errors.Wrap(err, "failed to show review comment list")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertReviewCommentListToOutput(reviewComments))
}

func (c *ReviewQueryController) ListFavoriteReview(ctx echo.Context) error {
	p := &param.ListFeedReviewParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation list favorite review")
	}

	q := converter.ConvertListFeedReviewParamToQuery(p)

	reviews, err := c.ReviewQueryService.ListFavoriteReview(p.UserID, q)
	if err != nil {
		return errors.Wrap(err, "failed list favorite review")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertQueryReviewListToOutput(reviews))
}

func (c *ReviewQueryController) ListReviewCommentReply(ctx echo.Context) error {
	p := &param.ListReviewCommentReply{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation list review comment reply")
	}

	replies, err := c.ReviewQueryService.ListReviewCommentReplyByReviewCommentID(p.ReviewCommentID)
	if err != nil {
		return errors.Wrap(err, "failed list review comment reply")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertReviewCommentReplyListToOutput(replies))
}
