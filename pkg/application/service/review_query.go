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
		ShowReviewListByParams(query *query.ShowReviewListQuery, ouser entity.OptionalUser) (*entity.ReviewDetailWithIsFavoriteList, error)
		ListFeed(ouser entity.OptionalUser, userID int, query *query.FindListPaginationQuery) (*entity.ReviewDetailWithIsFavoriteList, error)
		ShowQueryReview(id int, ouser entity.OptionalUser) (*entity.ReviewDetailWithIsFavorite, error)
		ShowReview(id int) (*entity.Review, error)
		ListReviewCommentByReviewID(reviewID int, limit int, ouser entity.OptionalUser) ([]*entity.ReviewCommentWithIsFavorite, error)
		ListFavoriteReview(ouser entity.OptionalUser, userID int, query *query.FindListPaginationQuery) (*entity.ReviewDetailWithIsFavoriteList, error)
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
func (s *ReviewQueryServiceImpl) ShowReviewListByParams(query *query.ShowReviewListQuery, ouser entity.OptionalUser) (*entity.ReviewDetailWithIsFavoriteList, error) {
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
		metasearchSubAreaID = area.MetasearchSubAreaID
		metasearchSubSubAreaID = area.MetasearchSubSubAreaID
	}
	if query.SubAreaID != 0 {
		subArea, err := s.AreaCategoryQueryRepository.FindByID(query.SubAreaID)
		if err != nil {
			return nil, errors.Wrap(err, "failed to find area_category")
		}
		metasearchAreaID = subArea.MetasearchAreaID
		metasearchSubAreaID = subArea.MetasearchSubAreaID
		metasearchSubSubAreaID = subArea.MetasearchSubSubAreaID
	}
	if query.SubSubAreaID != 0 {
		subSubArea, err := s.AreaCategoryQueryRepository.FindByID(query.SubSubAreaID)
		if err != nil {
			return nil, errors.Wrap(err, "failed to find area_category")
		}
		metasearchAreaID = subSubArea.MetasearchAreaID
		metasearchSubAreaID = subSubArea.MetasearchSubAreaID
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

	// 認証されているユーザーの場合、お気に入りしているかのフラグも取得する
	if ouser.Authenticated {
		reviews, err := s.ReviewQueryRepository.ShowReviewWithIsFavoriteListByParams(query, ouser.ID)
		if err != nil {
			return nil, errors.Wrap(err, "failed to show review list from repo")
		}
		return reviews, nil
	}

	reviews, err := s.ReviewQueryRepository.ShowReviewListByParams(query)
	if err != nil {
		return nil, errors.Wrap(err, "failed to show review list from repo")
	}

	return reviews, nil
}

func (s *ReviewQueryServiceImpl) ListFeed(ouser entity.OptionalUser, userID int, query *query.FindListPaginationQuery) (*entity.ReviewDetailWithIsFavoriteList, error) {
	if ouser.Authenticated {
		return s.ReviewQueryRepository.FindFeedReviewWithIsFavoriteListByUserID(ouser.ID, userID, query)
	}
	return s.ReviewQueryRepository.FindFeedReviewListByUserID(userID, query)
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

func (s *ReviewQueryServiceImpl) ListReviewCommentReplyByReviewCommentID(reviewCommentID int) ([]*entity.ReviewCommentReply, error) {
	return s.ReviewQueryRepository.FindReviewCommentReplyListByReviewCommentID(reviewCommentID)
}
