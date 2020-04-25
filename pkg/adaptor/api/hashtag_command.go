package api

import (
	"net/http"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type HashtagCommandController struct {
	converter.Converters
	service.HashtagCommandService
}

var HashtagCommandControllerSet = wire.NewSet(
	wire.Struct(new(HashtagCommandController), "*"),
)

func (c *HashtagCommandController) FollowHashtag(ctx echo.Context, user entity.User) error {
	p := input.FollowParam{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrap(err, "validation follow hashtag input")
	}

	if err := c.HashtagCommandService.FollowHashtag(&user, p.ID); err != nil {
		return errors.Wrap(err, "failed to follow hashtag")
	}

	return ctx.JSON(http.StatusOK, "ok")
}

func (c *HashtagCommandController) UnfollowHashtag(ctx echo.Context, user entity.User) error {
	p := input.FollowParam{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrap(err, "validation unfollow hashtag input")
	}

	if err := c.HashtagCommandService.UnfollowHashtag(&user, p.ID); err != nil {
		return errors.Wrap(err, "failed to un follow hashtag")
	}

	return ctx.JSON(http.StatusOK, "ok")
}
