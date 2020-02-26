package model

import (
	"time"

	"github.com/pkg/errors"
)

const timeFrontFmt = "2006-01-02T15:04+09:00"
const timeFromFrontLayout = "2006-01-02"

// time.Time型からフロント用のformatに変換した文字列を返す
func TimeFmtToFrontStr(time time.Time) string {
	return time.Format(timeFrontFmt)
}

// フロントから日付指定で来たformatでtime.Time型へパース
func ParseTimeFromFrontStr(timeStr string) (time.Time, error) {
	fromStrTime, err := time.Parse(timeFromFrontLayout, timeStr)
	if err != nil {
		return time.Now(), errors.Wrap(err, "failed parse invalid arg")
	}

	return fromStrTime, nil
}
