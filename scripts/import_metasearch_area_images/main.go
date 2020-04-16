package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/config"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

type Script struct {
	DB         *gorm.DB
	AWSSession *session.Session
	Config     *config.Config
}

type Config struct {
	ImportMetasearchAreaImages struct {
		MetasearchBucket string `yaml:"metasearch_bucket"`
	} `yaml:"import_meatsearch_area_images"`
}

const (
	areaBucketPrefix    = "static/area_category"
	subAreaBucketPrefix = "static/sub_area_category"
)

// local: docker-compose run --rm app go run ./scripts/import_metasearch_area_images
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

	return script.importMetasearchAreaCategoryImages()
}

func (s Script) importMetasearchAreaCategoryImages() error {
	var config Config
	if err := s.Config.Scripts.Decode(&config); err != nil {
		return errors.Wrap(err, "failed to load script config")
	}

	svc := s3.New(s.AWSSession)

	areaReq := &s3.ListObjectsInput{
		Bucket: aws.String(config.ImportMetasearchAreaImages.MetasearchBucket),
		Prefix: aws.String(areaBucketPrefix),
	}

	ao, err := svc.ListObjects(areaReq)
	if err != nil {
		return errors.Wrap(err, "failed to get s3 tmp object")
	}

	subAreaReq := &s3.ListObjectsInput{
		Bucket: aws.String(config.ImportMetasearchAreaImages.MetasearchBucket),
		Prefix: aws.String(subAreaBucketPrefix),
	}

	so, err := svc.ListObjects(subAreaReq)
	if err != nil {
		return errors.Wrap(err, "failed to get s3 tmp object")
	}

	os := append(ao.Contents, so.Contents...)

	for _, metasearchID := range os {
		var category entity.AreaCategory
		if err := s.DB.Where("metasearch_area_id = ?", s.trimKey(*metasearchID.Key)).Or("metasearch_sub_area_id = ?", s.trimKey(*metasearchID.Key)).Or("metasearch_sub_area_id = ?", s.trimKey(*metasearchID.Key)).Or("metasearch_sub_sub_area_id = ?", s.trimKey(*metasearchID.Key)).First(&category).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				continue
			}
			return errors.Wrap(err, "failed to find area_category")
		}
		switch category.Type {
		case model.AreaCategoryTypeArea:
			copyReq := s.copyObject(config.ImportMetasearchAreaImages.MetasearchBucket, config.ImportMetasearchAreaImages.MetasearchBucket+"/"+*metasearchID.Key, areaBucketPrefix+"/"+strconv.Itoa(category.ID))
			_, err = svc.CopyObject(copyReq)
			if err != nil {
				return errors.Wrap(err, "failed to copy s3 object")
			}
		case model.AreaCategoryTypeSubArea:
			copyReq := s.copyObject(config.ImportMetasearchAreaImages.MetasearchBucket, config.ImportMetasearchAreaImages.MetasearchBucket+"/"+*metasearchID.Key, subAreaBucketPrefix+"/"+strconv.Itoa(category.ID))
			_, err = svc.CopyObject(copyReq)
			if err != nil {
				return errors.Wrap(err, "failed to copy s3 object")
			}
		}
	}

	return nil
}

func (s Script) copyObject(bucket, source, key string) *s3.CopyObjectInput {
	return &s3.CopyObjectInput{
		Bucket:     aws.String(bucket),
		CopySource: aws.String(source),
		Key:        aws.String(key),
	}
}

func (s Script) trimKey(str string) string {
	tmp := strings.Replace(str, areaBucketPrefix+"/", "", -1)
	return strings.Replace(tmp, subAreaBucketPrefix+"/", "", -1)
}
