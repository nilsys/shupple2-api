package query

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	ListCfReturnGiftQuery struct {
		ProjectID     int
		UserID        int
		SessionStatus model.SessionStatus
		FindListPaginationQuery
	}
)
