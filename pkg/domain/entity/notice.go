package entity

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

type (
	Notice struct {
		ID               int
		UserID           int
		TriggeredUserID  int
		TriggeredUser    *User `gorm:"foreignkey:TriggeredUserID"`
		ActionType       model.NoticeActionType
		ActionTargetType model.NoticeActionTargetType
		ActionTargetID   int
		IsRead           bool
		Times
	}
)

func NewNotice(userID int, triggeredUserID int, actionType model.NoticeActionType, actionTargetType model.NoticeActionTargetType, actionTargetID int) *Notice {
	return &Notice{
		UserID:           userID,
		TriggeredUserID:  triggeredUserID,
		ActionType:       actionType,
		ActionTargetType: actionTargetType,
		ActionTargetID:   actionTargetID,
		IsRead:           false,
	}
}

func (notice Notice) IsOwnNotice() bool {
	return notice.UserID == notice.TriggeredUserID
}
