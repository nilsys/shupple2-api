package entity

import (
	"strconv"
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
		Times
	}

	ReviewCommentWithIsFavorite struct {
		ID            int `gorm:"primary_key"`
		UserID        int
		ReviewID      int
		User          *User `gorm:"foreignkey:UserID"`
		Body          string
		ReplyCount    int
		FavoriteCount int
		IsFavorited   bool
		Times
	}

	ReviewCommentReply struct {
		ID              int `gorm:"primary_key"`
		UserID          int
		ReviewCommentID int
		User            *User `gorm:"foreignkey:UserID"`
		Body            string
		Times
	}

	UserFavoriteReviewComment struct {
		UserID          int `gorm:"primary_key"`
		ReviewCommentID int `gorm:"primary_key"`
	}

	UserFavoriteReviewCommentReply struct {
		UserID               int
		ReviewCommentReplyID int
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

func NewUserFavoriteReviewCommentReply(userID, reviewCommentReplyID int) *UserFavoriteReviewCommentReply {
	return &UserFavoriteReviewCommentReply{
		UserID:               userID,
		ReviewCommentReplyID: reviewCommentReplyID,
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

func (r *ReviewCommentWithIsFavorite) TableName() string {
	return "review_comment"
}
