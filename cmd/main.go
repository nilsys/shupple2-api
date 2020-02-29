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

	api.GET("/users/ranking", app.UserQueryController.ShowUserRanking)
	api.GET("/posts", app.PostQueryController.ListPost)
	api.POST("/posts", app.PostCommandController.Store)
	api.GET("/posts/:id", app.PostQueryController.Show)
	api.GET("/posts/:id", app.PostQueryController.ShowQuery)
	// api.GET("/posts/:id", app.PostQueryController.Show)
	api.GET("/user/:id/posts/feed", app.PostQueryController.ListFeedPost)
	api.GET("/reviews", app.ReviewQueryController.LisReview)
	api.GET("/user/:id/reviews/feed", app.ReviewQueryController.ListFeedReview)
	api.GET("/comics", app.ComicQueryController.ListComic)
	api.GET("/comics/:id", app.ComicQueryController.Show)
	api.GET("/search/suggestions", app.SearchQueryController.ShowSearchSuggestionList)
	api.GET("/feature_posts", app.FeatureQueryController.ListFeature)
	api.GET("/feature_posts/:id", app.FeatureQueryController.ShowQuery)
	api.GET("/vlogs", app.VlogQueryController.ListVlog)
	api.GET("/hashtags/recommend", app.HashtagQueryController.ListRecommendHashtag)
	api.GET("/healthcheck", app.HealthCheckController.HealthCheck)
}
