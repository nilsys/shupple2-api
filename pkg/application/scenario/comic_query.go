package scenario

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	ComicQueryScenario interface {
		Show(id int, ouser *entity.OptionalUser) (*entity.ComicDetail, map[int]bool, error)
		List(query *query.FindListPaginationQuery, ouser *entity.OptionalUser) (*entity.ComicList, error)
	}

	ComicQueryQueryScenarioImpl struct {
		service.ComicQueryService
		repository.UserQueryRepository
	}
)

var ComicQueryScenarioSet = wire.NewSet(
	wire.Struct(new(ComicQueryQueryScenarioImpl), "*"),
	wire.Bind(new(ComicQueryScenario), new(*ComicQueryQueryScenarioImpl)),
)

func (s *ComicQueryQueryScenarioImpl) Show(id int, ouser *entity.OptionalUser) (*entity.ComicDetail, map[int]bool, error) {
	var idIsFollowMap map[int]bool

	comic, err := s.ComicQueryService.Show(id, ouser)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed show comic")
	}

	if ouser.IsAuthorized() {
		// 認証されている場合、Comic.Userをfollowしているかフラグを取得
		idIsFollowMap, err = s.IsFollowing(ouser.ID, []int{comic.UserID})
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed list user_following")
		}
	}

	return comic, idIsFollowMap, nil
}

// MEMO: 現時点ではid:IsFollowのMapが必要ない
func (s *ComicQueryQueryScenarioImpl) List(query *query.FindListPaginationQuery, ouser *entity.OptionalUser) (*entity.ComicList, error) {
	list, err := s.ComicQueryService.List(query, ouser)
	if err != nil {
		return nil, errors.Wrap(err, "failed list comic")
	}

	return list, nil
}
