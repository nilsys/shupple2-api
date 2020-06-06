package wordpress

import (
	"bytes"
	"encoding/json"

	"github.com/pkg/errors"
)

type Post struct {
	ID            int              `json:"id"`
	Date          JSTTime          `json:"date"`
	DateGmt       UTCTime          `json:"date_gmt"`
	GUID          Text             `json:"guid"`
	Modified      JSTTime          `json:"modified"`
	ModifiedGmt   UTCTime          `json:"modified_gmt"`
	Slug          URLEscapedString `json:"slug"`
	Status        Status           `json:"status"`
	Type          string           `json:"type"`
	Title         Text             `json:"title"`
	Content       ProtectableText  `json:"content"`
	Author        int              `json:"author"`
	FeaturedMedia int              `json:"featured_media"`
	Sticky        bool             `json:"sticky"`
	Meta          PostMeta         `json:"meta"`
	Categories    []int            `json:"categories"`
	Tags          []int            `json:"tags"`
}

type PostMeta struct {
	SEOTitle           string `json:"seo_title"`
	MetaDescription    string `json:"meta_description"`
	IsAdsRemovedInPage bool   `json:"is_ads_removed_in_page"`
}

func (p *PostMeta) UnmarshalJSON(body []byte) error {
	if bytes.Equal(body, arrayJSONBytes) {
		return nil
	}

	type Alias PostMeta
	return errors.WithStack(json.Unmarshal(body, (*Alias)(p)))
}
