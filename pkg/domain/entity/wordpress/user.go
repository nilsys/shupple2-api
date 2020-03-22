package wordpress

type User struct {
	ID          int            `json:"id"`
	Name        string         `json:"name"`
	URL         string         `json:"url"`
	Description string         `json:"description"`
	Link        string         `json:"link"`
	Slug        string         `json:"slug"`
	AvatarURLs  AvatarURLs     `json:"avatar_urls"`
	Attributes  UserAttributes `json:"acf"`
}

type UserAttributes struct {
	MediaUserID int `json:"media_user_id"`
}

type AvatarURLs struct {
	Num24 string `json:"24"`
	Num48 string `json:"48"`
	Num96 string `json:"96"`
}
