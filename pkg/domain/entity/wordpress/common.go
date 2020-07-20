package wordpress

import (
	"net/url"
	"strconv"
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/util"
	"gopkg.in/guregu/null.v3"
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

	RelatedPost struct {
		ID int `json:"ID"`
	}

	URLEscapedString string

	IntString int

	JSTTime time.Time
	UTCTime time.Time

	NullableJSTTime null.Time
)

func (s *URLEscapedString) UnmarshalText(data []byte) error {
	unespcaed, err := url.QueryUnescape(string(data))
	if err != nil {
		return err
	}
	*s = URLEscapedString(unespcaed)
	return nil
}

func (i *IntString) UnmarshalText(data []byte) error {
	result, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}
	*i = IntString(result)
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

func (t *NullableJSTTime) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		*t = NullableJSTTime(null.Time{})
	}

	parsed, err := time.ParseInLocation(timeJSONFormat, string(data), util.JSTLoc)
	*t = NullableJSTTime(null.TimeFrom(parsed))
	return err
}
