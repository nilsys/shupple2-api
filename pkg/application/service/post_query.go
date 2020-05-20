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
		ShowQueryByID(id int) (*entity.PostDetailWithHashtagAndIsFavorite, error)
		ShowQueryByIDForAuth(id, userID int) (*entity.PostDetailWithHashtagAndIsFavorite, error)
		ShowQueryBySlug(slug string) (*entity.PostDetailWithHashtagAndIsFavorite, error)
		ShowQueryBySlugForAuth(slug string, userID int) (*entity.PostDetailWithHashtagAndIsFavorite, error)
		ListByParams(query *query.FindPostListQuery) (*entity.PostList, error)
		ListByParamsForAuth(query *query.FindPostListQuery, userID int) (*entity.PostList, error)
		ListFeed(targetUserID int, query *query.FindListPaginationQuery) (*entity.PostList, error)
		ListFeedForAuth(userID, targetUserID int, query *query.FindListPaginationQuery) (*entity.PostList, error)
		ListFavoritePost(targetUserID int, query *query.FindListPaginationQuery) (*entity.PostList, error)
		ListFavoritePostForAuth(userID, targetUserID int, query *query.FindListPaginationQuery) (*entity.PostList, error)
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

func (s *PostQueryServiceImpl) ShowQueryByID(id int) (*entity.PostDetailWithHashtagAndIsFavorite, error) {
	post, err := s.PostQueryRepository.FindPostDetailWithHashtagByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed find post by id")
	}

	return post, nil
}

func (s *PostQueryServiceImpl) ShowQueryByIDForAuth(id, userID int) (*entity.PostDetailWithHashtagAndIsFavorite, error) {
	post, err := s.PostQueryRepository.FindPostDetailWithHashtagAndIsFavoriteByID(id, userID)
	if err != nil {
		return nil, errors.Wrap(err, "failed find post by id")
	}

	return post, nil
}

func (s *PostQueryServiceImpl) ShowQueryBySlug(slug string) (*entity.PostDetailWithHashtagAndIsFavorite, error) {
	post, err := s.PostQueryRepository.FindPostDetailWithHashtagBySlug(slug)
	if err != nil {
		return nil, errors.Wrap(err, "failed find post by id")
	}

	return post, nil
}

func (s *PostQueryServiceImpl) ShowQueryBySlugForAuth(slug string, userID int) (*entity.PostDetailWithHashtagAndIsFavorite, error) {
	post, err := s.PostQueryRepository.FindPostDetailWithHashtagAndIsFavoriteBySlug(slug, userID)
	if err != nil {
		return nil, errors.Wrap(err, "failed find post by id")
	}

	return post, nil
}

// 記事一覧参照
func (s *PostQueryServiceImpl) ListByParams(query *query.FindPostListQuery) (*entity.PostList, error) {
	posts, err := s.PostQueryRepository.FindListByParams(query)
	if err != nil {
		return nil, errors.Wrap(err, "failed find post by params")
	}
	return posts, nil
}

func (s *PostQueryServiceImpl) ListByParamsForAuth(query *query.FindPostListQuery, userID int) (*entity.PostList, error) {
	posts, err := s.PostQueryRepository.FindListWithIsFavoriteByParams(query, userID)
	if err != nil {
		return nil, errors.Wrap(err, "failed find post with is_favorite by params")
	}

	return posts, nil
}

// ユーザーがフォローしたユーザー or ハッシュタグの記事一覧参照
func (s *PostQueryServiceImpl) ListFeed(targetUserID int, query *query.FindListPaginationQuery) (*entity.PostList, error) {
	return s.PostQueryRepository.FindFeedListByUserID(targetUserID, query)
}

func (s *PostQueryServiceImpl) ListFeedForAuth(userID, targetUserID int, query *query.FindListPaginationQuery) (*entity.PostList, error) {
	return s.PostQueryRepository.FindFeedListWithIsFavoriteByUserID(userID, targetUserID, query)
}

// ユーザーがいいねした記事一覧参照
func (s *PostQueryServiceImpl) ListFavoritePost(targetUserID int, query *query.FindListPaginationQuery) (*entity.PostList, error) {
	return s.PostQueryRepository.FindFavoriteListByUserID(targetUserID, query)
}

func (s *PostQueryServiceImpl) ListFavoritePostForAuth(userID, targetUserID int, query *query.FindListPaginationQuery) (*entity.PostList, error) {
	return s.PostQueryRepository.FindFavoriteListWithIsFavoriteByUserID(userID, targetUserID, query)
}
