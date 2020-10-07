package facebook

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/huandu/facebook/v2"

	"github.com/stayway-corp/stayway-media-api/pkg/config"
)

type (
	EngagementAndID struct {
		Engagement Engagement `json:"engagement"`
		ID         string     `json:"id"`
	}

	Engagement struct {
		ReactionCount      int `json:"reaction_count"`
		CommentCount       int `json:"comment_count"`
		ShareCount         int `json:"share_count"`
		CommentPluginCount int `json:"comment_plugin_count"`
	}

	EngagementAndIDList []*EngagementAndID
)

const (
	paramsMethodKey                      = "method"
	paramsRelativeURLKey                 = "relative_url"
	paramsRelativeURLPrefix              = "?id="
	paramsRelativeURLSuffix              = "&fields=engagement"
	paramsRelativeTrailingSlashURLSuffix = "/&fields=engagement"
)

func (e EngagementAndIDList) GetShareCountBySuffixKey(slug string) int {
	resolve := 0
	for _, engagement := range e {
		if strings.HasSuffix(engagement.ID, slug) || strings.HasSuffix(engagement.ID, slug+"/") {
			resolve += engagement.Engagement.ShareCount
		}
	}
	return resolve
}

func GetRelativeURLParams(url *config.URL) facebook.Params {
	return facebook.Params{
		paramsMethodKey:      http.MethodGet,
		paramsRelativeURLKey: fmt.Sprintf("%s%s%s", paramsRelativeURLPrefix, url.String(), paramsRelativeURLSuffix),
	}
}

func GetRelativeTrailingSlashURLParams(url *config.URL) facebook.Params {
	return facebook.Params{
		paramsMethodKey:      http.MethodGet,
		paramsRelativeURLKey: fmt.Sprintf("%s%s%s", paramsRelativeURLPrefix, url.String(), paramsRelativeTrailingSlashURLSuffix),
	}
}
