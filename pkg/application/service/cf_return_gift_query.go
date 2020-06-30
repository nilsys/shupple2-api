package service

import (
	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	CfReturnGiftQueryService interface {
		ListByCfProjectID(projectID int) (*entity.CfReturnGiftList, error)
	}

	CfReturnGiftQueryServiceImpl struct {
		repository.CfReturnGiftQueryRepository
	}
)

var CfReturnGiftQueryServiceSet = wire.NewSet(
	wire.Struct(new(CfReturnGiftQueryServiceImpl), "*"),
	wire.Bind(new(CfReturnGiftQueryService), new(*CfReturnGiftQueryServiceImpl)),
)

func (s *CfReturnGiftQueryServiceImpl) ListByCfProjectID(projectID int) (*entity.CfReturnGiftList, error) {
	return s.CfReturnGiftQueryRepository.FindByCfProjectID(projectID)
}
