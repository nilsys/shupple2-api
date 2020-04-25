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

type WordpressCallbackController struct {
	converter.Converters
	service.WordpressCallbackService
}

var WordpressCallbackControllerSet = wire.NewSet(
	wire.Struct(new(WordpressCallbackController), "*"),
)

func (c *WordpressCallbackController) Import(ctx echo.Context) error {
	param := &input.ImportWordpressEntityParam{}
	if err := BindAndValidate(ctx, param); err != nil {
		return errors.Wrap(err, "invalid import wordpress entity input")
	}

	if err := c.WordpressCallbackService.Import(param.EntityType, param.ID); err != nil {
		return errors.Wrap(err, "failed to import wordpress entity")
	}

	return ctx.NoContent(http.StatusNoContent)
}
