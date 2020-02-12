package model

// queryのORDERで使用する値を返す
// デフォルトは新着順を表すNEW
func (sortBy SortBy) GetPostOrderQuery() string {
	switch sortBy {
	case SortByRANKING:
		return "favorite_count desc"
	case SortByNEW:
		return "created_at desc"
	default:
		return "created_at desc"
	}
}
