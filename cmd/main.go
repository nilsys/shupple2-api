package main

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/logger"

	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
)

const (
	configFilePath = "config.yaml"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	app, err := InitializeApp(configFilePath)
	if err != nil {
		return errors.Wrap(err, "failed to initialize app")
	}
	logger.Configure(app.Config.Logger)

	app.Echo.Debug = app.Config.IsDev()
	app.Echo.HTTPErrorHandler = api.ErrorHandler
	app.Echo.Use(middleware.Logger())
	app.Echo.Use(middleware.CORS())
	setRoutes(app)

	return app.Echo.Start(":3000")
}

func setRoutes(app *App) {
	api := app.Echo.Group("/api")

	api.POST("/posts", app.PostCommandController.Store)
	api.GET("/posts/:id", app.PostQueryController.FindByID)
	api.GET("/review", app.ReviewQueryController.LisReview)
}
