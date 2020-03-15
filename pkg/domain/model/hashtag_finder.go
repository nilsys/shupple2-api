package model

import "regexp"

var re = regexp.MustCompile(`#\S+`)

// 文字列の中からhashtagを探す
func FindHashtags(str string) []string {
	return re.FindAllString(str, -1)
}
