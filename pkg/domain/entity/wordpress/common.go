package wordpress

import (
	"net/url"
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

const (
	timeJSONFormat = `2006-01-02T15:04:05`
)

var (
	arrayJSONBytes = []byte("[]")
	falseJSONBytes = []byte("false")
)

type (
	Text struct {
		Rendered string `json:"rendered"`
	}

	ProtectableText struct {
		Rendered  string `json:"rendered"`
		Protected bool   `json:"protected"`
	}

	Self struct {
		Href string `json:"href"`
	}

	Collection struct {
		Href string `json:"href"`
	}

	URLEscapedString string

	JSTTime time.Time
	UTCTime time.Time
)

func (s *URLEscapedString) UnmarshalText(data []byte) error {
	unespcaed, err := url.QueryUnescape(string(data))
	if err != nil {
		return err
	}
	*s = URLEscapedString(unespcaed)
	return nil
}

func (t *JSTTime) UnmarshalText(data []byte) error {
	parsed, err := time.ParseInLocation(timeJSONFormat, string(data), util.JSTLoc)
	*t = JSTTime(parsed)
	return err
}

func (t *UTCTime) UnmarshalText(data []byte) error {
	parsed, err := time.ParseInLocation(timeJSONFormat, string(data), time.UTC)
	*t = UTCTime(parsed)
	return err
}
