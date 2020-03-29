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

func RemoveDuplicatesFromStringSlice(ss []string) []string {
	result := make([]string, 0, len(ss))
	existing := make(map[string]struct{})

	for _, s := range ss {
		if _, exist := existing[s]; !exist {
			result = append(result, s)
			existing[s] = struct{}{}
		}
	}

	return result
}
