package model

import (
	"regexp"
)

var taggedRegex = regexp.MustCompile(`@(\S+)`)

// 文字列の中からタグ付されたユーザを探す
func FindTaggedUser(str string) []string {
	var taggedUsers []string

	for _, matchStrings := range taggedRegex.FindAllStringSubmatch(str, -1) {
		taggedUsers = append(taggedUsers, matchStrings[1])
	}
	return taggedUsers
}
