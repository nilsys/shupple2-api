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
		List(query *query.FindCfProjectQuery, ouser *entity.OptionalUser) (*entity.CfProjectDetailList, *entity.UserRelationFlgMap, map[int]bool, error)
		ListSupportComment(projectID, limit int, ouser *entity.OptionalUser) (*entity.CfProjectSupportCommentList, error)
		Show(id int, ouser *entity.OptionalUser) (*entity.CfProjectDetail, *entity.UserRelationFlgMap, map[int]bool, error)
		ListSupported(user *entity.User, query *query.FindListPaginationQuery) (*entity.CfProjectDetailList, *entity.UserRelationFlgMap, map[int]bool, error)
	}

	CfProjectQueryScenarioImpl struct {
		service.CfProjectQueryService
		service.UserQueryService
		repository.CfProjectQueryRepository
	}
)

var CfProjectQueryScenarioSet = wire.NewSet(
	wire.Struct(new(CfProjectQueryScenarioImpl), "*"),
	wire.Bind(new(CfProjectQueryScenario), new(*CfProjectQueryScenarioImpl)),
)

func (s *CfProjectQueryScenarioImpl) List(query *query.FindCfProjectQuery, ouser *entity.OptionalUser) (*entity.CfProjectDetailList, *entity.UserRelationFlgMap, map[int]bool, error) {
	idRelationFlgMap := &entity.UserRelationFlgMap{}
	var idIsSupportMap map[int]bool

	list, err := s.CfProjectQueryService.List(query)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed list cf_project")
	}

	if ouser.IsAuthorized() {
		// 認証されている場合、CfProject.Userをfollow, blockしているかフラグを取得
		idRelationFlgMap, err = s.UserQueryService.RelationFlgMaps(ouser.ID, list.UserIDs())
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "failed list user_following")
		}

		// CfProjectを支援したかフラグを取得
		idIsSupportMap, err = s.IsSupported(ouser.ID, list.IDs())
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "failed list user_support")
		}
	}

	return list, idRelationFlgMap, idIsSupportMap, nil
}

// MEMO: 現時点ではid:IsFollowのMapが必要ない
func (s *CfProjectQueryScenarioImpl) ListSupportComment(projectID, limit int, ouser *entity.OptionalUser) (*entity.CfProjectSupportCommentList, error) {
	list, err := s.CfProjectQueryService.ListSupportComment(projectID, limit)
	if err != nil {
		return nil, errors.Wrap(err, "failed list cf_project")
	}

	return list, nil
}

func (s *CfProjectQueryScenarioImpl) Show(id int, ouser *entity.OptionalUser) (*entity.CfProjectDetail, *entity.UserRelationFlgMap, map[int]bool, error) {
	idRelationFlgMap := &entity.UserRelationFlgMap{}
	var idIsSupportMap map[int]bool

	project, err := s.CfProjectQueryService.Show(id)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed show cf_project")
	}

	if ouser.IsAuthorized() {
		// 認証されている場合、CfProject.Userをfollow, blockしているかフラグを取得
		idRelationFlgMap, err = s.UserQueryService.RelationFlgMaps(ouser.ID, []int{project.UserID})
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "failed list user_following")
		}

		// CfProjectを支援したかフラグを取得
		idIsSupportMap, err = s.IsSupported(ouser.ID, []int{project.ID})
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "failed list user_support")
		}
	}

	return project, idRelationFlgMap, idIsSupportMap, nil
}

func (s *CfProjectQueryScenarioImpl) ListSupported(user *entity.User, query *query.FindListPaginationQuery) (*entity.CfProjectDetailList, *entity.UserRelationFlgMap, map[int]bool, error) {
	projects, err := s.CfProjectQueryService.ListSupported(user, query)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed list cf_project")
	}

	idRelationFlgMap, err := s.UserQueryService.RelationFlgMaps(user.ID, projects.UserIDs())
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed list user_following")
	}

	idIsSupportMap, err := s.IsSupported(user.ID, projects.IDs())
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed list user_support")
	}

	return projects, idRelationFlgMap, idIsSupportMap, nil
}
