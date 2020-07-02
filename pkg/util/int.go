package util

import "github.com/dustin/go-humanize"

// 1000 => "1,000"
func WithComma(price int) string {
	return humanize.Comma(int64(price))
}
