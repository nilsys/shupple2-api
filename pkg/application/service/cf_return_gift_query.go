package service

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	CfReturnGiftQueryService interface {
		ListByCfProjectID(projectID int) (*entity.CfReturnGiftList, *entity.CfReturnGiftSoldCountList, error)
	}

	CfReturnGiftQueryServiceImpl struct {
		repository.CfReturnGiftQueryRepository
	}
)

var CfReturnGiftQueryServiceSet = wire.NewSet(
	wire.Struct(new(CfReturnGiftQueryServiceImpl), "*"),
	wire.Bind(new(CfReturnGiftQueryService), new(*CfReturnGiftQueryServiceImpl)),
)

func (s *CfReturnGiftQueryServiceImpl) ListByCfProjectID(projectID int) (*entity.CfReturnGiftList, *entity.CfReturnGiftSoldCountList, error) {
	gifts, err := s.CfReturnGiftQueryRepository.FindByCfProjectID(projectID)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed find cf_return_gift")
	}
	soldCountList, err := s.CfReturnGiftQueryRepository.FindSoldCountByReturnGiftIDs(context.Background(), gifts.IDs())
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed find cf_return_gift.sold_count")
	}
	return gifts, soldCountList, nil
}
