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

func NewArrangeScheduleRequest(userID, matchingUserID int, date time.Time, remark string) *ArrangeScheduleRequest {
	return &ArrangeScheduleRequest{
		UserID:         userID,
		MatchingUserID: matchingUserID,
		Date:           date,
		Remark:         remark,
	}
}
