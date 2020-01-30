package api

import (
	"net/http"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
)

type (
	PostQueryController struct {
		PostService service.PostQueryService
	}
)

var PostQueryControllerSet = wire.NewSet(
	wire.Struct(new(PostQueryController), "*"),
)

func (c *PostQueryController) FindByID(ctx echo.Context) error {
	q := param.GetPost{}
	if err := BindAndValidate(ctx, q); err != nil {
		return errors.Wrapf(err, "validation get post parameter")
	}

	post, err := c.PostService.FindByID(q.ID)
	if err != nil {
		return errors.Wrap(err, "failed to get post")
	}

	return ctx.JSON(http.StatusOK, post)
}

// TODO: 命名
func (c *PostQueryController) FindByParams(ctx echo.Context) error {
	params := param.FindPostListParam{}
	if err := BindAndValidate(ctx, params); err != nil {
		return errors.Wrapf(err, "validation find post list parameter")
	}

	query := converter.ConvertFindPostListParamToQuery(params)
	postAndCategories, err := c.PostService.FindByParams(query)
	if err != nil {
		return errors.Wrap(err, "failed to find post list")
	}

	return ctx.JSON(http.StatusOK, postAndCategories)
}
