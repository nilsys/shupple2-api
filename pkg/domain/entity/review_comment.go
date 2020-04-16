package entity

import (
	"strconv"
	"time"
)

type (
	ReviewComment struct {
		ID            int `gorm:"primary_key"`
		UserID        int
		ReviewID      int
		User          *User `gorm:"foreignkey:UserID"`
		Body          string
		ReplyCount    int
		FavoriteCount int
		CreatedAt     time.Time `gorm:"default:current_timestamp"`
		UpdatedAt     time.Time `gorm:"default:current_timestamp"`
		DeletedAt     *time.Time
	}

	ReviewCommentReply struct {
		ID              int `gorm:"primary_key"`
		UserID          int
		ReviewCommentID int
		User            *User `gorm:"foreignkey:UserID"`
		Body            string
		CreatedAt       time.Time `gorm:"-;default:current_timestamp"`
		UpdatedAt       time.Time `gorm:"-;default:current_timestamp"`
	}

	UserFavoriteReviewComment struct {
		UserID          int `gorm:"primary_key"`
		ReviewCommentID int `gorm:"primary_key"`
	}
)

func NewReviewComment(userID, reviewID int, body string) *ReviewComment {
	// IDはオートインクリメントされるのでなにも入れない
	return &ReviewComment{
		UserID:   userID,
		ReviewID: reviewID,
		Body:     body,
	}
}

func NewUserFavoriteReviewComment(userID, reviewCommentID int) *UserFavoriteReviewComment {
	return &UserFavoriteReviewComment{
		UserID:          userID,
		ReviewCommentID: reviewCommentID,
	}
}

func (r *ReviewComment) IsOwner(userID int) bool {
	return r.UserID == userID
}

func (r *ReviewComment) WebURL() string {
	return "https://stayway.jp/tourism" + strconv.Itoa(r.ID)
}

func (r *ReviewCommentReply) WebURL() string {
	return "https://stayway.jp/tourism" + strconv.Itoa(r.ID)
}
