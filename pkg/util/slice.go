package util

import (
	"strconv"
	"strings"
)

func JoinIntSlice(val []int, sep string) string {
	var builder strings.Builder
	for i, v := range val {
		if i > 0 {
			_, _ = builder.WriteString(sep)
		}
		builder.WriteString(strconv.Itoa(v))
	}

	return builder.String()
}
