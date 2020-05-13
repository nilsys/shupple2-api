package main

import (
	"log"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	repositoryImpl "github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity/wordpress"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

const (
	debug   = false
	perPage = 100
)

type Script struct {
	DB            *gorm.DB
	Config        *config.Config
	MediaUploader *s3manager.Uploader
	AWSConfig     config.AWS
	WordpressRepo repository.WordpressQueryRepository
	UserRepo      repository.UserCommandRepository
	UserService   service.UserCommandService
}

type CustomizedWordpressRepo struct {
	*repositoryImpl.WordpressQueryRepositoryImpl
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	script, err := InitializeScript("./config.yaml")
	if err != nil {
		return errors.Wrap(err, "failed to initialize script")
	}

	return script.Run()
}

func (s Script) Run() error {
	var rows []*entity.User
	if err := s.DB.Find(&rows).Error; err != nil {
		return errors.WithStack(err)
	}

	for _, user := range rows {
		if !user.WordpressID.Valid {
			continue
		}

		wpUser, err := s.WordpressRepo.FindUserByID(int(user.WordpressID.Int64))
		if err != nil {
			return errors.WithStack(err)
		}

		if err := s.storeWithAvatar(user, wpUser); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

// このScriptのためだけにexportするのは微妙なので、UserCommandRepositoryImplからコピペ
func (s Script) storeWithAvatar(user *entity.User, wpUser *wordpress.User) error {
	var (
		avatar *wordpress.MediaBody
		err    error
	)

	if wpUser.Meta.WPUserAvatar != 0 {
		avatar, err = s.WordpressRepo.FetchMediaBodyByID(wpUser.Meta.WPUserAvatar)
	} else {
		avatar, err = s.WordpressRepo.FetchResource(wpUser.AvatarURLs.Num96)
	}
	if err != nil {
		return errors.Wrap(err, "failed to download avatar")
	}
	defer avatar.Body.Close()

	if err := s.UserRepo.StoreWithAvatar(user, avatar.Body, avatar.ContentType); err != nil {
		return errors.Wrap(err, "faield to store user avatar")
	}

	return nil
}
