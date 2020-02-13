package query

// repositoryで使用する検索用構造体
type (
	// 一覧検索でlimit, offsetの指定のみある場合
	FindListPaginationQuery struct {
		Limit  int
		Offset int
	}
)
