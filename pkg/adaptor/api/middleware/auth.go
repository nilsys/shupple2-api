package middleware

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/uma-co82/shupple2-api/pkg/domain/model/serror"
	"github.com/uma-co82/shupple2-api/pkg/domain/repository"

	"github.com/labstack/echo/v4"
	"github.com/uma-co82/shupple2-api/pkg/application"
	"github.com/uma-co82/shupple2-api/pkg/domain/entity"
)

type (
	AuthorizedHandleFunc func(ctx echo.Context, user entity.UserTiny) error
)

type Authorize struct {
	application.AuthService
	repository.UserQueryRepository
}

const authScheme = "JWT "

func (a Authorize) Auth(f AuthorizedHandleFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		token := getTokenStr(context)
		if token == "" {
			return serror.New(nil, serror.CodeUnauthorized, "unauthorized")
		}
		firebaseID, err := a.AuthService.Authorize(token)
		if err != nil {
			return err
		}

		user, err := a.UserQueryRepository.FindByFirebaseID(firebaseID)
		if err != nil {
			return errors.Wrap(err, "failed ref user")
		}
		return f(context, user)
	}
}

func getTokenStr(ctx echo.Context) string {
	token := ctx.Request().Header.Get(echo.HeaderAuthorization)
	return strings.TrimPrefix(token, authScheme)
}
