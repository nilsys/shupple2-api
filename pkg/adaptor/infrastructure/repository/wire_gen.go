// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package repository

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"net/url"
)

import (
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Injectors from wire.go:

func InitializeTest(configFilePath config.FilePath) (*Test, error) {
	configConfig, err := config.GetConfig(configFilePath)
	if err != nil {
		return nil, err
	}
	db, err := ProvideDB(configConfig)
	if err != nil {
		return nil, err
	}
	session, err := ProvideAWSSession(configConfig)
	if err != nil {
		return nil, err
	}
	uploader := ProvideS3Uploader(session)
	dao := DAO{
		UnderlyingDB: db,
	}
	areaCategoryCommandRepositoryImpl := &AreaCategoryCommandRepositoryImpl{
		DAO: dao,
	}
	areaCategoryQueryRepositoryImpl := &AreaCategoryQueryRepositoryImpl{
		DB: db,
	}
	themeCategoryCommandRepositoryImpl := &ThemeCategoryCommandRepositoryImpl{
		DAO: dao,
	}
	themeCategoryQueryRepositoryImpl := &ThemeCategoryQueryRepositoryImpl{
		DB: db,
	}
	comicCommandRepositoryImpl := &ComicCommandRepositoryImpl{
		DAO: dao,
	}
	comicQueryRepositoryImpl := &ComicQueryRepositoryImpl{
		DB: db,
	}
	featureCommandRepositoryImpl := &FeatureCommandRepositoryImpl{
		DAO: dao,
	}
	featureQueryRepositoryImpl := &FeatureQueryRepositoryImpl{
		DB: db,
	}
	spotCategoryCommandRepositoryImpl := &SpotCategoryCommandRepositoryImpl{
		DAO: dao,
	}
	spotCategoryQueryRepositoryImpl := &SpotCategoryQueryRepositoryImpl{
		DB: db,
	}
	touristSpotCommandRepositoryImpl := &TouristSpotCommandRepositoryImpl{
		DAO: dao,
	}
	touristSpotQueryRepositoryImpl := &TouristSpotQueryRepositoryImpl{
		DB: db,
	}
	postCommandRepositoryImpl := &PostCommandRepositoryImpl{
		DAO: dao,
	}
	postQueryRepositoryImpl := &PostQueryRepositoryImpl{
		DB: db,
	}
	userQueryRepositoryImpl := &UserQueryRepositoryImpl{
		DB: db,
	}
	aws := configConfig.AWS
	userCommandRepositoryImpl := &UserCommandRepositoryImpl{
		DAO:           dao,
		MediaUploader: uploader,
		AWSConfig:     aws,
		AWSSession:    session,
	}
	vlogCommandRepositoryImpl := &VlogCommandRepositoryImpl{
		DAO: dao,
	}
	vlogQueryRepositoryImpl := &VlogQueryRepositoryImpl{
		DB: db,
	}
	reviewCommandRepositoryImpl := &ReviewCommandRepositoryImpl{
		DAO:        dao,
		AWSSession: session,
		AWSConfig:  aws,
	}
	reviewQueryRepositoryImpl := &ReviewQueryRepositoryImpl{
		DB: db,
	}
	wordpress := _wireWordpressValue
	staywayMedia := _wireStaywayMediaValue
	wordpressQueryRepositoryImpl := NewWordpressQueryRepositoryImpl(wordpress, staywayMedia)
	test := &Test{
		Config:                             configConfig,
		DB:                                 db,
		AWS:                                session,
		Uploader:                           uploader,
		AreaCategoryCommandRepositoryImpl:  areaCategoryCommandRepositoryImpl,
		AreaCategoryQueryRepositoryImpl:    areaCategoryQueryRepositoryImpl,
		ThemeCategoryCommandRepositoryImpl: themeCategoryCommandRepositoryImpl,
		ThemeCategoryQueryRepositoryImpl:   themeCategoryQueryRepositoryImpl,
		ComicCommandRepositoryImpl:         comicCommandRepositoryImpl,
		ComicQueryRepositoryImpl:           comicQueryRepositoryImpl,
		FeatureCommandRepositoryImpl:       featureCommandRepositoryImpl,
		FeatureQueryRepositoryImpl:         featureQueryRepositoryImpl,
		SpotCategoryCommandRepositoryImpl:  spotCategoryCommandRepositoryImpl,
		SpotCategoryQueryRepositoryImpl:    spotCategoryQueryRepositoryImpl,
		TouristSpotCommandRepositoryImpl:   touristSpotCommandRepositoryImpl,
		TouristSpotQueryRepositoryImpl:     touristSpotQueryRepositoryImpl,
		PostCommandRepositoryImpl:          postCommandRepositoryImpl,
		PostQueryRepositoryImpl:            postQueryRepositoryImpl,
		UserQueryRepositoryImpl:            userQueryRepositoryImpl,
		UserCommandRepositoryImpl:          userCommandRepositoryImpl,
		VlogCommandRepositoryImpl:          vlogCommandRepositoryImpl,
		VlogQueryRepositoryImpl:            vlogQueryRepositoryImpl,
		ReviewCommandRepositoryImpl:        reviewCommandRepositoryImpl,
		ReviewQueryRepositoryImpl:          reviewQueryRepositoryImpl,
		WordpressQueryRepositoryImpl:       wordpressQueryRepositoryImpl,
	}
	return test, nil
}

var (
	_wireWordpressValue = config.Wordpress{
		BaseURL: config.URL{
			URL: url.URL{
				Scheme: "https",
				Host:   "stg-admin.stayway.jp",
				Path:   "/tourism",
			},
		},
	}
	_wireStaywayMediaValue = config.StaywayMedia{
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
	}
)

// wire.go:

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
	*SpotCategoryCommandRepositoryImpl
	*SpotCategoryQueryRepositoryImpl
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

var configValuesSet = wire.NewSet(wire.Value(config.Wordpress{
	BaseURL: config.URL{
		URL: url.URL{
			Scheme: "https",
			Host:   "stg-admin.stayway.jp",
			Path:   "/tourism",
		},
	},
}), wire.Value(config.StaywayMedia{
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
