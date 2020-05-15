package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api"
	staywayMiddleware "github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/middleware"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/logger"
	"github.com/stayway-corp/stayway-media-api/pkg/config"

	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

type App struct {
	Config                          *config.Config
	DB                              *gorm.DB
	Echo                            *echo.Echo
	AuthorizeWrapper                staywayMiddleware.Authorize
	PostQueryController             api.PostQueryController
	PostFavoriteCommandController   api.PostFavoriteCommandController
	CategoryQueryController         api.CategoryQueryController
	ComicQueryController            api.ComicQueryController
	ReviewQueryController           api.ReviewQueryController
	ReviewCommandController         api.ReviewCommandController
	ReviewFavoriteCommandController api.ReviewFavoriteCommandController
	RssController                   api.RSSController
	HashtagQueryController          api.HashtagQueryController
	HashtagCommandController        api.HashtagCommandController
	SearchQueryController           api.SearchQueryController
	FeatureQueryController          api.FeatureQueryController
	VlogQueryController             api.VlogQueryController
	UserQueryController             api.UserQueryController
	UserCommandController           api.UserCommandController
	HealthCheckController           api.HealthCheckController
	WordpressCallbackController     api.WordpressCallbackController
	SitemapController               api.SitemapController
	S3CommandController             api.S3CommandController
	TouristSpotQueryController      api.TouristSpotQueryController
	InteresetQueryController        api.InterestQueryController
	ThemeQueryController            api.ThemeQueryController
	AreaQueryController             api.AreaQueryController
	InnQueryController              api.InnQueryController
	NoticeQueryController           api.NoticeQueryController
	ReportCommandController         api.ReportCommandController
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
	app.Echo.HTTPErrorHandler = api.ErrorHandler
	app.Echo.Use(middleware.Logger())
	app.Echo.Use(middleware.CORS())
	setRoutes(app)

	return app.Echo.Start(":3000")
}

