package entity

import "time"

type (
	ArrangeScheduleRequestTiny struct {
		ID                  int `gorm:"primary_key"`
		UserID              int
		MatchingUserID      int
		DateTime            time.Time
		Remark              string
		MatchingUserApprove bool
		Times
	}

	ArrangeScheduleRequest struct {
		ArrangeScheduleRequestTiny
		User         *User `gorm:"foreignkey:ID;association_foreignkey:UserID"`
		MatchingUser *User `gorm:"foreignkey:ID;association_foreignkey:MatchingUserID"`
	}
)

func NewArrangeScheduleRequest(userID, matchingUserID int, date time.Time, remark string) *ArrangeScheduleRequestTiny {
	return &ArrangeScheduleRequestTiny{
		UserID:         userID,
		MatchingUserID: matchingUserID,
		DateTime:       date,
		Remark:         remark,
	}
}
