package api

import (
	"net/http"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type HashtagCommandController struct {
	service.HashtagCommandService
}

var HashtagCommandControllerSet = wire.NewSet(
	wire.Struct(new(HashtagCommandController), "*"),
)

func (c *HashtagCommandController) FollowHashtag(ctx echo.Context, user entity.User) error {
	p := param.FollowParam{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrap(err, "validation follow hashtag param")
	}

	if err := c.HashtagCommandService.FollowHashtag(&user, p.ID); err != nil {
		return errors.Wrap(err, "failed to follow hashtag")
	}

	return ctx.JSON(http.StatusOK, "ok")
}

func (c *HashtagCommandController) UnfollowHashtag(ctx echo.Context, user entity.User) error {
	p := param.FollowParam{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrap(err, "validation unfollow hashtag param")
	}

	if err := c.HashtagCommandService.UnfollowHashtag(&user, p.ID); err != nil {
		return errors.Wrap(err, "failed to un follow hashtag")
	}

	return ctx.JSON(http.StatusOK, "ok")
}
