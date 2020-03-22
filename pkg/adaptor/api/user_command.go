package api

import (
	"net/http"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

	"github.com/google/wire"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
)

type UserCommandController struct {
	service.UserCommandService
}

var UserCommandControllerSet = wire.NewSet(
	wire.Struct(new(UserCommandController), "*"),
)

func (c *UserCommandController) SignUp(ctx echo.Context) error {
	p := param.StoreUser{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrap(err, "validation store user param")
	}

	err := c.UserCommandService.SignUp(converter.ConvertStoreUserParamToEntity(&p), p.CognitoToken, p.MigrationCode)
	if err != nil {
		return errors.Wrap(err, "failed to store user")
	}

	return ctx.JSON(http.StatusOK, "ok")
}

func (c *UserCommandController) Follow(ctx echo.Context, user entity.User) error {
	p := param.FollowParam{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrap(err, "validation follow user param")
	}

	if err := c.UserCommandService.Follow(&user, p.ID); err != nil {
		return errors.Wrap(err, "failed to store follow user")
	}

	return ctx.JSON(http.StatusOK, "ok")
}

func (c *UserCommandController) Unfollow(ctx echo.Context, user entity.User) error {
	p := param.FollowParam{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrap(err, "validation unFollow user param")
	}

	if err := c.UserCommandService.Unfollow(&user, p.ID); err != nil {
		return errors.Wrap(err, "failed to delete follow user")
	}

	return ctx.JSON(http.StatusOK, "ok")
}
