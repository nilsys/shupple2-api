package api

import (
	"net/http"

	"github.com/google/wire"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
)

type InterestQueryController struct {
	service.InterestQueryService
}

var InterestQueryControllerSet = wire.NewSet(
	wire.Struct(new(InterestQueryController), "*"),
)

func (c *InterestQueryController) ListAll(ctx echo.Context) error {
	interests, err := c.InterestQueryService.ListAll()
	if err != nil {
		return errors.Wrap(err, "failed list all interest")
	}
	return ctx.JSON(http.StatusOK, interests)
}
