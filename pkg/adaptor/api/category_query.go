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
	q := param.ListAreaParams{}
	if err := BindAndValidate(ctx, &q); err != nil {
		return errors.Wrap(err, "failed to BindAndValidate")
	}

	categories, err := c.CategoryQueryService.ListAreaByParams(q.AreaGroupID, q.PerPage, q.ExcludeID)
	if err != nil {
		return errors.Wrap(err, "failed to get categories")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertCategoriesToOutput(categories))
}

func (c *CategoryQueryController) ShowBySlug(ctx echo.Context) error {
	p := param.ShowPostBySlug{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrap(err, "validation show by slug parameter")
	}

	category, err := c.CategoryQueryService.ShowBySlug(p.Slug)
	if err != nil {
		return errors.Wrap(err, "failed to show by slug")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertCategoryToOutput(category))
}

// IDに紐づくAreaを返す
func (c *CategoryQueryController) ShowAreaByID(ctx echo.Context) error {
	q := param.GetArea{}
	if err := BindAndValidate(ctx, &q); err != nil {
		return errors.Wrapf(err, "ID is invalid")
	}

	category, err := c.CategoryQueryService.ShowAreaByID(q.ID)
	if err != nil {
		return errors.Wrap(err, "failed to category")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertCategoryToOutput(category))
}

// SubAreaのListを取得して返す
func (c *CategoryQueryController) ListSubArea(ctx echo.Context) error {
	q := param.ListSubAreaParams{}
	if err := BindAndValidate(ctx, &q); err != nil {
		return errors.Wrap(err, "failed to BindAndValidate CategoryParam")
	}

	categories, err := c.CategoryQueryService.ListSubAreaByParams(q.AreaID, q.PerPage, q.ExcludeID)
	if err != nil {
		return errors.Wrap(err, "failed to get categories")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertCategoriesToOutput(categories))
}

// IDに紐づくSubAreaを返す
func (c *CategoryQueryController) ShowSubAreaByID(ctx echo.Context) error {
	q := param.GetArea{}
	if err := BindAndValidate(ctx, &q); err != nil {
		return errors.Wrapf(err, "ID is invalid")
	}

	category, err := c.CategoryQueryService.ShowSubAreaByID(q.ID)
	if err != nil {
		return errors.Wrap(err, "failed to category")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertCategoryToOutput(category))
}

// SubSubAreaのListを取得して返す
func (c *CategoryQueryController) ListSubSubArea(ctx echo.Context) error {
	q := param.ListSubSubAreaParams{}
	if err := BindAndValidate(ctx, &q); err != nil {
		return errors.Wrap(err, "failed to BindAndValidate CategoryParam")
	}

	categories, err := c.CategoryQueryService.ListSubSubAreaByParams(q.SubAreaID, q.PerPage, q.ExcludeID)
	if err != nil {
		return errors.Wrap(err, "failed to get categories")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertCategoriesToOutput(categories))
}

func (c *CategoryQueryController) ShowSubSubAreaByID(ctx echo.Context) error {
	q := param.GetArea{}
	if err := BindAndValidate(ctx, &q); err != nil {
		return errors.Wrapf(err, "ID is invalid")
	}

	category, err := c.CategoryQueryService.ShowSubSubAreaByID(q.ID)
	if err != nil {
		return errors.Wrap(err, "failed to category")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertCategoryToOutput(category))
}
