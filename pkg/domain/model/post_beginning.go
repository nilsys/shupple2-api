package model

import (
	"strings"
	"unicode/utf8"

	"github.com/pkg/errors"

	"github.com/PuerkitoBio/goquery"
)

const (
	postBeginningCnt = 200
)

// PostのBody(HTML)をPost.Beginningへ変換
func PostBodyToBeginning(body string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return "", errors.Wrap(err, "failed doc from reader")
	}
	resolve := doc.Text()

	if utf8.RuneCountInString(resolve) < postBeginningCnt {
		return resolve, nil
	}

	return string([]rune(resolve)[:postBeginningCnt]), nil
}
