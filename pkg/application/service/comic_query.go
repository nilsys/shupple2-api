package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	// Comic参照系サービス
	ComicQueryService interface {
		Show(id int) (*entity.QueryComic, error)
		ShowList(query *query.FindListPaginationQuery) ([]*entity.Comic, error)
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
func (s *ComicQueryServiceImpl) Show(id int) (*entity.QueryComic, error) {
	comic, err := s.ComicQueryRepository.FindByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find comic")
	}

	return comic, nil
}

// Comic一覧参照
func (s *ComicQueryServiceImpl) ShowList(query *query.FindListPaginationQuery) ([]*entity.Comic, error) {
	return s.ComicQueryRepository.FindListOrderByCreatedAt(query)
}
