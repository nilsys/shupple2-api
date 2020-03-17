package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

type (
	// Review参照系レポジトリ
	ReviewQueryRepository interface {
		ShowReviewListByParams(query *query.ShowReviewListQuery) ([]*entity.QueryReview, error)
		FindFeedReviewListByUserID(userID int, query *query.FindListPaginationQuery) ([]*entity.QueryReview, error)
		FindReviewCommentListByReviewID(reviewID int, limit int) ([]*entity.ReviewComment, error)
		FindQueryReviewByID(id int) (*entity.QueryReview, error)
	}

	// Reviewコマンド系レポジトリ
	ReviewCommandRepository interface {
		StoreReview(c context.Context, review *entity.Review) error
		PersistReviewMedia(reviewMedia *entity.ReviewMedia) error
	}
)
