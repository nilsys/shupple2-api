package command

import "time"

type (
	StoreArrangeScheduleRequest struct {
		MatchingUserID int
		Date           time.Time
		Remark         string
	}
)
