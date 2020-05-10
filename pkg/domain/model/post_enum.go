package model

// queryのORDERで使用する値を返す
// デフォルトは新着順を表すNEW
func (sortBy MediaSortBy) GetPostOrderQuery() string {
	switch sortBy {
	case MediaSortByRANKING:
		return "favorite_count desc"
	case MediaSortByNEW:
		return "created_at desc"
	case MediaSortByRECOMMEND:
		return "weekly_views desc"
	default:
		return "created_at desc"
	}
}
