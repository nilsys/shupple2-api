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

type VlogQueryController struct {
	converter.Converters
	scenario.VlogQueryScenario
}

var VlogQueryControllerSet = wire.NewSet(
	wire.Struct(new(VlogQueryController), "*"),
)

// TODO: review_countをtourist_spotテーブルに追加して、Review投稿時にIncrementする様にする、その際にscriptを書いて既存のReviewの数を含める
func (c *VlogQueryController) Show(ctx echo.Context, ouser entity.OptionalUser) error {
	p := &input.ShowVlog{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "required vlog id")
	}

	vlog, touristSpotReviewCountMap, areaCategoriesMap, themeCategoriesMap, idRelationFlgMap, err := c.VlogQueryScenario.Show(p.ID, &ouser)
	if err != nil {
		return errors.Wrapf(err, "failed show vlog id=%d", p.ID)
	}

	return ctx.JSON(http.StatusOK, c.ConvertVlogDetail(vlog, touristSpotReviewCountMap, areaCategoriesMap, themeCategoriesMap, idRelationFlgMap))
}
func (c *VlogQueryController) ListVlog(ctx echo.Context, ouser entity.OptionalUser) error {
	param := &input.ListVlogParam{}
	if err := BindAndValidate(ctx, param); err != nil {
		return errors.Wrap(err, "invalid show vlogs input")
	}

	vlogs, areaCategoriesMap, themeCategoriesMap, err := c.VlogQueryScenario.List(c.ConvertListVlogParamToQuery(param), &ouser)
	if err != nil {
		return errors.Wrap(err, "failed show vlog list")
	}

	return ctx.JSON(http.StatusOK, c.ConvertVlogListToOutput(vlogs, areaCategoriesMap, themeCategoriesMap))
}
