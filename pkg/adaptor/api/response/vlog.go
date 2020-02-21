package response

type Vlog struct {
	Thumbnail       string     `json:"thumbnail"`
	AreaCategories  []Category `json:"areaCategories"`
	ThemeCategories []Category `json:"themeCategories"`
	Title           string     `json:"title"`
}
