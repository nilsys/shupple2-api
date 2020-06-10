package output

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

// フロント返却用Review
type (
	Review struct {
		ID            int                `json:"id"`
		InnID         int64              `json:"innId"`
		TouristSpotID int64              `json:"touristSpotId"`
		Score         int                `json:"score"`
		Body          string             `json:"body"`
		FavoriteCount int                `json:"favoriteCount"`
		Media         []ReviewMedia      `json:"media"`
		Views         int                `json:"views"`
		Accompanying  string             `json:"accompanying"`
		CreatedAt     model.TimeResponse `json:"createdAt"`
		UpdatedAt     model.TimeResponse `json:"updatedAt"`
		TravelDate    model.YearMonth    `json:"travelDate"`
		CommentCount  int                `json:"commentCount"`
		IsFavorite    bool               `json:"isFavorite"`
		Hashtag       []Hashtag          `json:"hashtag"`
		Creator       Creator            `json:"creator"`
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
		IsFavorite    bool               `json:"isFavorite"`
		CreatedAt     model.TimeResponse `json:"createdAt"`
	}

	ReviewCommentReply struct {
		ID            int                `json:"id"`
		UserSummary   *UserSummary       `json:"user"`
		Body          string             `json:"body"`
		IsFavorite    bool               `json:"isFavorite"`
		FavoriteCount int                `json:"favoriteCount"`
		CreatedAt     model.TimeResponse `json:"createdAt"`
	}

	ReviewDetailIsFavoriteList struct {
		TotalNumber int       `json:"totalNumber"`
		Reviews     []*Review `json:"reviews"`
	}
)

func NewReviewComment(userSummary *UserSummary, body string, createdAt model.TimeResponse, id, replyCount, favoriteCount int, isFavorite bool) *ReviewComment {
	return &ReviewComment{
		ID:            id,
		UserSummary:   userSummary,
		Body:          body,
		ReplyCount:    replyCount,
		FavoriteCount: favoriteCount,
		IsFavorite:    isFavorite,
		CreatedAt:     createdAt,
	}
}
