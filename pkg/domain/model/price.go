package model

import (
	"strconv"
	"strings"
)

// 1000 => "1,000"
func PriceFormatCnv(price int) string {
	return strings.Replace(strconv.Itoa(price), "000", ",000", -1)
}
