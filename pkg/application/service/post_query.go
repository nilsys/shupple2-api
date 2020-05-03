package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	// Post参照系サービス
	PostQueryService interface {
		ShowByID(id int) (*entity.Post, error)
		ShowQueryByID(id int, ouser entity.OptionalUser) (*entity.PostDetailWithHashtagAndIsFavorite, error)
		ShowQueryBySlug(slug string) (*entity.PostDetailWithHashtag, error)
		ListByParams(query *query.FindPostListQuery, ouser entity.OptionalUser) (*entity.PostList, error)
		ListFeed(ouser entity.OptionalUser, targetUserID int, query *query.FindListPaginationQuery) (*entity.PostList, error)
		ListFavoritePost(ouser entity.OptionalUser, targetUserID int, query *query.FindListPaginationQuery) (*entity.PostList, error)
	}

	// Post参照系サービス実装
	PostQueryServiceImpl struct {
		PostQueryRepository repository.PostQueryRepository
	}
)

var PostQueryServiceSet = wire.NewSet(
	wire.Struct(new(PostQueryServiceImpl), "*"),
	wire.Bind(new(PostQueryService), new(*PostQueryServiceImpl)),
)

func (s *PostQueryServiceImpl) ShowByID(id int) (*entity.Post, error) {
	post, err := s.PostQueryRepository.FindByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get post")
	}

	return post, nil
}

func (s *PostQueryServiceImpl) ShowQueryByID(id int, ouser entity.OptionalUser) (*entity.PostDetailWithHashtagAndIsFavorite, error) {
	if ouser.Authenticated {
		return s.PostQueryRepository.FindPostDetailWithHashtagAndIsFavoriteByID(id, ouser.ID)
	}
	return s.PostQueryRepository.FindPostDetailWithHashtagByID(id)
}

func (s *PostQueryServiceImpl) ShowQueryBySlug(slug string) (*entity.PostDetailWithHashtag, error) {
	return s.PostQueryRepository.FindPostDetailWithHashtagBySlug(slug)
}

// 記事一覧参照
func (s *PostQueryServiceImpl) ListByParams(query *query.FindPostListQuery, ouser entity.OptionalUser) (*entity.PostList, error) {
	if ouser.Authenticated {
		posts, err := s.PostQueryRepository.FindListWithIsFavoriteByParams(query, ouser.ID)
		if err != nil {
			return nil, errors.Wrap(err, "failed find post with is_favorite by params")
		}
		return posts, nil
	}

	posts, err := s.PostQueryRepository.FindListByParams(query)
	if err != nil {
		return nil, errors.Wrap(err, "failed find post by params")
	}

	return posts, nil
}

// ユーザーがフォローしたユーザー or ハッシュタグの記事一覧参照
func (s *PostQueryServiceImpl) ListFeed(ouser entity.OptionalUser, targetUserID int, query *query.FindListPaginationQuery) (*entity.PostList, error) {
	if ouser.Authenticated {
		return s.PostQueryRepository.FindFeedListWithIsFavoriteByUserID(ouser.ID, targetUserID, query)
	}
	return s.PostQueryRepository.FindFeedListByUserID(targetUserID, query)
}

// ユーザーがいいねした記事一覧参照
func (s *PostQueryServiceImpl) ListFavoritePost(ouser entity.OptionalUser, targetUserID int, query *query.FindListPaginationQuery) (*entity.PostList, error) {
	if ouser.Authenticated {
		return s.PostQueryRepository.FindFavoriteListWithIsFavoriteByUserID(ouser.ID, targetUserID, query)
	}
	return s.PostQueryRepository.FindFavoriteListByUserID(targetUserID, query)
}
