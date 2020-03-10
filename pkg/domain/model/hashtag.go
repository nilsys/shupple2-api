package model

import "regexp"

var re = regexp.MustCompile(`#\S+`)

func FindHashtags(str string) []string {
	return re.FindAllString(str, -1)
}
