package api

import (
	"net/http"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"

	"github.com/pkg/errors"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
)

type (
	CategoryQueryController struct {
		converter.Converters
		CategoryQueryService service.CategoryQueryService
	}
)

var CategoryQueryControllerSet = wire.NewSet(
	wire.Struct(new(CategoryQueryController), "*"),
)

func (c *CategoryQueryController) ShowBySlug(ctx echo.Context) error {
	p := input.ShowPostBySlug{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrap(err, "validation show by slug parameter")
	}

	category, err := c.CategoryQueryService.ShowBySlug(p.Slug)
	if err != nil {
		return errors.Wrap(err, "failed to show by slug")
	}

	return ctx.JSON(http.StatusOK, c.ConvertCategoryToOutput(category))
}
