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
		InnTypeID              int
		InnDiscerningType      int
		ThemeID                int
		HashTag                string
		SortBy                 model.MediaSortBy
		Keyward                string
		Limit                  int
		OffSet                 int
	}
)
