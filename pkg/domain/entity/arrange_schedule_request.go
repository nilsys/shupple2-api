package entity

import "time"

type (
	ArrangeScheduleRequest struct {
		ID                  int `gorm:"primary_key"`
		UserID              int
		MatchingUserID      int
		Date                time.Time
		Remark              string
		MatchingUserApprove bool
		Times
	}
)
