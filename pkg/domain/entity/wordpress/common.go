package wordpress

import (
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

const (
	timeJSONFormat = `"2006-01-02T15:04:05"`
)

var (
	arrayJSONBytes = []byte("[]")
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

	Time time.Time
)

func (t *Time) UnmarshalJSON(data []byte) error {
	parsed, err := time.ParseInLocation(timeJSONFormat, string(data), util.JSTLoc)
	*t = Time(parsed)
	return err
}
