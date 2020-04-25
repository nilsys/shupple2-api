package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

func ConvertListNoticeToOutput(notices []*entity.Notice) *output.NoticeList {
	responses := make([]*output.Notice, len(notices))
	unReadCount := 0
	for i, notice := range notices {
		responses[i] = convertNoticeToOutput(notice)
		if !notice.IsRead {
			unReadCount++
		}
	}

	return &output.NoticeList{
		Notices:     responses,
		UnreadCount: unReadCount,
	}
}

func convertNoticeToOutput(notice *entity.Notice) *output.Notice {
	user := output.NewUserSummary(notice.TriggeredUser.ID, notice.TriggeredUser.UID, notice.TriggeredUser.Name, notice.TriggeredUser.IconURL())
	return output.NewNotice(user, &notice.ActionType, &notice.ActionTargetType, notice.ActionTargetID, notice.IsRead)
}
