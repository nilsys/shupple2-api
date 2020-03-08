package response

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

// フロント返却用Post
type (
	Post struct {
		ID              int                `json:"id"`
		Thumbnail       string             `json:"thumbnail"`
		AreaCategories  []Category         `json:"areaCategories"`
		ThemeCategories []Category         `json:"themeCategories"`
		Title           string             `json:"title"`
		Creator         Creator            `json:"creator"`
		LikeCount       int                `json:"likeCount"`
		UpdatedAt       model.TimeResponse `json:"updatedAt"`
	}

	PostShow struct {
		ID              int        `json:"id"`
		Thumbnail       string     `json:"thumbnail"`
		Title           string     `json:"title"`
		Body            []PostBody `json:"body"`
		TOC             string     `json:"toc"`
		FavoriteCount   int        `json:"favoriteCount"`
		FacebookCount   int        `json:"facebookCount"`
		TwitterCount    int        `json:"twitterCount"`
		Views           int        `json:"views"`
		Creator         Creator    `json:"creator"`
		AreaCategories  []Category `json:"areaCategories"`
		ThemeCategories []Category `json:"themeCategories"`
		Hashtags        []Hashtag  `json:"hashtags"`
		UpdatedAt       string     `json:"updatedAt"`
		CreatedAt       string     `json:"createdAt"`
	}

	PostBody struct {
		Page int    `json:"page"`
		Body string `json:"body"`
	}
)

func NewPostBody(page int, body string) PostBody {
	return PostBody{
		Page: page,
		Body: body,
	}
}
