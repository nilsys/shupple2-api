package model

import (
	"encoding/json"
	"time"
)

type Date time.Time

const dateUnmarshalFmt = `"2006-01-02"`
const dateMarshalFmt = "2006-01-02"

func (date *Date) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	err := date.parseData(data)
	return err
}

func (date Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(date).Format(dateMarshalFmt))
}

func (date *Date) parseData(data []byte) error {
	tm, err := time.Parse(dateUnmarshalFmt, string(data))
	if err != nil {
		return err
	}
	*date = Date(tm)
	return nil
}
