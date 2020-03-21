package entity

import "time"

type (
	ReviewComment struct {
		ID            int `gorm:"column:id"`
		UserID        int
		ReviewID      int
		User          *User `gorm:"foreignkey:UserID"`
		Body          string
		ReplyCount    int
		FavoriteCount int
		CreatedAt     time.Time `gorm:"default:current_timestamp"`
		UpdatedAt     time.Time `gorm:"default:current_timestamp"`
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
