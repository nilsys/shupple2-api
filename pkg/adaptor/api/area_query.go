package api

import (
	"net/http"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
)

type (
	AreaQueryController struct {
		AreaQueryService service.AreaQueryService
	}
)

var AreaQueryControllerSet = wire.NewSet(
	wire.Struct(new(AreaQueryController), "*"),
)

// AreaのListを取得して返す
func (c *AreaQueryController) ListArea(ctx echo.Context) error {
	p := param.ListAreaParams{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrap(err, "failed to BindAndValidate")
	}

	categories, err := c.AreaQueryService.ListAreaByParams(p.AreaGroupID, p.PerPage, p.ExcludeID)
	if err != nil {
		return errors.Wrap(err, "failed to get categories")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertCategoriesToOutput(categories))
}

// IDに紐づくAreaを返す
func (c *AreaQueryController) ShowAreaByID(ctx echo.Context) error {
	p := param.GetArea{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrapf(err, "ID is invalid")
	}

	category, err := c.AreaQueryService.ShowAreaByID(p.ID)
	if err != nil {
		return errors.Wrap(err, "failed to category")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertCategoryToOutput(category))
}

// SubAreaのListを取得して返す
func (c *AreaQueryController) ListSubArea(ctx echo.Context) error {
	p := param.ListSubAreaParams{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrap(err, "failed to BindAndValidate CategoryParam")
	}

	categories, err := c.AreaQueryService.ListSubAreaByParams(p.AreaID, p.PerPage, p.ExcludeID)
	if err != nil {
		return errors.Wrap(err, "failed to get categories")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertCategoriesToOutput(categories))
}

// IDに紐づくSubAreaを返す
func (c *AreaQueryController) ShowSubAreaByID(ctx echo.Context) error {
	p := param.GetArea{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrapf(err, "ID is invalid")
	}

	category, err := c.AreaQueryService.ShowSubAreaByID(p.ID)
	if err != nil {
		return errors.Wrap(err, "failed to category")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertCategoryToOutput(category))
}

// SubSubAreaのListを取得して返す
func (c *AreaQueryController) ListSubSubArea(ctx echo.Context) error {
	p := param.ListSubSubAreaParams{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrap(err, "failed to BindAndValidate CategoryParam")
	}

	categories, err := c.AreaQueryService.ListSubSubAreaByParams(p.SubAreaID, p.PerPage, p.ExcludeID)
	if err != nil {
		return errors.Wrap(err, "failed to get categories")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertCategoriesToOutput(categories))
}

func (c *AreaQueryController) ShowSubSubAreaByID(ctx echo.Context) error {
	p := param.GetArea{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrapf(err, "ID is invalid")
	}

	category, err := c.AreaQueryService.ShowSubSubAreaByID(p.ID)
	if err != nil {
		return errors.Wrap(err, "failed to category")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertCategoryToOutput(category))
}
