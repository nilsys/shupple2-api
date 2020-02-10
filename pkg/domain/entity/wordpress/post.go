package wordpress

type Post struct {
	ID            int             `json:"id"`
	Date          Time            `json:"date"`
	DateGmt       Time            `json:"date_gmt"`
	GUID          Text            `json:"guid"`
	Modified      Time            `json:"modified"`
	ModifiedGmt   Time            `json:"modified_gmt"`
	Slug          string          `json:"slug"`
	Status        string          `json:"status"`
	Type          string          `json:"type"`
	Link          string          `json:"link"`
	Title         Text            `json:"title"`
	Content       ProtectableText `json:"content"`
	Excerpt       ProtectableText `json:"excerpt"`
	Author        int             `json:"author"`
	FeaturedMedia int             `json:"featured_media"`
	CommentStatus string          `json:"comment_status"`
	PingStatus    string          `json:"ping_status"`
	Sticky        bool            `json:"sticky"`
	Template      string          `json:"template"`
	Format        string          `json:"format"`
	Meta          []interface{}   `json:"meta"`
	Categories    []int           `json:"categories"`
	ArticleTheme  []interface{}   `json:"article_theme"`
	Links         PostLinks       `json:"_links"`
}

type About struct {
	Href string `json:"href"`
}

type Author struct {
	Embeddable bool   `json:"embeddable"`
	Href       string `json:"href"`
}

type Replies struct {
	Embeddable bool   `json:"embeddable"`
	Href       string `json:"href"`
}

type VersionHistory struct {
	Count int    `json:"count"`
	Href  string `json:"href"`
}

type PredecessorVersion struct {
	ID   int    `json:"id"`
	Href string `json:"href"`
}

type WpFeaturedmedia struct {
	Embeddable bool   `json:"embeddable"`
	Href       string `json:"href"`
}

type WpAttachment struct {
	Href string `json:"href"`
}

type WpTerm struct {
	Taxonomy   string `json:"taxonomy"`
	Embeddable bool   `json:"embeddable"`
	Href       string `json:"href"`
}

type Curies struct {
	Name      string `json:"name"`
	Href      string `json:"href"`
	Templated bool   `json:"templated"`
}

type PostLinks struct {
	Self               []Self               `json:"self"`
	Collection         []Collection         `json:"collection"`
	About              []About              `json:"about"`
	Author             []Author             `json:"author"`
	Replies            []Replies            `json:"replies"`
	VersionHistory     []VersionHistory     `json:"version-history"`
	PredecessorVersion []PredecessorVersion `json:"predecessor-version"`
	WpFeaturedmedia    []WpFeaturedmedia    `json:"wp:featuredmedia"`
	WpAttachment       []WpAttachment       `json:"wp:attachment"`
	WpTerm             []WpTerm             `json:"wp:term"`
	Curies             []Curies             `json:"curies"`
}
