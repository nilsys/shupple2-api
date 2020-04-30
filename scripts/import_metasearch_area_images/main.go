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
		MediaBucket      string `yaml:"media_bucket"`
	} `yaml:"import_meatsearch_area_images"`
}

const (
	metasearchAreaBucket    = "area_category"
	metasearchSubAreaBucket = "sub_area_category"
	areaBucketPrefix        = "static/area_category"
	subAreaBucketPrefix     = "static/sub_area_category"
	subSubAreaBucketPrefix  = "static/sub_sub_area_category"
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

	// scriptなので動かすときにcredential必要
	svc := s3.New(s.AWSSession)

	areaReq := &s3.ListObjectsInput{
		Bucket: aws.String(config.ImportMetasearchAreaImages.MetasearchBucket),
		Prefix: aws.String(metasearchAreaBucket),
	}

	ao, err := svc.ListObjects(areaReq)
	if err != nil {
		return errors.Wrap(err, "failed to get s3 tmp object")
	}

	subAreaReq := &s3.ListObjectsInput{
		Bucket: aws.String(config.ImportMetasearchAreaImages.MetasearchBucket),
		Prefix: aws.String(metasearchSubAreaBucket),
	}

	so, err := svc.ListObjects(subAreaReq)
	if err != nil {
		return errors.Wrap(err, "failed to get s3 tmp object")
	}

	os := append(ao.Contents, so.Contents...)

	// subareaに対して
	for _, metasearchID := range os {
		var category entity.AreaCategory
		key := s.trimKey(*metasearchID.Key)
		if key == "" {
			continue
		}
		if err := s.DB.Where("metasearch_area_id = ? AND type = ?", key, model.AreaCategoryTypeSubArea).First(&category).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				continue
			}
			return errors.Wrap(err, "failed to find area_category")
		}
		copyReq := s.copyObject(config.ImportMetasearchAreaImages.MetasearchBucket+"/"+*metasearchID.Key, config.ImportMetasearchAreaImages.MediaBucket, subAreaBucketPrefix+"/"+strconv.Itoa(category.ID)+"/"+s.getFileName(*metasearchID.Key))
		_, err = svc.CopyObject(copyReq)
		if err != nil {
			return errors.Wrap(err, "failed to copy s3 object")
		}
	}

	// sub_sub_areaに対して
	for _, metasearchID := range os {
		var category entity.AreaCategory
		key := s.trimKey(*metasearchID.Key)
		if key == "" {
			continue
		}
		if err := s.DB.Where("metasearch_sub_area_id = ? AND type = ?", key, model.AreaCategoryTypeSubSubArea).First(&category).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				continue
			}
			return errors.Wrap(err, "failed to find area_category")
		}
		copyReq := s.copyObject(config.ImportMetasearchAreaImages.MetasearchBucket+"/"+*metasearchID.Key, config.ImportMetasearchAreaImages.MediaBucket, subSubAreaBucketPrefix+"/"+strconv.Itoa(category.ID)+"/"+s.getFileName(*metasearchID.Key))
		_, err = svc.CopyObject(copyReq)
		if err != nil {
			return errors.Wrap(err, "failed to copy s3 object")
		}
	}
	return nil
}

func (s Script) copyObject(from, bucket, key string) *s3.CopyObjectInput {
	return &s3.CopyObjectInput{
		CopySource: aws.String(from),
		Bucket:     aws.String(bucket),
		Key:        aws.String(key),
		ACL:        aws.String(s3.ObjectCannedACLPublicRead),
	}
}

func (s Script) trimKey(str string) string {
	tmp := strings.Split(str, "/")
	if len(tmp) < 2 {
		return ""
	}
	return tmp[1]
}

func (s Script) getFileName(str string) string {
	tmp := strings.Split(str, "/")
	if len(tmp) < 2 {
		return ""
	}
	return tmp[2]
}
