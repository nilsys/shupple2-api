package response

// フロント返却用Comic
type Comic struct {
	ID        int    `json:"id"`
	Slug      string `json:"slug"`
	Title     string `json:"title"`
	Thumbnail string `json:"thumbnail"`
}

// フロント返却用Comic詳細
type ShowComic struct {
	ID        int     `json:"id"`
	Slug      string  `json:"slug"`
	Title     string  `json:"title"`
	Thumbnail string  `json:"thumbnail"`
	Body      string  `json:"body"`
	CreatedAt string  `json:"createdAt"`
	Creator   Creator `json:"creator"`
}
