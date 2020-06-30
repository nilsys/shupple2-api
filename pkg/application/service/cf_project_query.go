package service

import (
	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	CfProjectQueryService interface {
		ListSupportComment(projectID, limit int) ([]*entity.CfProjectSupportComment, error)
		Show(id int) (*entity.CfProject, error)
		List(query *query.FindCfProjectQuery) (*entity.CfProjectList, error)
	}

	CfProjectQueryServiceImpl struct {
		repository.CfProjectQueryRepository
	}
)

var CfProjectQueryServiceSet = wire.NewSet(
	wire.Struct(new(CfProjectQueryServiceImpl), "*"),
	wire.Bind(new(CfProjectQueryService), new(*CfProjectQueryServiceImpl)),
)

func (s *CfProjectQueryServiceImpl) ListSupportComment(projectID, limit int) ([]*entity.CfProjectSupportComment, error) {
	return s.CfProjectQueryRepository.FindSupportCommentListByCfProjectID(projectID, limit)
}

func (s *CfProjectQueryServiceImpl) Show(id int) (*entity.CfProject, error) {
	return s.CfProjectQueryRepository.FindByID(id)
}

func (s *CfProjectQueryServiceImpl) List(query *query.FindCfProjectQuery) (*entity.CfProjectList, error) {
	return s.CfProjectQueryRepository.FindListByQuery(query)
}
