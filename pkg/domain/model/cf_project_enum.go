package model

func (s CfProjectSortBy) GetCfProjectOrderQuery() string {
	switch s {
	case CfProjectSortByLargeAmount:
		return "achieved_price desc"
	default:
		return "created_at desc"
	}
}
