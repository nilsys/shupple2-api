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

		avatar, err := s.downloadAvatar(int(user.WordpressID.Int64))
		if err != nil {
			return errors.WithStack(err)
		}
		defer avatar.Body.Close()

		if err := s.UserRepo.StoreWithAvatar(user, avatar.Body, avatar.ContentType); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func (s Script) downloadAvatar(userID int) (*wordpress.MediaBody, error) {
	wpUsers, err := s.WordpressRepo.FindUsersByIDs([]int{userID})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if len(wpUsers) == 0 {
		return nil, errors.New("not found user")
	}
	wpUser := wpUsers[0]

	return s.WordpressRepo.DownloadAvatar(wpUser.AvatarURLs.Num96)
}
