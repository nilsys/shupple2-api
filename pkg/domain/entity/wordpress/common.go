package wordpress

import (
	"net/url"
	"strconv"
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

const (
	timeJSONFormat = `2006-01-02T15:04:05`
	dateJSONFormat = `2006-01-02`
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

	IntString int

	JSTTime time.Time
	UTCTime time.Time
	JSTDate time.Time
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

func (t *JSTDate) UnmarshalText(data []byte) error {
	parsed, err := time.ParseInLocation(dateJSONFormat, string(data), util.JSTLoc)
	*t = JSTDate(parsed)
	return err
}
