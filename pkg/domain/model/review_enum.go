package model

// queryのORDERで使用する値を返す
// デフォルトは新着順を表すNEW
func (sortBy ReviewSortBy) GetReviewOrderQuery() string {
	switch sortBy {
	case ReviewSortByNEW:
		return "created_at DESC"
	case ReviewSortByRECOMMEND:
		return "weekly_views DESC"
	default:
		return "created_at DESC"
	}
}

// Joinしている時にテーブル指定をする
func (sortBy ReviewSortBy) GetReviewOrderQueryForJoin() string {
	switch sortBy {
	case ReviewSortByNEW:
		return "review.created_at DESC"
	case ReviewSortByRECOMMEND:
		return "review.weekly_views DESC"
	default:
		return "review.created_at DESC"
	}
}
