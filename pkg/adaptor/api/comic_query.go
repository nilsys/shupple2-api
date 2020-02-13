package api

import (
	"net/http"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
)

type ComicQueryController struct {
	service.ComicQueryService
}

var ComicQueryControllerSet = wire.NewSet(
	wire.Struct(new(ComicQueryController), "*"),
)

func (c *ComicQueryController) Show(ctx echo.Context) error {
	param := &param.ShowComicParam{}
	if err := BindAndValidate(ctx, param); err != nil {
		return errors.Wrapf(err, "validation show comic param")
	}

	comicDetail, err := c.ComicQueryService.Show(param.ID)
	if err != nil {
		return errors.Wrap(err, "failed show comic")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertQueryComicOutput(comicDetail))
}

func (c *ComicQueryController) ListComic(ctx echo.Context) error {
	params := &param.ShowComicListParam{}
	if err := BindAndValidate(ctx, params); err != nil {
		return errors.Wrapf(err, "validation show comic list param")
	}

	comics, err := c.ComicQueryService.ShowList(converter.ConvertShowComicListParamToQuery(params))
	if err != nil {
		return errors.Wrapf(err, "failed show comic list")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertComicListToOutput(comics))
}
