package input

import (
	"github.com/uma-co82/shupple2-api/pkg/domain/model"
	"gopkg.in/guregu/null.v3"
)

type (
	StoreArrangeScheduleRequest struct {
		MatchingUserID IDParam
		DateTime       model.DateTime `json:"dateTime"`
		Remark         string         `json:"remark"`
		StartNow       null.Bool      `json:"startNow"`
	}
)
