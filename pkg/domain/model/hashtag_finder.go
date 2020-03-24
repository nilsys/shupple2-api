package model

import "regexp"

var re = regexp.MustCompile(`#(\S+)`)

// 文字列の中からhashtagを探す
func FindHashtags(str string) []string {
	hashtags := make([]string, 0)
	matches := re.FindAllStringSubmatch(str, -1)
	for _, s := range matches {
		hashtags = append(hashtags, s[1])
	}
	return hashtags
}
