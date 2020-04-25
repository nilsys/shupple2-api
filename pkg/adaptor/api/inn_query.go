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

type InnQueryController struct {
	converter.Converters
	service.InnQueryService
}

var InnQueryControllerSet = wire.NewSet(
	wire.Struct(new(InnQueryController), "*"),
)

func (c *InnQueryController) ListByParams(ctx echo.Context) error {
	p := &input.ListInn{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation list inn params")
	}

	inn, err := c.InnQueryService.ListInnByParams(p.AreaID, p.SubAreaID, p.SubSubAreaID, p.TouristSpotID)
	if err != nil {
		return errors.Wrap(err, "failed list inn")
	}

	return ctx.JSON(http.StatusOK, inn)
}
