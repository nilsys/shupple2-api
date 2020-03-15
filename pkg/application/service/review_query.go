package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
	"go.uber.org/zap"
)

type (
	// Review参照系サービス
	ReviewQueryService interface {
		ShowReviewListByParams(query *query.ShowReviewListQuery) ([]*entity.QueryReview, error)
		ShowListFeed(userID int, query *query.FindListPaginationQuery) ([]*entity.QueryReview, error)
		ShowReview(id int) (*entity.QueryReview, error)
		ListReviewCommentByReviewID(reviewID int, limit int) ([]*entity.ReviewComment, error)
	}

	// Review参照系サービス実装
	ReviewQueryServiceImpl struct {
		ReviewQueryRepository repository.ReviewQueryRepository
		InnQueryRepository    repository.InnQueryRepository
	}
)

var ReviewQueryServiceSet = wire.NewSet(
	wire.Struct(new(ReviewQueryServiceImpl), "*"),
	wire.Bind(new(ReviewQueryService), new(*ReviewQueryServiceImpl)),
)

// クエリで飛んで来た検索条件を用いreviewを検索
func (s *ReviewQueryServiceImpl) ShowReviewListByParams(query *query.ShowReviewListQuery) ([]*entity.QueryReview, error) {
	innIDs, err := s.InnQueryRepository.FindIDsByAreaID(query.AreaID, query.SubAreaID, query.SubSubAreaID)
	if err != nil {
		zap.Error(err)
	}

	// stayway-apiから取得したinn_idを検索に用いる
	query.InnIDs = innIDs

	reviews, err := s.ReviewQueryRepository.ShowReviewListByParams(query)
	if err != nil {
		return nil, errors.Wrap(err, "failed to show review list from repo")
	}

	return reviews, nil
}

func (s *ReviewQueryServiceImpl) ShowListFeed(userID int, query *query.FindListPaginationQuery) ([]*entity.QueryReview, error) {
	return s.ReviewQueryRepository.FindFeedReviewListByUserID(userID, query)
}

func (s *ReviewQueryServiceImpl) ShowReview(id int) (*entity.QueryReview, error) {
	return s.ReviewQueryRepository.FindQueryReviewByID(id)
}

func (s *ReviewQueryServiceImpl) ListReviewCommentByReviewID(reviewID int, limit int) ([]*entity.ReviewComment, error) {
	return s.ReviewQueryRepository.FindReviewCommentListByReviewID(reviewID, limit)
}
