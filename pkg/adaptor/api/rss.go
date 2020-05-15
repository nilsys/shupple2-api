package api

import (
	"net/http"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
)

type RSSController struct {
	service.RSSService
}

var RSSControllerSet = wire.NewSet(
	wire.Struct(new(RSSController), "*"),
)

func (c *RSSController) Show(ctx echo.Context) error {
	sitemap, contentType, err := c.RSSService.Show()
	if err != nil {
		return errors.Wrap(err, "failed to show rss")
	}

	return ctx.Blob(http.StatusOK, contentType, sitemap)
}
