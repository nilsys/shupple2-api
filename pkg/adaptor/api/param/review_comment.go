package param

type (
	CreateReviewCommentReply struct {
		ReviewCommentID int    `param:"id"`
		Body            string `json:"body"`
	}

	FavoriteReviewComment struct {
		ReviewCommentID int `param:"id"`
	}

	ListReviewCommentReply struct {
		ReviewCommentID int `param:"id" validate:"required"`
	}
)
