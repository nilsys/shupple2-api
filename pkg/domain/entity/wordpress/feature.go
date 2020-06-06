package wordpress

type Feature struct {
	ID            int               `json:"id"`
	Date          JSTTime           `json:"date"`
	Modified      UTCTime           `json:"modified"`
	Slug          URLEscapedString  `json:"slug"`
	Status        Status            `json:"status"`
	Title         Text              `json:"title"`
	Content       ProtectableText   `json:"content"`
	FeatureCat    []int             `json:"feature_cat"`
	Author        int               `json:"author"`
	FeaturedMedia int               `json:"featured_media"`
	Attributes    FeatureAttributes `json:"acf"`
}

type FeatureArticle struct {
	ID int `json:"ID"`
}

type FeatureAttributes struct {
	FeatureArticle []FeatureArticle `json:"feature_article"`
}
