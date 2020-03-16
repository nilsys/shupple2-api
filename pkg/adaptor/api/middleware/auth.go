package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	AuthorizeHandlerFunc func(ctx echo.Context, user entity.User) error
	AuthorizeWrapper     func(handlerFunc AuthorizeHandlerFunc) echo.HandlerFunc
)

const authScheme = "JWT "

func NewAuthorizeWrapper(config *config.Config, authService service.AuthService, userRepo repository.UserQueryRepository) AuthorizeWrapper {
	return func(f AuthorizeHandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			cognitoID, err := authService.Authorize(getTokenStr(ctx))
			if err != nil {
				return err
			}
			user, err := userRepo.FindByCognitoID(cognitoID)
			if err != nil {
				return serror.New(err, serror.CodeUnauthorized, "unauthorized")
			}
			return f(ctx, *user)
		}
	}
}

func getTokenStr(ctx echo.Context) string {
	token := ctx.Request().Header.Get(echo.HeaderAuthorization)
	return strings.TrimPrefix(token, authScheme)
}
