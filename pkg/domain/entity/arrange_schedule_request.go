package entity

import "time"

type (
	ArrangeScheduleRequest struct {
		UserID              int `gorm:"primary_key"`
		MatchingUserID      int `gorm:"primary_key"`
		Date                time.Time
		Remark              string
		MatchingUserApprove bool
		Times
	}
)

func NewArrangeScheduleRequest() ArrangeScheduleRequest {

}
