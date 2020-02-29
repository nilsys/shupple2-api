package api

import (
	"net/http"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"

	"github.com/pkg/errors"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
)

type (
	CategoryQueryController struct {
		CategoryQueryService service.CategoryQueryService
	}
)

var CategoryQueryControllerSet = wire.NewSet(
	wire.Struct(new(CategoryQueryController), "*"),
)

// AreaのListを取得して返す
func (c *CategoryQueryController) ListArea(ctx echo.Context) error {
	q := param.GetArea{}
	if err := BindAndValidate(ctx, &q); err != nil {
		return errors.Wrap(err, "failed to BindAndValidate")
	}

	categories, err := c.CategoryQueryService.ShowAreaListByParams(q.ParentCategoryID, q.PerPage, q.ExcludeID)
	if err != nil {
		return errors.Wrap(err, "failed to get categories")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertCategoriesToOutput(categories))
}

// IDに紐づくAreaを返す
func (c *CategoryQueryController) ShowAreaByID(ctx echo.Context) error {
	q := param.AreaID{}
	if err := BindAndValidate(ctx, &q); err != nil {
		return errors.Wrapf(err, "ID is invalid")
	}

	category, err := c.CategoryQueryService.ShowAreaByID(q.ID)
	if err != nil {
		return errors.Wrap(err, "failed to ShowAreaByID")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertCategoryToOutput(category))
}
