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
	*AreaCategoryCommandRepositoryImpl
	*AreaCategoryQueryRepositoryImpl
	*ThemeCategoryCommandRepositoryImpl
	*ThemeCategoryQueryRepositoryImpl
	*CfProjectQueryRepositoryImpl
	*CfProjectCommandRepositoryImpl
	*CfReturnGiftCommandRepositoryImpl
	*ComicCommandRepositoryImpl
	*ComicQueryRepositoryImpl
	*FeatureCommandRepositoryImpl
	*FeatureQueryRepositoryImpl
	*SpotCategoryCommandRepositoryImpl
	*SpotCategoryQueryRepositoryImpl
	*TouristSpotCommandRepositoryImpl
	*TouristSpotQueryRepositoryImpl
	*ShippingQueryRepositoryImpl
	*ShippingCommandRepositoryImpl
	*PostCommandRepositoryImpl
	*PostQueryRepositoryImpl
	*UserQueryRepositoryImpl
	*UserCommandRepositoryImpl
	*VlogCommandRepositoryImpl
	*VlogQueryRepositoryImpl
	*ReviewCommandRepositoryImpl
	*ReviewQueryRepositoryImpl
	*WordpressQueryRepositoryImpl
	*MetasearchAreaQueryRepositoryImpl
}

func InitializeTest(configFilePath config.FilePath) (*Test, error) {
	wire.Build(
		wire.Struct(new(Test), "*"),
		config.GetConfig,
		wire.FieldsOf(new(*config.Config), "AWS"),
		ProvideS3Uploader,
		RepositoriesSet,
	)

	return new(Test), nil
}
