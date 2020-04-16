package api

import (
	"net/http"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
)

type TouristSpotQueryController struct {
	service.TouristSpotQueryService
}

var TouristSpotQeuryControllerSet = wire.NewSet(
	wire.Struct(new(TouristSpotQueryController), "*"),
)

func (c *TouristSpotQueryController) Show(ctx echo.Context) error {
	p := &input.ShowTouristSpotParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation show post parameter")
	}

	touristSpot, err := c.TouristSpotQueryService.Show(p.ID)
	if err != nil {
		return errors.Wrap(err, "failed to get tourist_spot")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertQueryTouristSpotToOutput(touristSpot))
}

func (c *TouristSpotQueryController) ListTouristSpot(ctx echo.Context) error {
	p := &input.ListTouristSpotParams{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation list tourist_spot parameter")
	}

	q := converter.ConvertTouristSpotListParamToQuery(p)

	touristSpots, err := c.TouristSpotQueryService.List(q)
	if err != nil {
		return errors.Wrap(err, "failed to get tourist_spot list")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertTouristSpotToOutput(touristSpots))
}

func (c *TouristSpotQueryController) ListRecommendTouristSpot(ctx echo.Context) error {
	p := &input.ListRecommendTouristSpotParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation recommend list tourist_spot")
	}

	q := converter.ConvertRecommendTouristSpotListParamToQuery(p)

	touristSpots, err := c.TouristSpotQueryService.ListRecommend(q)
	if err != nil {
		return errors.Wrap(err, "failed to get recommend tourist_spot list")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertTouristSpotToOutput(touristSpots))
}
