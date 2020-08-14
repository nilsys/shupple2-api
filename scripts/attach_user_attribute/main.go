package main

import (
	"flag"
	"log"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"

	"github.com/pkg/errors"

	"github.com/jinzhu/gorm"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type Script struct {
	DB *gorm.DB
	repository.UserQueryRepository
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	script, err := InitializeScript("./config.yaml")
	if err != nil {
		return errors.Wrap(err, "faled to initialize script")
	}
	return script.AttachUserAttribute()
}

func (s Script) AttachUserAttribute() error {
	op := flag.String("attr", "", "attribute")
	uid := flag.String("uid", "", "uid")
	flag.Parse()
	if *op == "" || *uid == "" {
		return serror.New(nil, serror.CodeInvalidParam, "must attribute arg")
	}

	attr, err := model.ParseUserAttribute(*op)
	if err != nil {
		return serror.New(err, serror.CodeInvalidParam, "must attribute type")
	}

	user, err := s.UserQueryRepository.FindByUID(*uid)
	if err != nil {
		return errors.Wrap(err, "failed ref user")
	}

	userAttribute := entity.UserAttribute{
		UserID:    user.ID,
		Attribute: attr,
	}

	if err := s.DB.Save(&userAttribute).Error; err != nil {
		return errors.Wrap(err, "failed store user_attribute")
	}

	return nil
}
