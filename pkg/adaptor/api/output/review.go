package output

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

// フロント返却用Review
type (
	// TODO: usecaseが増えたら、命名考える
	Review struct {
		ID            int             `json:"id"`
		InnID         int             `json:"innId"`
		TouristSpotID int             `json:"touristSpotId"`
		Score         int             `json:"score"`
		Body          string          `json:"body"`
		FavoriteCount int             `json:"favoriteCount"`
		Media         []ReviewMedia   `json:"media"`
		Views         int             `json:"views"`
		Accompanying  string          `json:"accompanying"`
		UpdatedAt     string          `json:"udpatedAt"`
		TravelDate    model.YearMonth `json:"travelDate"`
		CommentCount  int             `json:"commentCount"`
		Hashtag       []Hashtag       `json:"hashtag"`
		Creator       Creator         `json:"creator"`
		// TODO:
		AssociatedContent ReviewTarget `json:"associatedContent"`
	}

	ReviewMedia struct {
		UUID     string `json:"uuid"`
		Mime     string `json:"mime"`
		URL      string `json:"url"`
		Priority int    `json:"priority"`
	}

	// レビューが紐付くinn or tourist_spotの情報
	ReviewTarget struct {
		Type string `json:"type"`
		Name string `json:"name"`
		Area string `json:"area"`
	}

	ReviewComment struct {
		ID            int                `json:"id"`
		UserSummary   *UserSummary       `json:"user"`
		Body          string             `json:"body"`
		ReplyCount    int                `json:"replyCount"`
		FavoriteCount int                `json:"favoriteCount"`
		CreatedAt     model.TimeResponse `json:"createdAt"`
	}

	ReviewCommentReply struct {
		ID          int                `json:"id"`
		UserSummary *UserSummary       `json:"user"`
		Body        string             `json:"body"`
		CreatedAt   model.TimeResponse `json:"createdAt"`
	}
)

func NewReviewComment(userSummary *UserSummary, body string, createdAt model.TimeResponse, id, replyCount, favoriteCount int) *ReviewComment {
	return &ReviewComment{
		ID:            id,
		UserSummary:   userSummary,
		Body:          body,
		ReplyCount:    replyCount,
		FavoriteCount: favoriteCount,
		CreatedAt:     createdAt,
	}
}
