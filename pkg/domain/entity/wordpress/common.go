package wordpress

import "time"

const (
	timeJSONFormat = `"2006-01-02T15:04:05"`
	locationString = "Asia/Tokyo"
)

var location *time.Location

func init() {
	l, err := time.LoadLocation(locationString)
	if err != nil {
		panic(err)
	}
	location = l
}

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
	parsed, err := time.ParseInLocation(timeJSONFormat, string(data), location)
	*t = Time(parsed)
	return err
}
