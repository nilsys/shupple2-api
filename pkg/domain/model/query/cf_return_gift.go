package query

type (
	ListCfReturnGiftQuery struct {
		ProjectID int `query:"projectId"`
		UserID    int `query:"userId"`
		FindListPaginationQuery
	}
)
