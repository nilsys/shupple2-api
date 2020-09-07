package service

import (
	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	// Post参照系サービス
	PostQueryService interface {
		ShowByID(id int, ouser *entity.OptionalUser) (*entity.PostDetailWithHashtagAndIsFavorite, error)
		ShowBySlug(slug string, ouser *entity.OptionalUser) (*entity.PostDetailWithHashtagAndIsFavorite, error)
		ListByParams(query *query.FindPostListQuery, ouser *entity.OptionalUser) (*entity.PostList, error)
		ListFeed(query *query.FindListPaginationQuery, user *entity.User) (*entity.PostList, error)
		ListFavoritePost(targetUserID int, query *query.FindListPaginationQuery, ouser *entity.OptionalUser) (*entity.PostList, error)
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

func (s *PostQueryServiceImpl) ShowByID(id int, ouser *entity.OptionalUser) (*entity.PostDetailWithHashtagAndIsFavorite, error) {
	if ouser.IsAuthorized() {
		return s.PostQueryRepository.FindPostDetailWithHashtagAndIsFavoriteByID(id, ouser.ID)
	}
	return s.PostQueryRepository.FindPostDetailWithHashtagByID(id)
}

func (s *PostQueryServiceImpl) ShowBySlug(slug string, ouser *entity.OptionalUser) (*entity.PostDetailWithHashtagAndIsFavorite, error) {
	if ouser.IsAuthorized() {
		return s.PostQueryRepository.FindPostDetailWithHashtagAndIsFavoriteBySlug(slug, ouser.ID)

	}
	return s.PostQueryRepository.FindPostDetailWithHashtagBySlug(slug)
}

// 記事一覧参照
func (s *PostQueryServiceImpl) ListByParams(query *query.FindPostListQuery, ouser *entity.OptionalUser) (*entity.PostList, error) {
	if ouser.IsAuthorized() {
		return s.PostQueryRepository.FindListWithIsFavoriteByParams(query, ouser.ID)
	}
	return s.PostQueryRepository.FindListByParams(query)
}

// ユーザーがフォローしたユーザー or ハッシュタグの記事一覧参照
func (s *PostQueryServiceImpl) ListFeed(query *query.FindListPaginationQuery, user *entity.User) (*entity.PostList, error) {
	return s.PostQueryRepository.FindFeedListWithIsFavoriteByUserID(user.ID, query)
}

// ユーザーがいいねした記事一覧参照
func (s *PostQueryServiceImpl) ListFavoritePost(targetUserID int, query *query.FindListPaginationQuery, ouser *entity.OptionalUser) (*entity.PostList, error) {
	if ouser.IsAuthorized() {
		return s.PostQueryRepository.FindFavoriteListWithIsFavoriteByUserID(ouser.ID, targetUserID, query)
	}
	return s.PostQueryRepository.FindFavoriteListByUserID(targetUserID, query)
}
