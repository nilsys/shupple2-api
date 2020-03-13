package wordpress

type Media struct {
	ID           int          `json:"id"`
	MediaDetails MediaDetails `json:"media_details"`
}

type MediaDetails struct {
	Width  int        `json:"width"`
	Height int        `json:"height"`
	File   string     `json:"file"`
	Sizes  MediaSizes `json:"sizes"`
}

type MediaSizes struct {
	Full MediaDetail `json:"full"`
}

type MediaDetail struct {
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	File      string `json:"file"`
	MimeType  string `json:"mime_type"`
	SourceURL string `json:"source_url"`
}
