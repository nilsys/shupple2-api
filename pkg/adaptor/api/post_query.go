package api

import (
	"net/http"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
)

type (
	PostQueryController struct {
		converter.Converters
		PostService service.PostQueryService
	}
)

var PostQueryControllerSet = wire.NewSet(
	wire.Struct(new(PostQueryController), "*"),
)

func (c *PostQueryController) Show(ctx echo.Context, ouser entity.OptionalUser) error {
	p := &input.GetPost{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrapf(err, "validation get post parameter")
	}

	post, err := c.PostService.ShowQueryByID(p.ID, ouser)
	if err != nil {
		return errors.Wrap(err, "failed to get post")
	}

	return ctx.JSON(http.StatusOK, c.ConvertPostDetailWithHashtagAndIsFavoriteToOutput(post))
}

func (c *PostQueryController) ShowBySlug(ctx echo.Context) error {
	p := &input.ShowPostBySlug{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation get post by slug parameter")
	}

	post, err := c.PostService.ShowQueryBySlug(p.Slug)
	if err != nil {
		return errors.Wrap(err, "failed to get post by slug")
	}

	return ctx.JSON(http.StatusOK, c.ConvertPostDetailWithHashtagToOutput(post))
}

func (c *PostQueryController) ListPost(ctx echo.Context, ouser entity.OptionalUser) error {
	p := &input.ListPostParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrapf(err, "validation find post list parameter")
	}

	query := c.ConvertFindPostListParamToQuery(p)

	posts, err := c.PostService.ListByParams(query, ouser)
	if err != nil {
		return errors.Wrap(err, "failed to find post list")
	}

	return ctx.JSON(http.StatusOK, c.ConvertPostListToOutput(posts))
}

func (c *PostQueryController) ListFeedPost(ctx echo.Context, ouser entity.OptionalUser) error {
	p := &input.ListFeedPostParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrapf(err, "validation find feed post list input")
	}

	q := c.ConvertListFeedPostParamToQuery(p)

	posts, err := c.PostService.ListFeed(ouser, p.UserID, q)
	if err != nil {
		return errors.Wrap(err, "failed to show feed posts")
	}

	return ctx.JSON(http.StatusOK, c.ConvertPostListToOutput(posts))
}

func (c *PostQueryController) ListFavoritePost(ctx echo.Context, ouser entity.OptionalUser) error {
	p := &input.ListFeedPostParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation list favorite post")
	}

	q := c.ConvertListFeedPostParamToQuery(p)

	posts, err := c.PostService.ListFavoritePost(ouser, p.UserID, q)
	if err != nil {
		return errors.Wrap(err, "failed list favorite post")
	}

	return ctx.JSON(http.StatusOK, c.ConvertPostListToOutput(posts))
}
