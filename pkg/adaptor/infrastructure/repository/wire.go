//+build wireinject

package repository

import (
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
	*CategoryCommandRepositoryImpl
	*CategoryQueryRepositoryImpl
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
}

func InitializeTest(configFilePath config.ConfigFilePath) (*Test, error) {
	wire.Build(
		wire.Struct(new(Test), "*"),
		config.GetConfig,
		wire.FieldsOf(new(*config.Config), "AWS"),
		ProvideAWSSession,
		ProvideS3Uploader,
		RepositoriesSet,
	)

	return new(Test), nil
}
