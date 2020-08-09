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
	CfProjectQueryScenario interface {
		List(query *query.FindCfProjectQuery, ouser *entity.OptionalUser) (*entity.CfProjectDetailList, map[int]bool, error)
		ListSupportComment(projectID, limit int, ouser *entity.OptionalUser) (*entity.CfProjectSupportCommentList, error)
		Show(id int, ouser *entity.OptionalUser) (*entity.CfProjectDetail, map[int]bool, error)
		ListSupported(user *entity.User, query *query.FindListPaginationQuery) (*entity.CfProjectDetailList, map[int]bool, error)
	}

	CfProjectQueryScenarioImpl struct {
		service.CfProjectQueryService
		repository.UserQueryRepository
	}
)

var CfProjectQueryScenarioSet = wire.NewSet(
	wire.Struct(new(CfProjectQueryScenarioImpl), "*"),
	wire.Bind(new(CfProjectQueryScenario), new(*CfProjectQueryScenarioImpl)),
)

func (s *CfProjectQueryScenarioImpl) List(query *query.FindCfProjectQuery, ouser *entity.OptionalUser) (*entity.CfProjectDetailList, map[int]bool, error) {
	var idIsFollowMap map[int]bool

	list, err := s.CfProjectQueryService.List(query)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed list cf_project")
	}

	if ouser.IsAuthorized() {
		// 認証されている場合、CfProject.Userをfollowしているかフラグを取得
		idIsFollowMap, err = s.IsFollowing(ouser.ID, list.UserIDs())
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed list user_following")
		}
	}

	return list, idIsFollowMap, nil
}

// MEMO: 現時点ではid:IsFollowのMapが必要ない
func (s *CfProjectQueryScenarioImpl) ListSupportComment(projectID, limit int, ouser *entity.OptionalUser) (*entity.CfProjectSupportCommentList, error) {
	list, err := s.CfProjectQueryService.ListSupportComment(projectID, limit)
	if err != nil {
		return nil, errors.Wrap(err, "failed list cf_project")
	}

	return list, nil
}

func (s *CfProjectQueryScenarioImpl) Show(id int, ouser *entity.OptionalUser) (*entity.CfProjectDetail, map[int]bool, error) {
	var idIsFollowMap map[int]bool

	project, err := s.CfProjectQueryService.Show(id)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed show cf_project")
	}

	if ouser.IsAuthorized() {
		// 認証されている場合、CfProject.Userをfollowしているかフラグを取得
		idIsFollowMap, err = s.IsFollowing(ouser.ID, []int{project.UserID})
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed list user_following")
		}
	}

	return project, idIsFollowMap, nil
}

func (s *CfProjectQueryScenarioImpl) ListSupported(user *entity.User, query *query.FindListPaginationQuery) (*entity.CfProjectDetailList, map[int]bool, error) {
	projects, err := s.CfProjectQueryService.ListSupported(user, query)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed list cf_project")
	}

	idIsFollowMap, err := s.IsFollowing(user.ID, projects.UserIDs())
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed list user_following")
	}

	return projects, idIsFollowMap, nil
}
