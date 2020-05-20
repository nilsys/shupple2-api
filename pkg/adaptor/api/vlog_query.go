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

type VlogQueryController struct {
	converter.Converters
	scenario.VlogQueryScenario
}

var VlogQueryControllerSet = wire.NewSet(
	wire.Struct(new(VlogQueryController), "*"),
)

func (c *VlogQueryController) Show(ctx echo.Context) error {
	p := &input.ShowVlog{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "required vlog id")
	}

	vlog, areaCategoriesMap, themeCategoriesMap, err := c.VlogQueryScenario.Show(p.ID)
	if err != nil {
		return errors.Wrapf(err, "failed show vlog id=%d", p.ID)
	}

	return ctx.JSON(http.StatusOK, c.ConvertVlogDetail(vlog, areaCategoriesMap, themeCategoriesMap))
}
func (c *VlogQueryController) ListVlog(ctx echo.Context) error {
	param := &input.ListVlogParam{}
	if err := BindAndValidate(ctx, param); err != nil {
		return errors.Wrap(err, "invalid show vlogs input")
	}

	vlogs, areaCategoriesMap, themeCategoriesMap, err := c.VlogQueryScenario.ListByParams(c.ConvertListVlogParamToQuery(param))
	if err != nil {
		return errors.Wrap(err, "failed show vlog list")
	}

	return ctx.JSON(http.StatusOK, c.ConvertVlogListToOutput(vlogs, areaCategoriesMap, themeCategoriesMap))
}
