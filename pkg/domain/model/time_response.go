package model

import (
	"encoding/json"
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

type TimeResponse time.Time

func (tr *TimeResponse) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	tm, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return err
	}
	*tr = TimeResponse(tm.In(time.UTC))
	return nil
}

func (tr TimeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(tr).In(util.JSTLoc).Format(time.RFC3339))
}
