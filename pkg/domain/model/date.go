package model

import (
	"encoding/json"
	"time"

	"github.com/uma-co82/shupple2-api/pkg/util"
)

type Date time.Time

const dateFormat = "2006-01-02"

func (date *Date) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	tm, err := time.Parse(dateFormat, str)
	if err != nil {
		return err
	}
	*date = Date(tm.In(time.UTC))
	return nil
}

func (date *Date) UnmarshalText(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	tm, err := time.Parse(dateFormat, string(data))
	if err != nil {
		return err
	}
	*date = Date(tm.In(time.UTC))
	return nil
}

func (date Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(date).In(util.JSTLoc).Format(dateFormat))
}
