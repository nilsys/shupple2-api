package api

import (
	"net/http"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
)

type UserQueryController struct {
	service.UserQueryService
}

var UserQueryControllerSet = wire.NewSet(
	wire.Struct(new(UserQueryController), "*"),
)

func (c *UserQueryController) ShowUserRanking(ctx echo.Context) error {
	p := &param.ListUserRanking{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation show user ranking list parameters")
	}

	q := converter.ConvertListRankinUserParamToQuery(p)

	users, err := c.UserQueryService.ShowUserRanking(q)
	if err != nil {
		return errors.Wrap(err, "failed to show user ranking list")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertUserRankingToOutput(users))
}
