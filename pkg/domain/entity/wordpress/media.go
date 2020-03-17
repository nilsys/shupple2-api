package wordpress

type Media struct {
	ID        int    `json:"id"`
	MimeType  string `json:"mime_type"`
	SourceURL string `json:"source_url"`
}
