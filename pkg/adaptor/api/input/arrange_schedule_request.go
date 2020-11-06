package input

import "github.com/uma-co82/shupple2-api/pkg/domain/model"

type (
	StoreArrangeScheduleRequest struct {
		MatchingUserID IDParam
		Date           model.Date
	}
)
