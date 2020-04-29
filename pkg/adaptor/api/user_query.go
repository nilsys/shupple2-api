package api

import (
	"net/http"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
)

type UserQueryController struct {
	converter.Converters
	service.UserQueryService
}

var UserQueryControllerSet = wire.NewSet(
	wire.Struct(new(UserQueryController), "*"),
)

func (c *UserQueryController) MyPage(ctx echo.Context, user entity.User) error {
	myPageUser, err := c.UserQueryService.Show(user.UID)
	if err != nil {
		return errors.Wrap(err, "failed to show user")
	}

	return ctx.JSON(http.StatusOK, c.ConvertUserDetailWithCountToOutPut(myPageUser))
}

func (c *UserQueryController) Show(ctx echo.Context) error {
	p := &input.ShowParam{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation show user parameters")
	}

	user, err := c.UserQueryService.Show(p.UID)
	if err != nil {
		return errors.Wrap(err, "failed to show user")
	}

	return ctx.JSON(http.StatusOK, c.ConvertUserDetailWithCountToOutPut(user))
}

func (c *UserQueryController) ShowUserRanking(ctx echo.Context) error {
	p := &input.ListUserRanking{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation show user ranking list parameters")
	}

	q := c.ConvertListRankinUserParamToQuery(p)

	users, err := c.UserQueryService.ShowUserRanking(q)
	if err != nil {
		return errors.Wrap(err, "failed to show user ranking list")
	}

	return ctx.JSON(http.StatusOK, c.ConvertUserRankingToOutput(users))
}

func (c *UserQueryController) ListFollowingUsers(ctx echo.Context) error {
	p := &input.ListFollowUser{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation list following user")
	}

	q := c.ConvertListFollowUserParamToQuery(p)

	users, err := c.UserQueryService.ListFollowing(q)
	if err != nil {
		return errors.Wrap(err, "failed to list user follow")
	}
	return ctx.JSON(http.StatusOK, c.ConvertUsersToUserSummaryList(users))
}

func (c *UserQueryController) ListFollowedUsers(ctx echo.Context) error {
	p := &input.ListFollowUser{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "validation list follower user")
	}

	q := c.ConvertListFollowUserParamToQuery(p)

	users, err := c.UserQueryService.ListFollowed(q)
	if err != nil {
		return errors.Wrap(err, "failed to list user follower")
	}
	return ctx.JSON(http.StatusOK, c.ConvertUsersToUserSummaryList(users))
}

func (c *UserQueryController) ListFavoritePostUser(ctx echo.Context, user entity.OptionalUser) error {
	p := &input.ListFavoriteMediaUser{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "")
	}

	users, err := c.UserQueryService.ListFavoritePostUser(p.MediaID, &user, c.ConvertListFavoriteMediaUserToQuery(p))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, c.ConvertUsersToUserSummaryList(users))
}

func (c *UserQueryController) ListFavoriteReviewUser(ctx echo.Context, user entity.OptionalUser) error {
	p := &input.ListFavoriteMediaUser{}
	if err := BindAndValidate(ctx, p); err != nil {
		return errors.Wrap(err, "")
	}

	users, err := c.UserQueryService.ListFavoriteReviewUser(p.MediaID, &user, c.ConvertListFavoriteMediaUserToQuery(p))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, c.ConvertUsersToUserSummaryList(users))
}
