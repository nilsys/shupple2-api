package api

import (
	"net/http"

	"github.com/stayway-corp/stayway-media-api/pkg/application/scenario"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
)

type FeatureQueryController struct {
	converter.Converters
	// 処理が複雑な物はシナリオクラスを挟み、serviceにアクセスしている
	FeatureQueryScenario scenario.FeatureQueryScenario
	FeatureQueryService  service.FeatureQueryService
}

var FeatureQueryControllerSet = wire.NewSet(
	wire.Struct(new(FeatureQueryController), "*"),
)

func (c *FeatureQueryController) ShowQuery(ctx echo.Context) error {
	p := &input.ShowFeatureParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation show feature")
	}

	queryFeature, areaCategories, themeCategories, err := c.FeatureQueryScenario.Show(p.ID)
	if err != nil {
		return errors.Wrap(err, "failed show query feature")
	}

	return ctx.JSON(http.StatusOK, c.ConvertFeatureDetailPostsToOutput(queryFeature, areaCategories, themeCategories))
}

func (c *FeatureQueryController) ListFeature(ctx echo.Context) error {
	p := &input.ShowFeatureListParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation show feature list input")
	}

	q := c.ConvertShowFeatureListParamToQuery(p)

	features, err := c.FeatureQueryService.List(q)
	if err != nil {
		return errors.Wrap(err, "failed show feature list")
	}

	return ctx.JSON(http.StatusOK, c.ConvertFeatureListToOutput(features))
}
