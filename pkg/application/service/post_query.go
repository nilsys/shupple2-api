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
		ShowQueryByID(id int) (*entity.PostDetailWithHashtag, error)
		ShowQueryBySlug(slug string) (*entity.PostDetailWithHashtag, error)
		ListFavoritePost(userID int, query *query.FindListPaginationQuery) ([]*entity.PostDetail, error)
		ListByParams(query *query.FindPostListQuery) (*entity.PostList, error)
		ListFeed(userID int, query *query.FindListPaginationQuery) ([]*entity.PostDetail, error)
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

func (s *PostQueryServiceImpl) ShowQueryByID(id int) (*entity.PostDetailWithHashtag, error) {
	return s.PostQueryRepository.FindPostDetailWithHashtagByID(id)
}

func (s *PostQueryServiceImpl) ShowQueryBySlug(slug string) (*entity.PostDetailWithHashtag, error) {
	return s.PostQueryRepository.FindPostDetailWithHashtagBySlug(slug)
}

// 記事一覧参照
func (s *PostQueryServiceImpl) ListByParams(query *query.FindPostListQuery) (*entity.PostList, error) {
	posts, err := s.PostQueryRepository.FindListByParams(query)
	if err != nil {
		return nil, errors.Wrap(err, "failed find post by params")
	}

	return posts, nil
}

// ユーザーがフォローしたユーザー or ハッシュタグの記事一覧参照
func (s *PostQueryServiceImpl) ListFeed(userID int, query *query.FindListPaginationQuery) ([]*entity.PostDetail, error) {
	return s.PostQueryRepository.FindFeedListByUserID(userID, query)
}

// ユーザーがいいねした記事一覧参照
func (s *PostQueryServiceImpl) ListFavoritePost(userID int, query *query.FindListPaginationQuery) ([]*entity.PostDetail, error) {
	return s.PostQueryRepository.FindFavoriteListByUserID(userID, query)
}
