package api

import (
	"net/http"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type HealthCheckController struct {
	repository.HealthCheckRepository
	Config *config.Config
}

var HealthCheckControllerSet = wire.NewSet(
	wire.Struct(new(HealthCheckController), "*"),
)

func (c *HealthCheckController) HealthCheck(ctx echo.Context) error {
	if err := c.HealthCheckRepository.CheckDBAlive(); err != nil {
		return errors.Wrap(err, "fail db health check")
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"version": c.Config.Version,
	})
}
