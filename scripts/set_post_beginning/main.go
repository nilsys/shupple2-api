package main

import (
	"log"
	"regexp"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"

	"github.com/pkg/errors"

	"github.com/jinzhu/gorm"
)

const (
	limit        = 100
	beginningCnt = 200
)

var (
	tagRe = regexp.MustCompile(`<.+?>`)
	emRe  = regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	wemRe = regexp.MustCompile(`[\s\p{Zs}]{2,}`)
)

type Script struct {
	DB *gorm.DB
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
	lastID := 0
	for {
		posts, err := s.FindNotHaveBeginning(lastID)
		if err != nil {
			return errors.Wrap(err, "failed find posts")
		}

		if len(posts) == 0 {
			break
		}

		for _, post := range posts {
			body, err := s.FindTopPage(post.ID)
			if err != nil {
				if serror.IsErrorCode(err, serror.CodeNotFound) {
					continue
				}
				return errors.Wrap(err, "failed find post_body")
			}

			resolve, err := model.PostBodyToBeginning(body.Body)
			if err != nil {
				return errors.Wrap(err, "failed extract body")
			}

			if err := s.UpdateBeginning(post.ID, resolve); err != nil {
				return errors.Wrap(err, "failed update post.beginning")
			}
		}

		lastID = posts[len(posts)-1].ID
	}
	return nil
}

func (s Script) FindNotHaveBeginning(lastID int) ([]*entity.Post, error) {
	var rows []*entity.Post
	if err := s.DB.Where("id > ? AND beginning = ''", lastID).Limit(limit).Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed find post")
	}
	return rows, nil
}

func (s Script) FindTopPage(id int) (*entity.PostBody, error) {
	var row entity.PostBody
	if err := s.DB.Where("post_id = ?", id).Order("page").First(&row).Error; err != nil {
		return nil, repository.ErrorToFindSingleRecord(err, "post_bdy(post_id=%d)", id)
	}
	return &row, nil
}

func (s Script) UpdateBeginning(id int, body string) error {
	if err := s.DB.Exec("UPDATE post SET beginning = ? WHERE id = ?", body, id).Error; err != nil {
		return errors.Wrap(err, "failed update post.beginning")
	}
	return nil
}
