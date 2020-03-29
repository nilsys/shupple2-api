package model

import (
	"regexp"
	"strings"

	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

var re = regexp.MustCompile(`#(\S+)`)

// 文字列の中からhashtagを探す
func FindHashtags(str string) []string {
	str = strings.ReplaceAll(str, "#", " #")
	hashtags := make([]string, 0)
	matches := re.FindAllStringSubmatch(str, -1)
	for _, s := range matches {
		hashtags = append(hashtags, strings.TrimSpace(s[1]))
	}
	return util.RemoveDuplicatesFromStringSlice(hashtags)
}
