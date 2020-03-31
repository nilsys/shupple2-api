package query

import (
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

type (
	FindUserRankingListQuery struct {
		AreaID       int
		SubAreaID    int
		SubSubAreaID int
		SortBy       model.UserSortBy
		FromDate     time.Time
		ToDate       time.Time
		Limit        int
		Offset       int
	}

	FindFollowUser struct {
		ID     int
		Limit  int
		Offset int
	}
)
