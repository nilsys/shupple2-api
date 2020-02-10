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

	return ctx.JSON(http.StatusOK, r)
}
