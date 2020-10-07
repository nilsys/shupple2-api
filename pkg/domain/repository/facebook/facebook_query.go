package facebook

import (
	"github.com/huandu/facebook/v2"
	facebook2 "github.com/stayway-corp/stayway-media-api/pkg/domain/entity/facebook"
)

type (
	QueryRepository interface {
		GetShareCountByURLBatchRequest(query []facebook.Params) (facebook2.EngagementAndIDList, error)
	}
)