func setRoutes(app *App) {
	api := app.Echo.Group("/api/v1/tourism")
	auth := app.AuthorizeWrapper

	{
		posts := api.Group("/posts")
		posts.GET("", auth.Optional(app.PostQueryController.ListPost))
		posts.GET("/:id", auth.Optional(app.PostQueryController.Show))
		posts.GET("/:slug/slug", app.PostQueryController.ShowBySlug)
		posts.PUT("/:id/favorite", auth.Require(app.PostFavoriteCommandController.Store))
		posts.DELETE("/:id/favorite", auth.Require(app.PostFavoriteCommandController.Delete))
	}

	{
		users := api.Group("/users")
		users.POST("", app.UserCommandController.SignUp)
		users.PUT("", auth.Require(app.UserCommandController.Update))
		users.GET("", auth.Require(app.UserQueryController.MyPage))
		users.GET("/:uid", app.UserQueryController.Show)
		users.GET("/ranking", app.UserQueryController.ShowUserRanking)
		users.POST("/:id/follow", auth.Require(app.UserCommandController.Follow))
		users.DELETE("/:id/follow", auth.Require(app.UserCommandController.Unfollow))
		users.GET("/:id/following", app.UserQueryController.ListFollowingUsers)
		users.GET("/:id/followed", app.UserQueryController.ListFollowedUsers)
		users.GET("/:id/feed/posts", auth.Optional(app.PostQueryController.ListFeedPost))
		users.GET("/:id/feed/reviews", auth.Optional(app.ReviewQueryController.ListFeedReview))
		users.GET("/:id/favorite/posts", auth.Optional(app.PostQueryController.ListFavoritePost))
		users.GET("/:id/favorite/reviews", auth.Optional(app.ReviewQueryController.ListFavoriteReview))
		users.GET("/favorite/reviews/:id", auth.Optional(app.UserQueryController.ListFavoriteReviewUser))
		users.GET("/favorite/posts/:id", auth.Optional(app.UserQueryController.ListFavoritePostUser))
	}

	{
		reviews := api.Group("/reviews")
		reviews.GET("", auth.Optional(app.ReviewQueryController.LisReview))
		reviews.DELETE("/comment/:id", auth.Require(app.ReviewCommandController.DeleteReviewComment))
		reviews.POST("/:id/comment", auth.Require(app.ReviewCommandController.StoreReviewComment))
		reviews.GET("/:id/comment", auth.Optional(app.ReviewQueryController.ListReviewCommentByReviewID))
		reviews.POST("", auth.Require(app.ReviewCommandController.Store))
		reviews.PUT("/:id", auth.Require(app.ReviewCommandController.Update))
		reviews.DELETE("/:id", auth.Require(app.ReviewCommandController.Delete))
		reviews.GET("/:id", auth.Optional(app.ReviewQueryController.ShowReview))
		reviews.POST("/:id/comment", auth.Require(app.ReviewCommandController.StoreReviewComment))
		reviews.GET("/comment/:id/reply", app.ReviewQueryController.ListReviewCommentReply)
		reviews.POST("/comment/:id/reply", auth.Require(app.ReviewCommandController.StoreReviewCommentReply))
		reviews.PUT("/comment/:id/favorite", auth.Require(app.ReviewCommandController.FavoriteReviewComment))
		reviews.DELETE("/comment/:id/favorite", auth.Require(app.ReviewCommandController.UnfavoriteReviewComment))
		reviews.PUT("/:id/favorite", auth.Require(app.ReviewFavoriteCommandController.Store))
		reviews.DELETE("/:id/favorite", auth.Require(app.ReviewFavoriteCommandController.Delete))
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
		categories := api.Group("/categories")
		categories.GET("/:slug/slug", app.CategoryQueryController.ShowBySlug)
	}

	{
		areas := api.Group("/areas")
		areas.GET("", app.AreaQueryController.ListArea)
		areas.GET("/:id", app.AreaQueryController.ShowAreaByID)
	}

	{
		subAreas := api.Group("/sub_areas")
		subAreas.GET("", app.AreaQueryController.ListSubArea)
		subAreas.GET("/:id", app.AreaQueryController.ShowSubAreaByID)
	}

	{
		subSubAreas := api.Group("/sub_sub_areas")
		subSubAreas.GET("", app.AreaQueryController.ListSubSubArea)
		subSubAreas.GET("/:id", app.AreaQueryController.ShowSubSubAreaByID)
	}

	{
		vlogs := api.Group("/vlogs")
		vlogs.GET("", app.VlogQueryController.ListVlog)
		vlogs.GET("/:id", app.VlogQueryController.Show)
	}

	{
		hashtags := api.Group("/hashtags")
		hashtags.GET("/recommend", app.HashtagQueryController.ListRecommendHashtag)
		hashtags.POST("/:id/follow", auth.Require(app.HashtagCommandController.FollowHashtag))
		hashtags.DELETE("/:id/follow", auth.Require(app.HashtagCommandController.UnfollowHashtag))
	}

	{
		interests := api.Group("/interests")
		interests.GET("", app.InteresetQueryController.ListAll)
	}

	{
		inns := api.Group("/inns")
		inns.GET("", app.InnQueryController.ListByParams)
	}

	{
		notices := api.Group("/notice")
		notices.GET("", auth.Require(app.NoticeQueryController.ListNotices))
	}

	{
		themes := api.Group("/themes")
		themes.GET("", app.ThemeQueryController.List)
	}

	{
		subThemes := api.Group("/sub_themes")
		subThemes.GET("", app.ThemeQueryController.ListSubThemeByParentID)
	}

	{
		reports := api.Group("/reports")
		reports.POST("", auth.Require(app.ReportCommandController.Report))
		reports.POST("/submit", app.ReportCommandController.MarkAsDone, staywayMiddleware.KeyAuth(app.Config.Slack.CallbackKey))
	}

	api.POST("/s3", auth.Require(app.S3CommandController.Post))
	api.GET("/search/suggestions", app.SearchQueryController.ListSearchSuggestion)
	api.POST(
		"/wordpress/import",
		app.WordpressCallbackController.Import,
		staywayMiddleware.KeyAuth(app.Config.Wordpress.CallbackKey),
	)

	app.Echo.GET("/tourism/sitemap*", app.SitemapController.Show)
	app.Echo.GET("/tourism/smartnews", app.RssController.Show)
	app.Echo.GET("/tourism/wp-content/plugins/google-sitemap-generator/sitemap.xsl", app.SitemapController.Show)

	app.Echo.GET("/healthcheck", app.HealthCheckController.HealthCheck)
}
