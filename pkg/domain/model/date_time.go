package model

import (
	"encoding/json"
	"time"

	"github.com/uma-co82/shupple2-api/pkg/util"
)

type DateTime time.Time

const (
	dateTimeFmt = "2006-01-02 15:04"
)

func (d *DateTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	tm, err := time.Parse(dateTimeFmt, str)
	if err != nil {
		return err
	}
	*d = DateTime(tm.In(time.UTC))
	return nil
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d).In(util.JSTLoc).Format(dateTimeFmt))
}
