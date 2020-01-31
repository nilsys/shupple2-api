package response

// フロント返却用Post
type Post struct {
	ID               int      `json:"id"`
	Thumbnail        string   `json:"thumbnail"`
	AreaCategories   []string `json:"areaCategories"`
	ThemeCategories  []string `json:"themeCategories"`
	Title            string   `json:"title"`
	CreatorThumbnail string   `json:"creatorThumbnail"`
	CreatorName      string   `json:"creatorName"`
	LikeCount        string   `json:"likeCount"`
	UpdatedAt        string   `json:"updatedAt"`
}
