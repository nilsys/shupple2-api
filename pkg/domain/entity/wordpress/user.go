package wordpress

import (
	"bytes"
	"encoding/json"

	"github.com/pkg/errors"
)

type User struct {
	ID          int              `json:"id"`
	Name        string           `json:"name"`
	URL         string           `json:"url"`
	Description string           `json:"description"`
	Link        string           `json:"link"`
	Slug        URLEscapedString `json:"slug"`
	AvatarURLs  AvatarURLs       `json:"avatar_urls"`
	Meta        UserMeta         `json:"meta"`
}

type UserMeta struct {
	Facebook     string `json:"facebook"`
	Twitter      string `json:"twitter"`
	Instagram    string `json:"instagram"`
	Youtube      string `json:"youtube"`
	WPUserAvatar int    `json:"wp_user_avatar"`
}

type AvatarURLs struct {
	Num24 string `json:"24"`
	Num48 string `json:"48"`
	Num96 string `json:"96"`
}

func (u *UserMeta) UnmarshalJSON(body []byte) error {
	if bytes.Equal(body, arrayJSONBytes) {
		return nil
	}

	type Alias UserMeta
	return errors.WithStack(json.Unmarshal(body, (*Alias)(u)))
}
