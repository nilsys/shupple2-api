package model

import "time"

const timeFrontFmt = "2006-01-02T15:04+09:00"

// time.Time型からフロント用のformatに変換した文字列を返す
func TimeFmtToFrontStr(time time.Time) string {
	return time.Format(timeFrontFmt)
}
