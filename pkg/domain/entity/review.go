package entity

import (
	"fmt"
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

type (
	// table: review
	Review struct {
		ID            int
		UserID        int
		TouristSpotID int `gorm:"default:nil"`
		InnID         int `gorm:"default:nil"`
		Score         int
		MediaCount    int
		Body          string
		FavoriteCount int
		Views         int
		TravelDate    time.Time
		Accompanying  model.AccompanyingType
		Medias        []*ReviewMedia   `gorm:"foreignkey:ReviewID"`
		HashtagIDs    []*ReviewHashtag `gorm:"foreignkey:ReviewID"`
		CreatedAt     time.Time        `gorm:"-;default:current_timestamp"`
		UpdatedAt     time.Time        `gorm:"-;default:current_timestamp"`
		DeletedAt     *time.Time
	}

	// table: review_media
	ReviewMedia struct {
		ID       string
		MimeType string
		Priority int
		ReviewID int `gorm:"review_id"`
	}

	ReviewHashtag struct {
		ReviewID  int `gorm:"primary_key"`
		HashtagID int `gorm:"primary_key"`
	}

	// 参照用Review
	QueryReview struct {
		Review
		CommentCount int
		User         *User      `gorm:"foreignkey:UserID"`
		Hashtag      []*Hashtag `gorm:"many2many:review_hashtag;jointable_foreignkey:review_id;"`
	}
)

func NewReviewMedia(id string, mimeType string, priority int) *ReviewMedia {
	return &ReviewMedia{
		ID:       id,
		MimeType: mimeType,
		Priority: priority,
	}
}

func NewReviewHashtag(reviewID, hashtagID int) *ReviewHashtag {
	return &ReviewHashtag{
		ReviewID:  reviewID,
		HashtagID: hashtagID,
	}
}

// TODO: 仮置き
func (r *ReviewMedia) GenerateURL() string {
	return "https://stayway.jp/image/" + r.ID
}

func (r *QueryReview) TableName() string {
	return "review"
}

func (reviewMedia *ReviewMedia) S3Path() string {
	return fmt.Sprintf("review/%d/%s", reviewMedia.ReviewID, reviewMedia.ID)
}
