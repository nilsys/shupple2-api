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
	AreaQueryController struct {
		converter.Converters
		AreaQueryService service.AreaQueryService
	}
)

var AreaQueryControllerSet = wire.NewSet(
	wire.Struct(new(AreaQueryController), "*"),
)

// AreaのListを取得して返す
func (c *AreaQueryController) ListArea(ctx echo.Context) error {
	p := input.ListAreaParams{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrap(err, "failed to BindAndValidate")
	}

	areaCategories, err := c.AreaQueryService.ListAreaByParams(p.AreaGroup, p.PerPage, p.ExcludeID)
	if err != nil {
		return errors.Wrap(err, "failed to get areaCategories")
	}

	return ctx.JSON(http.StatusOK, c.ConvertAreaCategoriesWithPostCountToOutput(areaCategories))
}

// IDに紐づくAreaを返す
func (c *AreaQueryController) ShowAreaByID(ctx echo.Context) error {
	p := input.GetArea{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrapf(err, "ID is invalid")
	}

	areaCategoryDetail, err := c.AreaQueryService.ShowAreaByID(p.ID)
	if err != nil {
		return errors.Wrap(err, "failed to areaCategory")
	}

	return ctx.JSON(http.StatusOK, c.ConvertAreaCategoryDetailToOutput(areaCategoryDetail))
}

// SubAreaのListを取得して返す
func (c *AreaQueryController) ListSubArea(ctx echo.Context) error {
	p := input.ListSubAreaParams{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrap(err, "failed to BindAndValidate AreaCategoryParam")
	}

	areaCategories, err := c.AreaQueryService.ListSubAreaByParams(p.AreaID, p.ThemeID, p.PerPage, p.ExcludeID)
	if err != nil {
		return errors.Wrap(err, "failed to get areaCategories")
	}

	return ctx.JSON(http.StatusOK, c.ConvertAreaCategoriesWithPostCountToOutput(areaCategories))
}

// IDに紐づくSubAreaを返す
func (c *AreaQueryController) ShowSubAreaByID(ctx echo.Context) error {
	p := input.GetArea{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrapf(err, "ID is invalid")
	}

	areaCategory, err := c.AreaQueryService.ShowSubAreaByID(p.ID)
	if err != nil {
		return errors.Wrap(err, "failed to areaCategory")
	}

	return ctx.JSON(http.StatusOK, c.ConvertAreaCategoryDetailToOutput(areaCategory))
}

// SubSubAreaのListを取得して返す
func (c *AreaQueryController) ListSubSubArea(ctx echo.Context) error {
	p := input.ListSubSubAreaParams{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrap(err, "failed to BindAndValidate AreaCategoryParam")
	}

	areaCategories, err := c.AreaQueryService.ListSubSubAreaByParams(p.SubAreaID, p.ThemeID, p.PerPage, p.ExcludeID)
	if err != nil {
		return errors.Wrap(err, "failed to get areaCategories")
	}

	return ctx.JSON(http.StatusOK, c.ConvertAreaCategoriesWithPostCountToOutput(areaCategories))
}

func (c *AreaQueryController) ShowSubSubAreaByID(ctx echo.Context) error {
	p := input.GetArea{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrapf(err, "ID is invalid")
	}

	areaCategory, err := c.AreaQueryService.ShowSubSubAreaByID(p.ID)
	if err != nil {
		return errors.Wrap(err, "failed to areaCategory")
	}

	return ctx.JSON(http.StatusOK, c.ConvertAreaCategoryDetailToOutput(areaCategory))
}
