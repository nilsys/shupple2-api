package model

import (
	"encoding/json"
	"time"
)

type YearMonth struct {
	time.Time
}

const yearMonthFmt = `"2006-01"`
const yearMonthStrFmt = "2006-01"

func (yearMonth *YearMonth) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	err := yearMonth.parseDate(data)
	return err
}

func (yearMonth *YearMonth) MarshalJSON() ([]byte, error) {
	tm := yearMonth.Time
	return json.Marshal(tm.Format(yearMonthStrFmt))
}

func (yearMonth *YearMonth) parseDate(data []byte) error {
	tm, err := time.Parse(yearMonthFmt, string(data))
	if err != nil {
		return err
	}
	yearMonth.Time = tm
	return nil
}

func NewYearMonth(time time.Time) YearMonth {
	return YearMonth{time}
}
