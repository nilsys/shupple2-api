package service

import (
	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	// Comic参照系サービス
	ComicQueryService interface {
		Show(id int, ouser *entity.OptionalUser) (*entity.ComicDetail, error)
		List(query *query.FindListPaginationQuery, ouser *entity.OptionalUser) (*entity.ComicList, error)
	}

	// Comic参照系サービス実装
	ComicQueryServiceImpl struct {
		repository.ComicQueryRepository
		repository.UserQueryRepository
	}
)

var ComicQueryServiceSet = wire.NewSet(
	wire.Struct(new(ComicQueryServiceImpl), "*"),
	wire.Bind(new(ComicQueryService), new(*ComicQueryServiceImpl)),
)

// Comic参照
func (s *ComicQueryServiceImpl) Show(id int, ouser *entity.OptionalUser) (*entity.ComicDetail, error) {
	if ouser.Authenticated {
		return s.ComicQueryRepository.FindWithIsFavoriteByID(id, ouser.ID)
	}
	return s.ComicQueryRepository.FindByID(id)
}

// Comic一覧参照
func (s *ComicQueryServiceImpl) List(query *query.FindListPaginationQuery, ouser *entity.OptionalUser) (*entity.ComicList, error) {
	if ouser.Authenticated {
		return s.ComicQueryRepository.FindWithIsFavoriteListOrderByCreatedAt(query, ouser.ID)
	}
	return s.ComicQueryRepository.FindListOrderByCreatedAt(query)
}
