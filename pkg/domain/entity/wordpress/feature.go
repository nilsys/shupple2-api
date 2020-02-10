package wordpress

type Feature struct {
	ID         int               `json:"id"`
	Date       Time              `json:"date"`
	Modified   Time              `json:"modified"`
	Slug       string            `json:"slug"`
	Title      Text              `json:"title"`
	Content    ProtectableText   `json:"content"`
	FeatureCat []int             `json:"feature_cat"`
	Author     int               `json:"author"`
	Attributes FeatureAttributes `json:"acf"`
}

type FeatureArticle struct {
	ID int `json:"ID"`
}

type FeatureAttributes struct {
	FeatureArticle []FeatureArticle `json:"feature_article"`
}
