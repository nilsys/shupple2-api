package api

import (
	"net/http"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

	"github.com/stayway-corp/stayway-media-api/pkg/application/scenario"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
)

type FeatureQueryController struct {
	converter.Converters
	scenario.FeatureQueryScenario
}

var FeatureQueryControllerSet = wire.NewSet(
	wire.Struct(new(FeatureQueryController), "*"),
)

func (c *FeatureQueryController) Show(ctx echo.Context, ouser entity.OptionalUser) error {
	p := &input.ShowFeatureParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation show feature")
	}

	queryFeature, areaCategories, themeCategories, idIsFollowMap, err := c.FeatureQueryScenario.Show(p.ID, &ouser)
	if err != nil {
		return errors.Wrap(err, "failed show query feature")
	}

	return ctx.JSON(http.StatusOK, c.ConvertFeatureDetailPostsToOutput(queryFeature, areaCategories, themeCategories, idIsFollowMap))
}

func (c *FeatureQueryController) List(ctx echo.Context) error {
	p := &input.ShowFeatureListParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation show feature list input")
	}

	q := c.ConvertShowFeatureListParamToQuery(p)

	features, err := c.FeatureQueryScenario.List(q)
	if err != nil {
		return errors.Wrap(err, "failed show feature list")
	}

	return ctx.JSON(http.StatusOK, c.ConvertFeatureListToOutput(features))
}
