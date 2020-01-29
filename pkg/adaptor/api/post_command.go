package api

import (
	"net/http"

	"github.com/stayway-corp/stayway-media-api/pkg/application/service"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

type (
	PostCommandController struct {
		PostService service.PostCommandService
	}
)

var PostCommandControllerSet = wire.NewSet(
	wire.Struct(new(PostCommandController), "*"),
)

func (c *PostCommandController) Store(ctx echo.Context) error {
	// MEMO: 仮置き
	return ctx.JSON(http.StatusOK, nil)
}
