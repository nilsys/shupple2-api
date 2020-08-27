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
		Endpoint         string
		Times
	}

	NoticeList struct {
		List []*Notice
	}
)

func NewNotice(userID int, triggeredUserID int, actionType model.NoticeActionType, actionTargetType model.NoticeActionTargetType, actionTargetID int, endpoint string) *Notice {
	return &Notice{
		UserID:           userID,
		TriggeredUserID:  triggeredUserID,
		ActionType:       actionType,
		ActionTargetType: actionTargetType,
		ActionTargetID:   actionTargetID,
		IsRead:           false,
		Endpoint:         endpoint,
	}
}

func (notice Notice) IsOwnNotice() bool {
	return notice.UserID == notice.TriggeredUserID
}

func (n *NoticeList) UnreadIDs() []int {
	ids := make([]int, 0, len(n.List))

	for _, notice := range n.List {
		if !notice.IsRead {
			ids = append(ids, notice.ID)
		}
	}

	return ids
}
