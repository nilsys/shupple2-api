package query

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

type (
	// repositoryで使用するクエリ発行用構造体
	ShowReviewListQuery struct {
		UserID                 int
		InnID                  int
		TouristSpotID          int
		HashTag                string
		AreaID                 int
		SubAreaID              int
		SubSubAreaID           int
		MetasearchAreaID       int
		MetasearchSubAreaID    int
		MetasearchSubSubAreaID int
		SortBy                 model.ReviewSortBy
		Keyword                string
		ExcludeID              int
		Limit                  int
		OffSet                 int
		// InnQueryRepositoryから取得後にいれる
		InnIDs []int
	}
)

func (q *ShowReviewListQuery) SQLLikeKeyword() string {
	return "%" + q.Keyword + "%"
}
