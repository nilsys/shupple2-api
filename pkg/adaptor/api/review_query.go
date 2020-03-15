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

	reviews, err := c.ReviewQueryService.ShowListFeed(p.ID, q)
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

	review, err := c.ReviewQueryService.ShowReview(p.ID)
	if err != nil {
		return errors.Wrap(err, "failed show review")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertQueryReviewShowToOutput(review))
}
