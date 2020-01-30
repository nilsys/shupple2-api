package query

type (
	// repositoryで使用する検索用構造体
	FindPostListQuery struct {
		AreaID       int
		SubAreaID    int
		SubSubAreaID int
		ThemeID      int
		HashTag      string
		SortBy       string
		Limit        int
		OffSet       int
	}
)
