package wordpress

import (
	"bytes"
	"encoding/json"

	"github.com/pkg/errors"
)

type (
	Vlog struct {
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
		Template      string           `json:"template"`
		Categories    []int            `json:"categories"`
		Attributes    VlogAttributes   `json:"acf"`
	}

	Creator struct {
		ID              int    `json:"ID"`
		UserFirstname   string `json:"user_firstname"`
		UserLastname    string `json:"user_lastname"`
		Nickname        string `json:"nickname"`
		UserNicename    string `json:"user_nicename"`
		DisplayName     string `json:"display_name"`
		UserEmail       string `json:"user_email"`
		UserURL         string `json:"user_url"`
		UserRegistered  string `json:"user_registered"`
		UserDescription string `json:"user_description"`
		UserAvatar      string `json:"user_avatar"`
	}

	VlogAttributes struct {
		Youtube       string   `json:"youtube"`
		Series        string   `json:"series"`
		FilmEdit      Creators `json:"film_edit"`
		MovieLocation []int    `json:"movie_location"`
		YearMonth     string   `json:"year_month"`
		RunTime       string   `json:"run_time"`
		MovieTimeline string   `json:"movie_timeline"`
	}

	Creators []*Creator
)

func (cs *Creators) UnmarshalJSON(body []byte) error {
	if bytes.Equal(body, falseJSONBytes) {
		return nil
	}

	type Alias Creators
	return errors.WithStack(json.Unmarshal(body, (*Alias)(cs)))
}
