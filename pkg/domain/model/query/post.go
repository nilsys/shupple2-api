package query

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	// repositoryで使用する検索用構造体
	FindPostListQuery struct {
		UserID       int
		AreaID       int
		SubAreaID    int
		SubSubAreaID int
		ThemeID      int
		HashTag      string
		SortBy       model.SortBy
		Limit        int
		OffSet       int
	}
)
