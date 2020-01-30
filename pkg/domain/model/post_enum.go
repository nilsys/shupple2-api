package model

// 記事検索のソート順を指定する値オブジェクト
type SortBy int

const (
	RANKING SortBy = iota
	NEW
)

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
