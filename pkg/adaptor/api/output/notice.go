package output

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	Notice struct {
		User           *UserSummary                  `json:"user"`
		ActionType     *model.NoticeActionType       `json:"actionType"`
		ActionTarget   *model.NoticeActionTargetType `json:"actionTarget"`
		ActionTargetID int                           `json:"actionTargetId"`
		IsRead         bool                          `json:"isRead"`
	}

	NoticeList struct {
		Notices     []*Notice `json:"notices"`
		UnreadCount int       `json:"unreadCount"`
	}
)

func NewNotice(user *UserSummary, actionType *model.NoticeActionType, actionTarget *model.NoticeActionTargetType, actionTargetID int, isRead bool) *Notice {
	return &Notice{
		User:           user,
		ActionType:     actionType,
		ActionTarget:   actionTarget,
		ActionTargetID: actionTargetID,
		IsRead:         isRead,
	}
}
