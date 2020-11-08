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

type UserQueryController struct {
	service.UserQueryService
	converter.Converters
}

var UserQueryControllerSet = wire.NewSet(
	wire.Struct(new(UserQueryController), "*"),
)

func (c *UserQueryController) Show(ctx echo.Context, user *entity.UserTiny) error {
	resolve, err := c.UserQueryService.Show(user)
	if err != nil {
		return errors.Wrap(err, "failed show user")
	}

	return ctx.JSON(http.StatusOK, c.ConvertUser2Output(resolve))
}

func (c *UserQueryController) ShowByID(ctx echo.Context) error {
	i := &input.IDParam{}
	if err := BindAndValidate(ctx, i); err != nil {
		return errors.Wrap(err, "failed bind input")
	}

	user, err := c.UserQueryService.ShowByID(i.ID)
	if err != nil {
		return errors.Wrap(err, "failed show user")
	}

	return ctx.JSON(http.StatusOK, c.ConvertUser2Output(user))
}

func (c *UserQueryController) ShowMatchingUser(ctx echo.Context, user *entity.UserTiny) error {
	resolve, err := c.UserQueryService.ShowMatchingUser(user)
	if err != nil {
		return errors.Wrap(err, "failed show user")
	}

	return ctx.JSON(http.StatusOK, c.ConvertUser2Output(resolve))
}

/*
	マッチング後の評価をしていないユーザー一覧
*/
func (c *UserQueryController) ListNotReviewMainMatchingMatchingUser(ctx echo.Context, user *entity.UserTiny) error {
	resolve, err := c.UserQueryService.ListPendingMainMatchingMatchingUser(user)
	if err != nil {
		return errors.Wrap(err, "failed list not review main matching user")
	}
	return ctx.JSON(http.StatusOK, c.ConvertUserList2Output(resolve))
}

func (c *UserQueryController) ListMainMatchingUser(ctx echo.Context, user *entity.UserTiny) error {
	resolve, err := c.UserQueryService.ListMainMatchingUser(user)
	if err != nil {
		return errors.Wrap(err, "failed list main matching user")
	}
	return ctx.JSON(http.StatusOK, c.ConvertUserList2Output(resolve))
}
