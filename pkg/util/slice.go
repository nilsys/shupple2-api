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

func RemoveDuplicatesAndZeroFromIntSlice(val []int) []int {
	results := make([]int, 0, len(val))

	// 重複,0削除
	encountered := map[int]bool{}
	for i := 0; i < len(val); i++ {
		if !encountered[val[i]] && val[i] != 0 {
			encountered[val[i]] = true
			results = append(results, val[i])
		}
	}

	return results
}
