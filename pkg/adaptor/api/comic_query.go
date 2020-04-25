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

type ComicQueryController struct {
	converter.Converters
	service.ComicQueryService
}

var ComicQueryControllerSet = wire.NewSet(
	wire.Struct(new(ComicQueryController), "*"),
)

func (c *ComicQueryController) Show(ctx echo.Context) error {
	param := &input.ShowComicParam{}
	if err := BindAndValidate(ctx, param); err != nil {
		return errors.Wrapf(err, "validation show comic input")
	}

	comicDetail, err := c.ComicQueryService.Show(param.ID)
	if err != nil {
		return errors.Wrap(err, "failed show comic")
	}

	return ctx.JSON(http.StatusOK, c.ConvertQueryComicOutput(comicDetail))
}

func (c *ComicQueryController) ListComic(ctx echo.Context) error {
	params := &input.ShowComicListParam{}
	if err := BindAndValidate(ctx, params); err != nil {
		return errors.Wrapf(err, "validation show comic list input")
	}

	comics, err := c.ComicQueryService.ShowList(c.ConvertShowComicListParamToQuery(params))
	if err != nil {
		return errors.Wrapf(err, "failed show comic list")
	}

	return ctx.JSON(http.StatusOK, c.ConvertComicListToOutput(comics))
}
