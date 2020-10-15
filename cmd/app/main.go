package main

import (
	"github.com/labstack/echo/v4"
	"github.com/uma-co82/shupple2-api/pkg/adaptor/api"
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
	Config                *config.Config
	DB                    *gorm.DB
	Echo                  *echo.Echo
	HealthCheckController api.HealthCheckController
}

func run() error {
	//app, err := InitializeApp(config.DefaultConfigFilePath)
	//if err != nil {
	//	return errors.Wrap(err, "failed to initialize app")
	//}
	//logger.Configure(app.Config.Logger)
	//
	//if app.Config.Migrate.Auto {
	//	if err := repository.MigrateUp(app.Config.Database, app.Config.Migrate.FilesDir); err != nil {
	//		return errors.Wrap(err, "failed to migrate up")
	//	}
	//}
	//
	//app.Echo.Debug = app.Config.IsDev()
	//app.Echo.HTTPErrorHandler = api.CreateErrorHandler(app.Config.Env)
	//app.Echo.Use(middleware.RequestID())
	//app.Echo.Use(middleware.Logger())
	//app.Echo.Use(middleware.CORS())
	//setRoutes(app)
	//
	//return app.Echo.Start(":3000")
	return nil
}

func setRoutes(app *App) {
	api := app.Echo.Group("/api/v1")

	api.GET("/healthcheck", app.HealthCheckController.HealthCheck)
}
