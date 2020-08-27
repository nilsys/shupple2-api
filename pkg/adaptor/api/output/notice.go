package output

import (
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

type (
	Notice struct {
		User           *UserSummary                  `json:"user"`
		ActionType     *model.NoticeActionType       `json:"actionType"`
		ActionTarget   *model.NoticeActionTargetType `json:"actionTarget"`
		ActionTargetID int                           `json:"actionTargetId"`
		IsRead         bool                          `json:"isRead"`
		CreatedAt      model.TimeResponse            `json:"createdAt"`
		Endpoint       string                        `json:"endpoint"`
	}

	NoticeList struct {
		Notices     []*Notice `json:"notices"`
		UnreadCount int       `json:"unreadCount"`
	}
)

func NewNotice(user *UserSummary, actionType *model.NoticeActionType, actionTarget *model.NoticeActionTargetType, actionTargetID int, isRead bool, createdAt time.Time, endpoint string) *Notice {
	return &Notice{
		User:           user,
		ActionType:     actionType,
		ActionTarget:   actionTarget,
		ActionTargetID: actionTargetID,
		IsRead:         isRead,
		CreatedAt:      model.TimeResponse(createdAt),
		Endpoint:       endpoint,
	}
}
