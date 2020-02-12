package repository

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/query"
)

type (
	// Review参照系レポジトリ
	ReviewQueryRepository interface {
		ShowReviewListByParams(query *query.ShowReviewListQuery) ([]*entity.Review, error)
	}
)
