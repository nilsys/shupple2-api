package input

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	ListCfReturnGift struct {
		ProjectID     int                 `query:"projectId"`
		UserID        int                 `query:"userId"`
		SessionStatus model.SessionStatus `query:"sessionStatus"`
		PaginationQuery
	}
)

const (
	listCfReturnGiftDefaultPerPage = 20
)

func (i *PaginationQuery) GetListCfReturnGiftLimit() int {
	if i.PerPage == 0 {
		return listCfReturnGiftDefaultPerPage
	}
	return i.PerPage
}

func (i *PaginationQuery) GetListCfReturnGiftOffset() int {
	if i.Page == 1 || i.Page == 0 {
		return 0
	}
	return i.GetPostLimit() * (i.Page - 1)
}
