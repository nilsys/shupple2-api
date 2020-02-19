package repository

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

type (
	// Review参照系レポジトリ
	ReviewQueryRepository interface {
		ShowReviewListByParams(query *query.ShowReviewListQuery) ([]*entity.Review, error)
		FindFeedReviewListByUserID(userID int, query *query.FindListPaginationQuery) ([]*entity.Review, error)
	}
)
