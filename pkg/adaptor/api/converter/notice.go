package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

func (c Converters) ConvertListNoticeToOutput(notices []*entity.Notice) *output.NoticeList {
	responses := make([]*output.Notice, len(notices))
	unReadCount := 0
	for i, notice := range notices {
		responses[i] = c.convertNoticeToOutput(notice)
		if !notice.IsRead {
			unReadCount++
		}
	}

	return &output.NoticeList{
		Notices:     responses,
		UnreadCount: unReadCount,
	}
}

func (c Converters) convertNoticeToOutput(notice *entity.Notice) *output.Notice {
	return output.NewNotice(c.NewUserSummaryFromUser(notice.TriggeredUser), &notice.ActionType, &notice.ActionTargetType, notice.ActionTargetID, notice.IsRead)
}
