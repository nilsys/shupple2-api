package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	CfReturnGiftQueryService interface {
		ListByCfProjectID(projectID int) (*entity.CfReturnGiftWithCountList, error)
	}

	CfReturnGiftQueryServiceImpl struct {
		repository.CfReturnGiftQueryRepository
	}
)

var CfReturnGiftQueryServiceSet = wire.NewSet(
	wire.Struct(new(CfReturnGiftQueryServiceImpl), "*"),
	wire.Bind(new(CfReturnGiftQueryService), new(*CfReturnGiftQueryServiceImpl)),
)

// TODO: SoldCount, SupporterCountは別で取る
func (s *CfReturnGiftQueryServiceImpl) ListByCfProjectID(projectID int) (*entity.CfReturnGiftWithCountList, error) {
	gifts, err := s.CfReturnGiftQueryRepository.FindByCfProjectID(projectID)
	if err != nil {
		return nil, errors.Wrap(err, "failed find cf_return_gift")
	}
	return gifts, nil
}
