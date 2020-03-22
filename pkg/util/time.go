package util

import "time"

const (
	locationStringForTokyo = "Asia/Tokyo"
)

var JSTLoc *time.Location

func init() {
	l, err := time.LoadLocation(locationStringForTokyo)
	if err != nil {
		panic(err)
	}
	JSTLoc = l
}
