package command

import (
	"time"

	"gopkg.in/guregu/null.v3"
)

type (
	StoreArrangeScheduleRequest struct {
		MatchingUserID int
		Date           time.Time
		Remark         string
		StartNow       null.Bool
	}
)
