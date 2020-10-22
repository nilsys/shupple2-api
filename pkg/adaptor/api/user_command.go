package api

import (
	"net/http"

	"github.com/uma-co82/shupple2-api/pkg/domain/entity"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/uma-co82/shupple2-api/pkg/adaptor/api/converter"
	"github.com/uma-co82/shupple2-api/pkg/adaptor/api/input"
	"github.com/uma-co82/shupple2-api/pkg/application/service"
)

type UserCommandController struct {
	service.UserCommandService
	converter.Converters
}

var UserCommandControllerSet = wire.NewSet(
	wire.Struct(new(UserCommandController), "*"),
)

func (c *UserCommandController) SignUp(ctx echo.Context) error {
	i := &input.RegisterUser{}
	if err := BindAndValidate(ctx, i); err != nil {
		return errors.Wrap(err, "failed bind input")
	}

	if err := c.UserCommandService.SignUp(c.ConvertRegisterUserInput2Cmd(i), i.FirebaseToken); err != nil {
		return errors.Wrap(err, "failed sign up user")
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (c *UserCommandController) Matching(ctx echo.Context, user *entity.UserTiny) error {
	if err := c.UserCommandService.Matching(user); err != nil {
		return errors.Wrap(err, "failed matching user")
	}

	return ctx.NoContent(http.StatusNoContent)
}
