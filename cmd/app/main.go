package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"github.com/uma-co82/shupple2-api/pkg/adaptor/api"
	middleware2 "github.com/uma-co82/shupple2-api/pkg/adaptor/api/middleware"
	"github.com/uma-co82/shupple2-api/pkg/adaptor/infrastructure/repository"
	"github.com/uma-co82/shupple2-api/pkg/adaptor/logger"
	"github.com/uma-co82/shupple2-api/pkg/application/service"
	"github.com/uma-co82/shupple2-api/pkg/config"

	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	log.Println("teminated.")
}

type App struct {
	Config *config.Config
	DB     *gorm.DB
	Echo   *echo.Echo
	middleware2.Authorize
	api.HealthCheckController
	api.UserQueryController
	api.UserCommandController
	api.ArrangeScheduleRequestCommandController
	api.ArrangeScheduleRequestQueryController
	service.TransactionService
}

func run() error {
	app, err := InitializeApp(config.DefaultConfigFilePath)
	if err != nil {
		return errors.Wrap(err, "failed to initialize app")
	}
	logger.Configure(app.Config.Logger)

	if app.Config.Migrate.Auto {
		if err := repository.MigrateUp(app.Config.Database, app.Config.Migrate.FilesDir); err != nil {
			return errors.Wrap(err, "failed to migrate up")
		}
	}

	app.Echo.Debug = app.Config.IsDev()
	app.Echo.HTTPErrorHandler = api.CreateErrorHandler(app.Config.Env)
	app.Echo.Use(middleware.RequestID())
	app.Echo.Use(middleware.Logger())
	app.Echo.Use(middleware.CORS())
	setRoutes(app)

	return app.Echo.Start(":3000")
}

func setRoutes(app *App) {
	api := app.Echo.Group("/api/v1")
	auth := app.Authorize

	users := api.Group("/users")

	{
		users.POST("", app.UserCommandController.SignUp)
		users.GET("/:id", app.UserQueryController.ShowByID)
		users.POST("/matching", auth.Auth(app.UserCommandController.Matching))
		users.GET("/matching", auth.Auth(app.UserQueryController.ShowMatchingUser))
	}

	arrangeSchedule := api.Group("/arrange_schedule")

	{
		arrangeSchedule.POST("", auth.Auth(app.ArrangeScheduleRequestCommandController.Store))
	}

	api.GET("/healthcheck", app.HealthCheckController.HealthCheck)
}
