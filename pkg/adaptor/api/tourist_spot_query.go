package api

import (
	"net/http"

	"github.com/stayway-corp/stayway-media-api/pkg/application/scenario"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
)

type TouristSpotQueryController struct {
	converter.Converters
	scenario.TouristSpotQueryScenario
}

var TouristSpotQeuryControllerSet = wire.NewSet(
	wire.Struct(new(TouristSpotQueryController), "*"),
)

func (c *TouristSpotQueryController) Show(ctx echo.Context) error {
	p := &input.ShowTouristSpotParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation show post parameter")
	}

	touristSpot, areaCategoriesMap, themeCategoriesMap, err := c.TouristSpotQueryScenario.Show(p.ID)
	if err != nil {
		return errors.Wrap(err, "failed to get tourist_spot")
	}

	return ctx.JSON(http.StatusOK, c.ConvertTouristSpotToOutput(touristSpot, areaCategoriesMap, themeCategoriesMap))
}

func (c *TouristSpotQueryController) ListTouristSpot(ctx echo.Context) error {
	p := &input.ListTouristSpotParams{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation list tourist_spot parameter")
	}

	q := c.ConvertTouristSpotListParamToQuery(p)

	touristSpots, areaCategoriesMap, themeCategoriesMap, err := c.TouristSpotQueryScenario.List(q)
	if err != nil {
		return errors.Wrap(err, "failed to get tourist_spot list")
	}

	return ctx.JSON(http.StatusOK, c.ConvertTouristSpotListToOutput(touristSpots, areaCategoriesMap, themeCategoriesMap))
}

func (c *TouristSpotQueryController) ListRecommendTouristSpot(ctx echo.Context) error {
	p := &input.ListRecommendTouristSpotParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation recommend list tourist_spot")
	}

	q := c.ConvertRecommendTouristSpotListParamToQuery(p)

	touristSpots, areaCategoriesMap, themeCategoriesMap, err := c.TouristSpotQueryScenario.ListRecommend(q)
	if err != nil {
		return errors.Wrap(err, "failed to get recommend tourist_spot list")
	}

	return ctx.JSON(http.StatusOK, c.ConvertTouristSpotListToOutput(touristSpots, areaCategoriesMap, themeCategoriesMap))
}
