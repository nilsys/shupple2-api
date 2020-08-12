package output

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	PostTiny struct {
		ID             int                `json:"id"`
		Slug           string             `json:"slug"`
		Thumbnail      string             `json:"thumbnail"`
		Title          string             `json:"title"`
		Beginning      string             `json:"summary"`
		TOC            string             `json:"toc"`
		IsSticky       bool               `json:"isSticky"`
		HideAds        bool               `json:"hideAds"`
		FavoriteCount  int                `json:"favoriteCount"`
		FacebookCount  int                `json:"facebookCount"`
		TwitterCount   int                `json:"twitterCount"`
		Views          int                `json:"views"`
		SEOTitle       string             `json:"seoTitle"`
		SEODescription string             `json:"seoDescription"`
		CreatedAt      model.TimeResponse `json:"createdAt"`
		EditedAt       model.TimeResponse `json:"editedAt"`
	}

	// PostQueryController.ListPost()でのみ使用中
	PostWithCategoryDetail struct {
		PostTiny
		IsFavorite      bool                   `json:"isFavorite"`
		Creator         Creator                `json:"creator"`
		AreaCategories  []*AreaCategoryDetail  `json:"areaCategories"`
		ThemeCategories []*ThemeCategoryDetail `json:"themeCategories"`
	}

	PostWithCategoryDetailList struct {
		TotalNumber int                       `json:"totalNumber"`
		Posts       []*PostWithCategoryDetail `json:"posts"`
	}

	PostShow struct {
		PostTiny
		IsFavorite      bool                   `json:"isFavorite"`
		Body            []*PostBody            `json:"body"`
		Creator         Creator                `json:"creator"`
		AreaCategories  []*AreaCategoryDetail  `json:"areaCategories"`
		ThemeCategories []*ThemeCategoryDetail `json:"themeCategories"`
		Hashtags        []*Hashtag             `json:"hashtags"`
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
