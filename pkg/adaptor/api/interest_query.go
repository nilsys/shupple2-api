package api

import (
	"net/http"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"

	"github.com/google/wire"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
)

type InterestQueryController struct {
	converter.Converters
	service.InterestQueryService
}

var InterestQueryControllerSet = wire.NewSet(
	wire.Struct(new(InterestQueryController), "*"),
)

func (c *InterestQueryController) ListAll(ctx echo.Context) error {
	i := &input.ListInterest{}
	if err := BindAndValidate(ctx, i); err != nil {
		return errors.Wrap(err, "validation list inn params")
	}
	interests, err := c.InterestQueryService.ListAll(i.InterestGroup)
	if err != nil {
		return errors.Wrap(err, "failed list all interest")
	}
	return ctx.JSON(http.StatusOK, interests)
}
