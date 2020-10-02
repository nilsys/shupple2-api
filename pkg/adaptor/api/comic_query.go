package api

import (
	"net/http"

	"github.com/stayway-corp/stayway-media-api/pkg/application/scenario"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
)

type ComicQueryController struct {
	converter.Converters
	scenario.ComicQueryScenario
}

var ComicQueryControllerSet = wire.NewSet(
	wire.Struct(new(ComicQueryController), "*"),
)

func (c *ComicQueryController) Show(ctx echo.Context, ouser entity.OptionalUser) error {
	param := &input.ShowComicParam{}
	if err := BindAndValidate(ctx, param); err != nil {
		return errors.Wrapf(err, "validation show comic input")
	}

	comicDetail, idRelationFlgMap, err := c.ComicQueryScenario.Show(param.ID, &ouser)
	if err != nil {
		return errors.Wrap(err, "failed show comic")
	}

	return ctx.JSON(http.StatusOK, c.ConvertComicDetailToOutput(comicDetail, idRelationFlgMap))
}

func (c *ComicQueryController) ListComic(ctx echo.Context, ouser entity.OptionalUser) error {
	params := &input.ShowComicListParam{}
	if err := BindAndValidate(ctx, params); err != nil {
		return errors.Wrapf(err, "validation show comic list input")
	}

	comics, err := c.ComicQueryScenario.List(c.ConvertShowComicListParamToQuery(params), &ouser)
	if err != nil {
		return errors.Wrapf(err, "failed show comic list")
	}

	return ctx.JSON(http.StatusOK, c.ConvertComicListToOutput(comics))
}
