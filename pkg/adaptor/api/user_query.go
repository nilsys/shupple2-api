package api

import (
	"net/http"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

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

func (c *UserQueryController) Show(ctx echo.Context) error {
	p := &param.ShowParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation show user ranking list parameters")
	}

	user, err := c.UserQueryService.Show(p.ID)
	if err != nil {
		return errors.Wrap(err, "validation show user parameter")
	}

	return ctx.JSON(http.StatusOK, converter.ConvertUserDetailWithCountToOutPut(user))
}

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

func (c *UserQueryController) ListFollowingUsers(ctx echo.Context) error {
	p := &param.ListFollowUser{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation list following user")
	}

	q := converter.ConvertListFollowUserParamToQuery(p)

	users, err := c.UserQueryService.ListFollowing(q)
	if err != nil {
		return errors.Wrap(err, "failed to list user follow")
	}
	return ctx.JSON(http.StatusOK, converter.ConvertUsersToFollowUsers(users))
}

func (c *UserQueryController) ListFollowedUsers(ctx echo.Context) error {
	p := &param.ListFollowUser{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation list follower user")
	}

	q := converter.ConvertListFollowUserParamToQuery(p)

	users, err := c.UserQueryService.ListFollowed(q)
	if err != nil {
		return errors.Wrap(err, "failed to list user follower")
	}
	return ctx.JSON(http.StatusOK, converter.ConvertUsersToFollowUsers(users))
}

func (c *UserQueryController) ListFavoritePostUser(ctx echo.Context, user entity.OptionalUser) error {
	p := &param.ListFavoriteMediaUser{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "")
	}

	users, err := c.UserQueryService.ListFavoritePostUser(p.MediaID, &user, converter.ConvertListFavoriteMediaUserToQuery(p))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, converter.ConvertUsersToFollowUsers(users))
}

func (c *UserQueryController) ListFavoriteReviewUser(ctx echo.Context, user entity.OptionalUser) error {
	p := &param.ListFavoriteMediaUser{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "")
	}

	users, err := c.UserQueryService.ListFavoriteReviewUser(p.MediaID, &user, converter.ConvertListFavoriteMediaUserToQuery(p))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, converter.ConvertUsersToFollowUsers(users))
}
