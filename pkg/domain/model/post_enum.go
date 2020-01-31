package model

// 記事検索のソート順を指定する値オブジェクト
type SortBy int

const (
	RANKING SortBy = iota
	NEW
)

// stringからSortByを返す
// デフォルトは新着順を表すNEW
func NewSortBy(sortByStr string) SortBy {
	switch sortByStr {
	case RANKING.Value():
		return RANKING
	case NEW.Value():
		return NEW
	default:
		return NEW
	}
}

// queryのORDERで使用する値を返す
// デフォルトは新着順を表すNEW
func (sortBy SortBy) GetOrderQuery() string {
	switch sortBy {
	case RANKING:
		return "favorite_count desc"
	case NEW:
		return "created_at desc"
	default:
		return "created_at desc"
	}
}

// 文字列で扱う
func (sortBy SortBy) Value() string {
	switch sortBy {
	case RANKING:
		return "RANKING"
	case NEW:
		return "NEW"
	default:
		return ""
	}
}
