package entity

import "time"

type (
	ReviewComment struct {
		ID        int `gorm:"column:id"`
		UserID    int
		ReviewID  int
		User      *User     `gorm:"foreignkey:UserID"`
		Body      string    `gorm:"column:body"`
		CreatedAt time.Time `gorm:"default:current_timestamp"`
		UpdatedAt time.Time `gorm:"default:current_timestamp"`
	}
)

func NewReviewComment(userID, reviewID int, body string) *ReviewComment {
	// IDはオートインクリメントされるのでなにも入れない
	return &ReviewComment{
		UserID:    userID,
		ReviewID:  reviewID,
		Body:      body,
	}
}
