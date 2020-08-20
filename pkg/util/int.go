package util

import (
	"github.com/dustin/go-humanize"
	"gopkg.in/guregu/null.v3"
)

// 1000 => "1,000"
func WithComma(price int) string {
	return humanize.Comma(int64(price))
}

func NullIntFrom(i int) null.Int {
	return null.IntFrom(int64(i))
}
