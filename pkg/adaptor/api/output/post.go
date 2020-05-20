package output

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

// フロント返却用Post
type (
	Post struct {
		ID              int                `json:"id"`
		Thumbnail       string             `json:"thumbnail"`
		AreaCategories  []*AreaCategory    `json:"areaCategories"`
		ThemeCategories []*ThemeCategory   `json:"themeCategories"`
		Title           string             `json:"title"`
		Slug            string             `json:"slug"`
		Creator         Creator            `json:"creator"`
		FavoriteCount   int                `json:"favoriteCount"`
		Views           int                `json:"views"`
		HideAds         bool               `json:"hideAds"`
		IsFavorite      bool               `json:"isFavorite"`
		CreatedAt       model.TimeResponse `json:"createdAt"`
		UpdatedAt       model.TimeResponse `json:"updatedAt"`
	}

	// PostQueryController.ListPost()でのみ使用中
	PostWithCategoryDetail struct {
		ID              int                    `json:"id"`
		Thumbnail       string                 `json:"thumbnail"`
		AreaCategories  []*AreaCategoryDetail  `json:"areaCategories"`
		ThemeCategories []*ThemeCategoryDetail `json:"themeCategories"`
		Title           string                 `json:"title"`
		Slug            string                 `json:"slug"`
		Creator         Creator                `json:"creator"`
		FavoriteCount   int                    `json:"favoriteCount"`
		Views           int                    `json:"views"`
		HideAds         bool                   `json:"hideAds"`
		IsFavorite      bool                   `json:"isFavorite"`
		CreatedAt       model.TimeResponse     `json:"createdAt"`
		UpdatedAt       model.TimeResponse     `json:"updatedAt"`
	}

	PostWithCategoryDetailList struct {
		TotalNumber int                       `json:"totalNumber"`
		Posts       []*PostWithCategoryDetail `json:"posts"`
	}

	PostList struct {
		TotalNumber int     `json:"totalNumber"`
		Posts       []*Post `json:"posts"`
	}

	PostShow struct {
		ID              int                    `json:"id"`
		Thumbnail       string                 `json:"thumbnail"`
		Title           string                 `json:"title"`
		Slug            string                 `json:"slug"`
		Body            []*PostBody            `json:"body"`
		TOC             string                 `json:"toc"`
		FavoriteCount   int                    `json:"favoriteCount"`
		FacebookCount   int                    `json:"facebookCount"`
		TwitterCount    int                    `json:"twitterCount"`
		Views           int                    `json:"views"`
		SEOTitle        string                 `json:"seoTitle"`
		SEODescription  string                 `json:"seoDescription"`
		HideAds         bool                   `json:"hideAds"`
		IsFavorited     bool                   `json:"isFavorited"`
		Creator         Creator                `json:"creator"`
		AreaCategories  []*AreaCategoryDetail  `json:"areaCategories"`
		ThemeCategories []*ThemeCategoryDetail `json:"themeCategories"`
		Hashtags        []*Hashtag             `json:"hashtags"`
		UpdatedAt       string                 `json:"updatedAt"`
		CreatedAt       string                 `json:"createdAt"`
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
