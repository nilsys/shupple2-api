package response

// フロント返却用Post
type (
	Post struct {
		ID              int        `json:"id"`
		Thumbnail       string     `json:"thumbnail"`
		AreaCategories  []Category `json:"areaCategories"`
		ThemeCategories []Category `json:"themeCategories"`
		Title           string     `json:"title"`
		Creator         Creator    `json:"creator"`
		LikeCount       int        `json:"likeCount"`
		UpdatedAt       string     `json:"updatedAt"`
	}
)
