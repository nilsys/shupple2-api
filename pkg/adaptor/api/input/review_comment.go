package input

type (
	CreateReviewCommentReply struct {
		ReviewCommentID int    `param:"id" validate:"required"`
		Body            string `json:"body"`
	}

	FavoriteReviewComment struct {
		ReviewCommentID int `param:"id"`
	}

	ListReviewCommentReply struct {
		ReviewCommentID int `param:"id" validate:"required"`
	}
)
