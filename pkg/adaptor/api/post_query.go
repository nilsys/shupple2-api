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

func (c *PostQueryController) Show(ctx echo.Context) error {
	q := &param.GetPost{}
	if err := BindAndValidate(ctx, q); err != nil {
		return errors.Wrapf(err, "validation get post parameter")
	}

	post, err := c.PostService.ShowQueryByID(q.ID)
	if err != nil {
		return errors.Wrap(err, "failed to get post")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertQueryShowPostToOutput(post))
}

func (c *PostQueryController) ListPost(ctx echo.Context) error {
	p := &param.ListPostParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrapf(err, "validation find post list parameter")
	}

	query := converter.ConvertFindPostListParamToQuery(p)

	posts, err := c.PostService.ShowListByParams(query)
	if err != nil {
		return errors.Wrap(err, "failed to find post list")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertPostToOutput(posts))
}

func (c *PostQueryController) ListFeedPost(ctx echo.Context) error {
	p := &param.ListFeedPostParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrapf(err, "validation find feed post list param")
	}

	q := converter.ConvertListFeedPostParamToQuery(p)

	posts, err := c.PostService.ShowListFeed(p.UserID, q)
	if err != nil {
		return errors.Wrap(err, "failed to show feed posts")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertPostToOutput(posts))
}
