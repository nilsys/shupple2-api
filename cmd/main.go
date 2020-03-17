package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api"
	staywayMiddleware "github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/middleware"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/logger"
	"github.com/stayway-corp/stayway-media-api/pkg/config"

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

type App struct {
	Config                      *config.Config
	Echo                        *echo.Echo
	AuthorizeWrapper            staywayMiddleware.AuthorizeWrapper
	PostCommandController       api.PostCommandController
	PostQueryController         api.PostQueryController
	CategoryQueryController     api.CategoryQueryController
	ComicQueryController        api.ComicQueryController
	ReviewQueryController       api.ReviewQueryController
	ReviewCommandController     api.ReviewCommandController
	HashtagQueryController      api.HashtagQueryController
	SearchQueryController       api.SearchQueryController
	FeatureQueryController      api.FeatureQueryController
	VlogQueryController         api.VlogQueryController
	UserQueryController         api.UserQueryController
	UserCommandController       api.UserCommandController
	HealthCheckController       api.HealthCheckController
	WordpressCallbackController api.WordpressCallbackController
	S3CommandController         api.S3CommandController
	TouristSpotQueryController  api.TouristSpotQueryController
	InteresetQueryController    api.InterestQueryController
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
	auth := app.AuthorizeWrapper

	{
		posts := api.Group("/posts")
		posts.GET("", app.PostQueryController.ListPost)
		posts.POST("", app.PostCommandController.Store)
		posts.GET("/:id", app.PostQueryController.Show)
	}

	{
		users := api.Group("/users")
		users.POST("", app.UserCommandController.SignUp)
		users.GET("/:id/posts/feed", app.PostQueryController.ListFeedPost)
		users.GET("/:id/reviews/feed", app.ReviewQueryController.ListFeedReview)
		users.GET("/ranking", app.UserQueryController.ShowUserRanking)
		users.GET("/:id/followee", app.UserQueryController.ListFolloweeUsers)
		users.GET("/:id/follower", app.UserQueryController.ListFollowerUsers)
	}

	{
		reviews := api.Group("/reviews")
		reviews.GET("", app.ReviewQueryController.LisReview)
		reviews.GET("/:id/comment", app.ReviewQueryController.ListReviewCommentByReviewID)
		reviews.POST("", app.ReviewCommandController.Store)
		reviews.GET("/:id", app.ReviewQueryController.ShowReview)
	}

	{
		comics := api.Group("/comics")
		comics.GET("", app.ComicQueryController.ListComic)
		comics.GET("/:id", app.ComicQueryController.Show)
	}

	{
		touristSpots := api.Group("/tourist_spots")
		touristSpots.GET("", app.TouristSpotQueryController.ListTouristSpot)
		touristSpots.GET("/recommend", app.TouristSpotQueryController.ListRecommendTouristSpot)
		touristSpots.GET("/:id", app.TouristSpotQueryController.Show)
	}

	{
		features := api.Group("/feature_posts")

		features.GET("", app.FeatureQueryController.ListFeature)
		features.GET("/:id", app.FeatureQueryController.ShowQuery)
	}

	{
		areas := api.Group("/areas")
		areas.GET("", app.CategoryQueryController.ListArea)
		areas.GET("/:id", app.CategoryQueryController.ShowAreaByID)
	}

	{
		subAreas := api.Group("/sub_areas")
		subAreas.GET("", app.CategoryQueryController.ListSubArea)
		subAreas.GET("/:id", app.CategoryQueryController.ShowSubAreaByID)
	}

	{
		subSubAreas := api.Group("/sub_sub_areas")
		subSubAreas.GET("", app.CategoryQueryController.ListSubSubArea)
		subSubAreas.GET("/:id", app.CategoryQueryController.ShowSubSubAreaByID)
	}

	{
		vlogs := api.Group("/vlogs")
		vlogs.GET("", app.VlogQueryController.ListVlog)
		vlogs.GET("/:id", app.VlogQueryController.Show)
	}

	{
		hashtags := api.Group("/hashtags")
		hashtags.GET("/recommend", app.HashtagQueryController.ListRecommendHashtag)
	}

	{
		interests := api.Group("/interests")
		interests.GET("", app.InteresetQueryController.ListAll)
	}

	api.POST("/s3", auth(app.S3CommandController.Post))
	api.GET("/healthcheck", app.HealthCheckController.HealthCheck)
	api.GET("/search/suggestions", app.SearchQueryController.ListSearchSuggestion)
	api.POST(
		"/wordpress/import",
		app.WordpressCallbackController.Import,
		staywayMiddleware.KeyAuth(app.Config.Wordpress.CallbackKey),
	)
}
