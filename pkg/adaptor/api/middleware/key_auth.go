package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const keyAuthQueryParamName = "key"

func KeyAuth(correctKey string) echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "query:" + keyAuthQueryParamName,
		Validator: func(key string, c echo.Context) (bool, error) {
			return key == correctKey, nil
		},
	})
}
