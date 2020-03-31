package response

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

// フロント返却用Post
type (
	Post struct {
		ID                      int                `json:"id"`
		Thumbnail               string             `json:"thumbnail"`
		AreaCategories          []*AreaCategory    `json:"areaCategories"`
		ThemeCategoryCategories []*ThemeCategory   `json:"themeCategoryCategories"`
		Title                   string             `json:"title"`
		Creator                 Creator            `json:"creator"`
		LikeCount               int                `json:"likeCount"`
		Views                   int                `json:"views"`
		HideAds                 bool               `json:"hideAds"`
		CreatedAt               model.TimeResponse `json:"createdAt"`
		UpdatedAt               model.TimeResponse `json:"updatedAt"`
	}

	PostList struct {
		TotalNumber int     `json:"totalNumber"`
		Posts       []*Post `json:"posts"`
	}

	PostShow struct {
		ID                      int              `json:"id"`
		Thumbnail               string           `json:"thumbnail"`
		Title                   string           `json:"title"`
		Body                    []*PostBody      `json:"body"`
		TOC                     string           `json:"toc"`
		FavoriteCount           int              `json:"favoriteCount"`
		FacebookCount           int              `json:"facebookCount"`
		TwitterCount            int              `json:"twitterCount"`
		Views                   int              `json:"views"`
		SEOTitle                string           `json:"seoTitle"`
		SEODescription          string           `json:"seoDescription"`
		HideAds                 bool             `json:"hideAds"`
		Creator                 Creator          `json:"creator"`
		AreaCategories          []*AreaCategory  `json:"areaCategories"`
		ThemeCategoryCategories []*ThemeCategory `json:"themeCategoryCategories"`
		Hashtags                []*Hashtag       `json:"hashtags"`
		UpdatedAt               string           `json:"updatedAt"`
		CreatedAt               string           `json:"createdAt"`
	}

	PostBody struct {
		Page int    `json:"page"`
		Body string `json:"body"`
	}
)

func NewPostBody(page int, body string) *PostBody {
	return &PostBody{
		Page: page,
		Body: body,
	}
}
