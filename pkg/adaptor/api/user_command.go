package api

import (
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

func (c *UserCommandController) SingUp(ctx echo.Context) error {
	i := &input.RegisterUser{}
	if err := BindAndValidate(ctx, i); err != nil {
		return errors.Wrap(err, "failed bind input")
	}

	if err := c.UserCommandService.SignUp(c.ConvertRegisterUserInput2Cmd(i), i.FirebaseToken); err != nil {

	}
}
