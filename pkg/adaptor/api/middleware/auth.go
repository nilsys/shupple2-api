package middleware

import (
	"strings"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"

	"github.com/google/wire"

	"github.com/labstack/echo/v4"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	AuthorizedHandlerFunc         func(ctx echo.Context, user entity.User) error
	OptionalAuthorizedHandlerFunc func(ctx echo.Context, user entity.OptionalUser) error
)

const authScheme = "JWT "

type Authorize struct {
	AuthService service.AuthService
	UserRepo    repository.UserQueryRepository
}

var AuthorizeSet = wire.NewSet(
	wire.Struct(new(Authorize), "*"),
)

func (a Authorize) Require(f AuthorizedHandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cognitoID, err := a.AuthService.Authorize(getTokenStr(ctx))
		if err != nil {
			return err
		}
		user, err := a.UserRepo.FindByCognitoID(cognitoID)
		if err != nil {
			return serror.New(err, serror.CodeUnauthorized, "unauthorized")
		}
		return f(ctx, *user)
	}
}

func (a Authorize) Optional(f OptionalAuthorizedHandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token := getTokenStr(ctx)
		if token == "" {
			return f(ctx, entity.OptionalUser{})
		}
		cognitoID, err := a.AuthService.Authorize(token)
		if err != nil {
			return err
		}
		user, err := a.UserRepo.FindByCognitoID(cognitoID)
		if err != nil {
			return serror.New(err, serror.CodeUnauthorized, "unauthorized")
		}
		return f(ctx, entity.OptionalUser{User: *user, Authenticated: true})
	}
}

func getTokenStr(ctx echo.Context) string {
	token := ctx.Request().Header.Get(echo.HeaderAuthorization)
	return strings.TrimPrefix(token, authScheme)
}
