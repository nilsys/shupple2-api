package wordpress

import "io"

type Media struct {
	ID        int    `json:"id"`
	MimeType  string `json:"mime_type"`
	SourceURL string `json:"source_url"`
}

type MediaBody struct {
	ContentType string
	Body        io.ReadCloser
}

type PhotoGalleryItem struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Caption      string `json:"caption"`
	FullImageURL string `json:"full_image_url"`
}
