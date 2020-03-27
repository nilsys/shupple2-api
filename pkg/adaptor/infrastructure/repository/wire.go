//+build wireinject

package repository

import (
	"net/url"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
)

type Test struct {
	Config   *config.Config
	DB       *gorm.DB
	AWS      *session.Session
	Uploader *s3manager.Uploader
	*AreaCategoryCommandRepositoryImpl
	*AreaCategoryQueryRepositoryImpl
	*ThemeCategoryCommandRepositoryImpl
	*ThemeCategoryQueryRepositoryImpl
	*ComicCommandRepositoryImpl
	*ComicQueryRepositoryImpl
	*FeatureCommandRepositoryImpl
	*FeatureQueryRepositoryImpl
	*LcategoryCommandRepositoryImpl
	*LcategoryQueryRepositoryImpl
	*TouristSpotCommandRepositoryImpl
	*TouristSpotQueryRepositoryImpl
	*PostCommandRepositoryImpl
	*PostQueryRepositoryImpl
	*UserQueryRepositoryImpl
	*UserCommandRepositoryImpl
	*VlogCommandRepositoryImpl
	*VlogQueryRepositoryImpl
	*ReviewCommandRepositoryImpl
	*ReviewQueryRepositoryImpl
	*WordpressQueryRepositoryImpl
}

var configValuesSet = wire.NewSet(
	wire.Value(config.Wordpress{
		BaseURL: config.URL{
			URL: url.URL{
				Scheme: "https",
				Host:   "stg-admin.stayway.jp",
				Path:   "/tourism",
			},
		},
	}),
	wire.Value(config.StaywayMedia{
		BaseURL: config.URL{
			URL: url.URL{
				Scheme: "https",
				Host:   "stg.stayway.jp",
				Path:   "/tourism",
			},
		},
		FilesURL: config.URL{
			URL: url.URL{
				Scheme: "https",
				Host:   "stg-files.stayway.jp",
			},
		},
	}),
)

func InitializeTest(configFilePath config.FilePath) (*Test, error) {
	wire.Build(
		wire.Struct(new(Test), "*"),
		config.GetConfig,
		wire.FieldsOf(new(*config.Config), "AWS"),
		configValuesSet,
		ProvideS3Uploader,
		RepositoriesSet,
	)

	return new(Test), nil
}
