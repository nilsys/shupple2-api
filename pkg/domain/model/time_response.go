package model

import (
	"encoding/json"
	"time"
)

type TimeResponse time.Time

const dateTimeFmt = `"2006-01-02T15:04+09:00"`
const dateTimeStrFmt = "2006-01-02T15:04+09:00"

func (tr TimeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(tr).Format(dateTimeStrFmt))
}

func (tr TimeResponse) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	err := tr.parseData(data)
	return err
}

func (tr *TimeResponse) parseData(data []byte) error {
	tm, err := time.Parse(dateTimeFmt, string(data))
	if err != nil {
		return err
	}
	*tr = TimeResponse(tm)
	return nil
}
