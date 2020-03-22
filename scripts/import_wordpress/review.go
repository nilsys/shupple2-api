package main

import (
	"bytes"
	"regexp"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/satori/uuid"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/command"
)

type (
	review struct {
		ID          int `gorm:"column:ID;primary_key"`
		PostContent string
		PostAuthor  int
		Meta        []*meta `gorm:"foreignkey:PostID"`
	}

	meta struct {
		PostID    int
		MetaKey   string
		MetaValue string
	}
)

func (review) TableName() string {
	return "wp_posts"
}

func (r review) GetMeta(key string) string {
	for _, m := range r.Meta {
		if m.MetaKey == key {
			return m.MetaValue
		}
	}

	return ""
}

func (meta) TableName() string {
	return "wp_postmeta"
}

func (s Script) importReview(wordpressDB *gorm.DB) error {
	const limit = 100
	lastID := 0
	for {
		reviews, err := s.findReview(wordpressDB, limit, lastID)
		if err != nil {
			return errors.Wrap(err, "failed to find reviews")
		}
		if len(reviews) == 0 {
			break
		}

		for _, r := range reviews {
			if r.ID == 98029 || r.ID == 113821 || r.ID == 113097 || r.ID == 113648 {
				continue
			}
			user, err := s.UserQueryRepository.FindByWordpressID(r.PostAuthor)
			if err != nil {
				return errors.WithStack(err)
			}

			reviewCommand, err := s.convertToCreateReviewCommand(r)
			if err != nil {
				return errors.Wrapf(err, "failed to import review(id=%d)", r.ID)
			}

			if err := s.ReviewCommandScenario.Create(user, reviewCommand); err != nil {
				return errors.WithStack(err)
			}
		}

		lastID = reviews[len(reviews)-1].ID
	}

	return nil
}

func (s Script) findReview(wordpressDB *gorm.DB, limit, sinceID int) ([]*review, error) {
	q := wordpressDB.Where("post_type = 'wpcr3_review' AND ID > ?", sinceID).Limit(limit).Preload("Meta")

	rows := make([]*review, 0, limit)
	if err := q.Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find review")
	}

	return rows, nil
}

func (s Script) convertToCreateReviewCommand(review *review) (*command.CreateReview, error) {
	if review.GetMeta("wpcr3_f2") != "" || review.GetMeta("wpcr3_f1") != "" {
		return nil, errors.New("f1 or f2 found")
	}

	media, err := s.uploadMedia(review.ID, review.GetMeta("wpcr3_f3"))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	spotID, err := strconv.Atoi(review.GetMeta("wpcr3_review_post"))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if _, err := s.TouristSpotService.ImportFromWordpressByID(spotID); err != nil {
		return nil, errors.WithStack(err)
	}

	score, err := strconv.Atoi(review.GetMeta("wpcr3_review_rating"))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	travelDate, err := parseTravelDate(review.GetMeta("wpcr3_f4"))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	accompanyingType, err := parseAccompanyingType(review.GetMeta("wpcr3_f5"))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &command.CreateReview{
		TravelDate:    travelDate,
		Accompanying:  accompanyingType,
		TouristSpotID: spotID,
		Score:         score,
		Body:          review.PostContent,
		MediaUUIDs:    media,
	}, nil
}

var yearMonthRegexp = regexp.MustCompile(`(\d+)年(\d+)月`)

func parseTravelDate(travelDate string) (model.YearMonth, error) {
	submatch := yearMonthRegexp.FindStringSubmatch(travelDate)
	if len(submatch) < 3 {
		return model.YearMonth{}, errors.Errorf("invalid travelDate %s", travelDate)
	}

	year, err := strconv.Atoi(submatch[1])
	if err != nil {
		return model.YearMonth{}, errors.WithStack(err)
	}

	month, err := strconv.Atoi(submatch[2])
	if err != nil {
		return model.YearMonth{}, errors.WithStack(err)
	}

	result := time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.Local)
	return model.YearMonth{result}, nil
}

func parseAccompanyingType(accompanyingType string) (model.AccompanyingType, error) {
	switch accompanyingType {
	case "ビジネス":
		return model.AccompanyingTypeBUISINESS, nil
	case "カップル・夫婦":
		return model.AccompanyingTypeCOUPLE, nil
	case "家族":
		return model.AccompanyingTypeFAMILY, nil
	case "友達":
		return model.AccompanyingTypeFRIEND, nil
	case "1人":
		return model.AccompanyingTypeONLY, nil
	case "小さな子連れ":
		return model.AccompanyingTypeWITHCHILD, nil
	}

	return model.AccompanyingType(0), errors.Errorf("invalid accompanyingType %s", accompanyingType)
}

func (s Script) uploadMedia(reviewID int, mediaIDStr string) ([]*command.CreateReviewMedia, error) {
	if mediaIDStr == "" {
		return []*command.CreateReviewMedia{}, nil
	}

	mediaID, err := strconv.Atoi(mediaIDStr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	wpMediaList, err := s.WordpressRepo.FindMediaByIDs([]int{mediaID})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if len(wpMediaList) == 0 {
		return nil, errors.New("wordpress media not found")
	}
	wpMedia := wpMediaList[0]

	media := &command.CreateReviewMedia{
		UUID:     uuid.NewV4().String(),
		MimeType: wpMedia.MimeType,
	}

	mediaBody, err := s.WordpressRepo.DownloadAvatar(wpMedia.SourceURL) // HTTP Get しているだけなので使いまわしてしまう
	if err != nil {
		return nil, errors.WithStack(err)
	}

	key := model.UploadedS3Path(media.UUID)
	_, err = s.MediaUploader.Upload(&s3manager.UploadInput{
		Bucket:      &s.AWSConfig.FilesBucket,
		Key:         &key,
		Body:        bytes.NewReader(mediaBody),
		ContentType: aws.String(media.MimeType),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return []*command.CreateReviewMedia{media}, nil
}
