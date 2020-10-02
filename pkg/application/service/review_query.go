package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/logger"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
	"go.uber.org/zap"
)

type (
	// Review参照系サービス
	ReviewQueryService interface {
		ListByParams(query *query.ShowReviewListQuery, ouser entity.OptionalUser) (*entity.ReviewDetailWithIsFavoriteList, error)
		ListFeed(user entity.User, query *query.FindListPaginationQuery) (*entity.ReviewDetailWithIsFavoriteList, error)
		ShowQueryReview(id int, ouser entity.OptionalUser) (*entity.ReviewDetailWithIsFavorite, error)
		ShowReview(id int) (*entity.Review, error)
		ListReviewCommentByReviewID(reviewID int, limit int, ouser entity.OptionalUser) ([]*entity.ReviewCommentWithIsFavorite, error)
		ListFavoriteReview(ouser entity.OptionalUser, userID int, query *query.FindListPaginationQuery) (*entity.ReviewDetailWithIsFavoriteList, error)
		ListReviewCommentReplyByReviewCommentID(reviewCommentID int, ouser *entity.OptionalUser) ([]*entity.ReviewCommentReplyWithIsFavorite, error)
	}

	// Review参照系サービス実装
	ReviewQueryServiceImpl struct {
		repository.ReviewQueryRepository
		repository.InnQueryRepository
		repository.MetasearchAreaQueryRepository
	}
)

var ReviewQueryServiceSet = wire.NewSet(
	wire.Struct(new(ReviewQueryServiceImpl), "*"),
	wire.Bind(new(ReviewQueryService), new(*ReviewQueryServiceImpl)),
)

// TODO: リファクタ
// クエリで飛んで来た検索条件を用いreviewを検索
func (s *ReviewQueryServiceImpl) ListByParams(q *query.ShowReviewListQuery, ouser entity.OptionalUser) (*entity.ReviewDetailWithIsFavoriteList, error) {
	if q.AreaID != 0 || q.SubAreaID != 0 || q.SubSubAreaID != 0 {
		findInnQuery := &query.FindInn{}

		if q.AreaID != 0 {
			metasearchAreas, err := s.MetasearchAreaQueryRepository.FindByAreaCategoryID(q.AreaID, model.AreaCategoryTypeArea)
			if err != nil {
				return nil, errors.Wrap(err, "failed to get area")
			}
			findInnQuery.SetMetaserachID(metasearchAreas)
		}
		if q.SubAreaID != 0 {
			metasearchAreas, err := s.MetasearchAreaQueryRepository.FindByAreaCategoryID(q.SubAreaID, model.AreaCategoryTypeSubArea)
			if err != nil {
				return nil, errors.Wrap(err, "failed to get sub area")
			}
			findInnQuery.SetMetaserachID(metasearchAreas)
		}
		if q.SubSubAreaID != 0 {
			metasearchAreas, err := s.MetasearchAreaQueryRepository.FindByAreaCategoryID(q.SubSubAreaID, model.AreaCategoryTypeSubSubArea)
			if err != nil {
				return nil, errors.Wrap(err, "failed to get sub sub area")
			}
			findInnQuery.SetMetaserachID(metasearchAreas)
		}

		innIDs, err := s.InnQueryRepository.FindByParams(findInnQuery)
		if err != nil {
			// errorは握り潰す
			logger.Error("failed metasearch inns api", zap.Error(err))
		}
		// stayway-apiから取得したinn_idを検索に用いる
		q.InnIDs = innIDs.IDs()
	}

	// 認証されているユーザーの場合、お気に入りしているかのフラグも取得する
	if ouser.Authenticated {
		reviews, err := s.ReviewQueryRepository.ShowReviewWithIsFavoriteListByParams(q, ouser.ID)
		if err != nil {
			return nil, errors.Wrap(err, "failed to show review list from repo")
		}
		return reviews, nil
	}

	reviews, err := s.ReviewQueryRepository.ShowReviewListByParams(q)
	if err != nil {
		return nil, errors.Wrap(err, "failed to show review list from repo")
	}

	return reviews, nil
}

func (s *ReviewQueryServiceImpl) ListFeed(user entity.User, query *query.FindListPaginationQuery) (*entity.ReviewDetailWithIsFavoriteList, error) {
	return s.ReviewQueryRepository.FindFeedReviewWithIsFavoriteListByUserID(user.ID, query)
}

func (s *ReviewQueryServiceImpl) ShowQueryReview(id int, ouser entity.OptionalUser) (*entity.ReviewDetailWithIsFavorite, error) {
	if ouser.Authenticated {
		return s.ReviewQueryRepository.FindQueryReviewWithIsFavoriteByID(id, ouser.ID)
	}
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

func (s *ReviewQueryServiceImpl) ListReviewCommentByReviewID(reviewID int, limit int, ouser entity.OptionalUser) ([]*entity.ReviewCommentWithIsFavorite, error) {
	if ouser.Authenticated {
		return s.ReviewQueryRepository.FindReviewCommentWithIsFavoriteListByReviewID(reviewID, limit, ouser.ID)
	}
	return s.ReviewQueryRepository.FindReviewCommentListByReviewID(reviewID, limit)
}

func (s *ReviewQueryServiceImpl) ListFavoriteReview(ouser entity.OptionalUser, userID int, query *query.FindListPaginationQuery) (*entity.ReviewDetailWithIsFavoriteList, error) {
	if ouser.Authenticated {
		return s.ReviewQueryRepository.FindFavoriteReviewWithIsFavoriteListByUserID(ouser.ID, userID, query)
	}
	return s.ReviewQueryRepository.FindFavoriteReviewListByUserID(userID, query)
}

func (s *ReviewQueryServiceImpl) ListReviewCommentReplyByReviewCommentID(reviewCommentID int, ouser *entity.OptionalUser) ([]*entity.ReviewCommentReplyWithIsFavorite, error) {
	if ouser.Authenticated {
		return s.ReviewQueryRepository.FindReviewCommentReplyWithIsFavoriteListByReviewCommentID(reviewCommentID, ouser.ID)
	}
	return s.ReviewQueryRepository.FindReviewCommentReplyListByReviewCommentID(reviewCommentID)
}
