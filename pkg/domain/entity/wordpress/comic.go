package wordpress

type Comic struct {
	ID            int              `json:"id"`
	Date          JSTTime          `json:"date"`
	Modified      UTCTime          `json:"modified"`
	Slug          URLEscapedString `json:"slug"`
	Status        Status           `json:"status"`
	Type          string           `json:"type"`
	Link          string           `json:"link"`
	Title         Text             `json:"title"`
	Content       ProtectableText  `json:"content"`
	Author        int              `json:"author"`
	FeaturedMedia int              `json:"featured_media"`
}
