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

type (
	ThemeQueryController struct {
		converter.Converters
		service.ThemeCategoryQueryService
	}
)

var ThemeQueryControllerSet = wire.NewSet(
	wire.Struct(new(ThemeQueryController), "*"),
)

func (c *ThemeQueryController) List(ctx echo.Context) error {
	q := input.ListThemeParams{}
	if err := BindAndValidate(ctx, &q); err != nil {
		return errors.Wrap(err, "failed to validate params")
	}

	categories, err := c.ThemeCategoryQueryService.ListThemeByParams(q.AreaID, q.SubAreaID, q.SubSubAreaID, q.ExcludeID)
	if err != nil {
		return errors.Wrap(err, "failed to get list of themes")
	}

	return ctx.JSON(http.StatusOK, c.ConvertThemeCategoriesWithPostCountToOutput(categories))
}

func (c *ThemeQueryController) ListSubThemeByParentID(ctx echo.Context) error {
	q := input.ListSubThemeParams{}
	if err := BindAndValidate(ctx, &q); err != nil {
		return errors.Wrap(err, "failed to validate params")
	}

	categories, err := c.ThemeCategoryQueryService.ListSubThemeByParams(q.ThemeID, q.AreaID, q.SubAreaID, q.SubSubAreaID, q.ExcludeID)
	if err != nil {
		return errors.Wrap(err, "failed to get subTheme list")
	}
	return ctx.JSON(http.StatusOK, c.ConvertThemeCategoriesWithPostCountToOutput(categories))
}
