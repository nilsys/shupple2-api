package entity

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/config"
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
		CommentCount  int
		Views         int
		TravelDate    time.Time
		Accompanying  model.AccompanyingType
		Medias        ReviewMediaList  `gorm:"foreignkey:ReviewID"`
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

	ReviewDetail struct {
		Review
		User    *User      `gorm:"foreignkey:UserID"`
		Hashtag []*Hashtag `gorm:"many2many:review_hashtag;jointable_foreignkey:review_id;"`
	}

	ReviewDetailWithIsFavorite struct {
		Review
		IsFavorite bool
		User       *User      `gorm:"foreignkey:UserID"`
		Hashtag    []*Hashtag `gorm:"many2many:review_hashtag;jointable_foreignkey:review_id;"`
	}

	UserFavoriteReview struct {
		UserID   int
		ReviewID int
	}

	ReviewDetailWithIsFavoriteList struct {
		TotalNumber int
		Reviews     []*ReviewDetailWithIsFavorite
	}

	ReviewMediaList []*ReviewMedia
)

func NewUserFavoriteReview(userID, reviewID int) *UserFavoriteReview {
	return &UserFavoriteReview{
		UserID:   userID,
		ReviewID: reviewID,
	}
}

func NewReviewMedia(id string, mimeType string, priority int) *ReviewMedia {
	return &ReviewMedia{
		ID:       id,
		MimeType: mimeType,
		Priority: priority,
	}
}

func (r *ReviewMedia) S3Path() string {
	return fmt.Sprintf("review/%d/%s", r.ReviewID, r.ID)
}

func (r *ReviewMedia) URL(filesURL config.URL) string {
	filesURL.Path = r.S3Path()
	return filesURL.String()
}

func (r *ReviewDetail) TableName() string {
	return "review"
}

func (r *Review) IsOwner(userID int) bool {
	return r.UserID == userID
}

func (r *Review) HashHashtagIDs() bool {
	return len(r.HashtagIDs) > 0
}

// TODO: 仮置き
func (r *Review) WebURL() string {
	return "https://stayway.jp/tourism/" + strconv.Itoa(r.ID)
}

func (rdi ReviewDetailWithIsFavorite) TableName() string {
	return "review"
}

func (list ReviewMediaList) Sort() {
	sort.Slice(list, func(i, j int) bool { return list[i].Priority < list[j].Priority })
}
