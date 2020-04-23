package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/logger"
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
		ShowQueryReview(id int) (*entity.QueryReview, error)
		ShowReview(id int) (*entity.Review, error)
		ListReviewCommentByReviewID(reviewID int, limit int) ([]*entity.ReviewComment, error)
		ListFavoriteReview(userID int, query *query.FindListPaginationQuery) ([]*entity.QueryReview, error)
		ListReviewCommentReplyByReviewCommentID(reviewCommentID int) ([]*entity.ReviewCommentReply, error)
	}

	// Review参照系サービス実装
	ReviewQueryServiceImpl struct {
		ReviewQueryRepository       repository.ReviewQueryRepository
		InnQueryRepository          repository.InnQueryRepository
		AreaCategoryQueryRepository repository.AreaCategoryQueryRepository
	}
)

var ReviewQueryServiceSet = wire.NewSet(
	wire.Struct(new(ReviewQueryServiceImpl), "*"),
	wire.Bind(new(ReviewQueryService), new(*ReviewQueryServiceImpl)),
)

// TODO: リファクタ
// クエリで飛んで来た検索条件を用いreviewを検索
func (s *ReviewQueryServiceImpl) ShowReviewListByParams(query *query.ShowReviewListQuery) ([]*entity.QueryReview, error) {
	var metasearchAreaID int
	var metasearchSubAreaID int
	var metasearchSubSubAreaID int
	// metasearch側のidで検索する為(metasearch側と乖離している為typeは調べない)
	if query.AreaID != 0 {
		area, err := s.AreaCategoryQueryRepository.FindByID(query.AreaID)
		if err != nil {
			return nil, errors.Wrap(err, "failed to find area_category")
		}
		metasearchAreaID = area.MetasearchAreaID
		metasearchSubAreaID = area.MetasearchAreaID
		metasearchSubSubAreaID = area.MetasearchAreaID
	}
	if query.SubAreaID != 0 {
		subArea, err := s.AreaCategoryQueryRepository.FindByID(query.SubAreaID)
		if err != nil {
			return nil, errors.Wrap(err, "failed to find area_category")
		}
		metasearchAreaID = subArea.MetasearchSubAreaID
		metasearchSubAreaID = subArea.MetasearchSubAreaID
		metasearchSubSubAreaID = subArea.MetasearchSubAreaID
	}
	if query.SubSubAreaID != 0 {
		subSubArea, err := s.AreaCategoryQueryRepository.FindByID(query.SubAreaID)
		if err != nil {
			return nil, errors.Wrap(err, "failed to find area_category")
		}
		metasearchAreaID = subSubArea.MetasearchSubSubAreaID
		metasearchSubAreaID = subSubArea.MetasearchSubSubAreaID
		metasearchSubSubAreaID = subSubArea.MetasearchSubSubAreaID
	}

	if metasearchAreaID != 0 || metasearchSubAreaID != 0 || metasearchSubSubAreaID != 0 {
		// 指定されたareaに紐づいているinnのidを取得
		innIDs, err := s.InnQueryRepository.FindIDsByAreaID(metasearchAreaID, metasearchSubAreaID, metasearchSubSubAreaID)
		if err != nil {
			// errorは握り潰す
			logger.Error("failed metasearch inns api", zap.Error(err))
		}
		// stayway-apiから取得したinn_idを検索に用いる
		query.InnIDs = innIDs
	}

	reviews, err := s.ReviewQueryRepository.ShowReviewListByParams(query)
	if err != nil {
		return nil, errors.Wrap(err, "failed to show review list from repo")
	}

	return reviews, nil
}

func (s *ReviewQueryServiceImpl) ShowListFeed(userID int, query *query.FindListPaginationQuery) ([]*entity.QueryReview, error) {
	return s.ReviewQueryRepository.FindFeedReviewListByUserID(userID, query)
}

func (s *ReviewQueryServiceImpl) ShowQueryReview(id int) (*entity.QueryReview, error) {
	return s.ReviewQueryRepository.FindQueryReviewByID(id)
}

func (s *ReviewQueryServiceImpl) ShowByID(reviewID int) (*entity.Review, error) {
	review, err := s.ReviewQueryRepository.FindByID(reviewID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed FindByID")
	}

	return review, nil
}

func (s *ReviewQueryServiceImpl) ShowReview(id int) (*entity.Review, error) {
	return s.ReviewQueryRepository.FindByID(id)
}

func (s *ReviewQueryServiceImpl) ListReviewCommentByReviewID(reviewID int, limit int) ([]*entity.ReviewComment, error) {
	return s.ReviewQueryRepository.FindReviewCommentListByReviewID(reviewID, limit)
}

func (s *ReviewQueryServiceImpl) ListFavoriteReview(userID int, query *query.FindListPaginationQuery) ([]*entity.QueryReview, error) {
	return s.ReviewQueryRepository.FindFavoriteListByUserID(userID, query)
}

func (s *ReviewQueryServiceImpl) ListReviewCommentReplyByReviewCommentID(reviewCommentID int) ([]*entity.ReviewCommentReply, error) {
	return s.ReviewQueryRepository.FindReviewCommentReplyListByReviewCommentID(reviewCommentID)
}
