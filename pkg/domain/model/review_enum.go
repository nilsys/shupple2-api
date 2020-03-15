package model

// queryのORDERで使用する値を返す
// デフォルトは新着順を表すNEW
func (sortBy ReviewSortBy) GetReviewOrderQuery() string {
	switch sortBy {
	case ReviewSortByNEW:
		return "updated_at DESC"
	case ReviewSortByRECOMMEND:
		return "views DESC"
	default:
		return "updated_at DESC"
	}
}
