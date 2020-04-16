package api

import (
	"net/http"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

	"github.com/google/wire"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
)

type UserCommandController struct {
	service.UserCommandService
}

var UserCommandControllerSet = wire.NewSet(
	wire.Struct(new(UserCommandController), "*"),
)

func (c *UserCommandController) SignUp(ctx echo.Context) error {
	p := input.StoreUser{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrap(err, "validation store user input")
	}

	err := c.UserCommandService.SignUp(converter.ConvertStoreUserParamToEntity(&p), p.CognitoToken, p.MigrationCode)
	if err != nil {
		return errors.Wrap(err, "failed to store user")
	}

	return ctx.JSON(http.StatusOK, "ok")
}

// uid以外の要素を更新可能
func (c *UserCommandController) Update(ctx echo.Context, user entity.User) error {
	p := input.UpdateUser{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrap(err, "validation store user image input")
	}

	err := c.UserCommandService.Update(&user, converter.ConvertUpdateUserParamToCmd(&p))
	if err != nil {
		return errors.Wrap(err, "failed to store user image")
	}

	return ctx.JSON(http.StatusOK, "ok")
}

func (c *UserCommandController) Follow(ctx echo.Context, user entity.User) error {
	p := input.FollowParam{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrap(err, "validation follow user input")
	}

	if err := c.UserCommandService.Follow(&user, p.ID); err != nil {
		return errors.Wrap(err, "failed to store follow user")
	}

	return ctx.JSON(http.StatusOK, "ok")
}

func (c *UserCommandController) Unfollow(ctx echo.Context, user entity.User) error {
	p := input.FollowParam{}
	if err := BindAndValidate(ctx, &p); err != nil {
		return errors.Wrap(err, "validation unFollow user input")
	}

	if err := c.UserCommandService.Unfollow(&user, p.ID); err != nil {
		return errors.Wrap(err, "failed to delete follow user")
	}

	return ctx.JSON(http.StatusOK, "ok")
}
