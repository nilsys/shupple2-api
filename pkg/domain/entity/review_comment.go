package entity

type (
	ReviewCommentTiny struct {
		ID            int `gorm:"primary_key"`
		UserID        int
		ReviewID      int
		User          *User `gorm:"foreignkey:UserID"`
		Body          string
		ReplyCount    int
		FavoriteCount int
		Times
	}

	ReviewCommentDetail struct {
		ReviewCommentTiny
		Review *Review `gorm:"foreignkey:ID;association_foreignkey:ReviewID"`
	}

	ReviewCommentWithIsFavorite struct {
		ID            int `gorm:"primary_key"`
		UserID        int
		ReviewID      int
		User          *User `gorm:"foreignkey:UserID"`
		Body          string
		ReplyCount    int
		FavoriteCount int
		IsFavorite    bool
		Times
	}

	ReviewCommentReplyTiny struct {
		ID              int `gorm:"primary_key"`
		UserID          int
		ReviewCommentID int
		User            *User `gorm:"foreignkey:UserID"`
		Body            string
		FavoriteCount   int
		Times
	}

	ReviewCommentReplyDetail struct {
		ReviewCommentReplyTiny
		ReviewCommentDetail *ReviewCommentDetail `gorm:"foreignkey:ID;association_foreignkey:ReviewCommentID"`
	}

	ReviewCommentReplyWithIsFavorite struct {
		ID              int `gorm:"primary_key"`
		UserID          int
		ReviewCommentID int
		User            *User `gorm:"foreignkey:UserID"`
		Body            string
		FavoriteCount   int
		IsFavorite      bool
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

func NewReviewComment(userID, reviewID int, body string) *ReviewCommentTiny {
	return &ReviewCommentTiny{
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

func (r *ReviewCommentReplyTiny) TableName() string {
	return "review_comment_reply"
}

func (r *ReviewCommentWithIsFavorite) TableName() string {
	return "review_comment"
}

func (r *ReviewCommentReplyWithIsFavorite) TableName() string {
	return "review_comment_reply"
}

func (r *ReviewCommentTiny) TableName() string {
	return "review_comment"
}

func (r *ReviewCommentDetail) TableName() string {
	return "review_comment"
}
