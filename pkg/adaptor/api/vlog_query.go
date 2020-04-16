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

type VlogQueryController struct {
	service.VlogQueryService
}

var VlogQueryControllerSet = wire.NewSet(
	wire.Struct(new(VlogQueryController), "*"),
)

func (c *VlogQueryController) Show(ctx echo.Context) error {
	p := &input.ShowVlog{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "required vlog id")
	}

	vlog, err := c.VlogQueryService.Show(p.ID)
	if err != nil {
		return errors.Wrapf(err, "failed show vlog id=%d", p.ID)
	}

	return ctx.JSON(http.StatusOK, converter.ConvertVlogDetail(vlog))
}
func (c *VlogQueryController) ListVlog(ctx echo.Context) error {
	param := &input.ListVlogParam{}
	if err := BindAndValidate(ctx, param); err != nil {
		return errors.Wrap(err, "invalid show vlogs input")
	}

	vlogs, err := c.VlogQueryService.ShowListByParams(converter.ConvertListVlogParamToQuery(param))
	if err != nil {
		return errors.Wrap(err, "failed show vlog list")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertVlogListToOutput(vlogs))
}
