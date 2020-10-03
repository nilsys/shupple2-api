package query

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	// repositoryで使用する検索用構造体
	FindPostListQuery struct {
		UserID                 int
		AreaID                 int
		SubAreaID              int
		SubSubAreaID           int
		ChildAreaID            int
		ChildSubAreaID         int
		ChildSubSubAreaID      int
		MetasearchAreaID       int
		MetasearchSubAreaID    int
		MetasearchSubSubAreaID int
		MetasearchInnTypeID    int
		MetasearchTagID        int
		ThemeID                int
		SubThemeID             int
		ThemeSlug              string
		HashTag                string
		SortBy                 model.MediaSortBy
		Keyword                string
		CfProjectID            int
		NoHaveAreaID           bool
		Limit                  int
		OffSet                 int
	}
)

func (q *FindPostListQuery) SQLLikeKeyword() string {
	return "%" + q.Keyword + "%"
}
