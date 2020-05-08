package api

import (
	"net/http"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/converter"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
)

type SitemapController struct {
	converter.Converters
	service.SitemapService
}

var SitemapControllerSet = wire.NewSet(
	wire.Struct(new(SitemapController), "*"),
)

func (c *SitemapController) Show(ctx echo.Context) error {
	path := ctx.Request().URL.Path
	sitemap, contentType, err := c.SitemapService.Show(path)
	if err != nil {
		return errors.Wrapf(err, "failed to show sitemap(%s)", path)
	}

	return ctx.Blob(http.StatusOK, contentType, sitemap)
}
