package service

import (
	"github.com/google/wire"

	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity/wordpress"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	CfReturnGiftCommandService interface {
		ImportFromWordpressByID(id int) error
	}

	CfReturnGiftCommandServiceImpl struct {
		repository.CfReturnGiftCommandRepository
		repository.WordpressQueryRepository
		WordpressService
		TransactionService
	}
)

var CfReturnGiftCommandServiceSet = wire.NewSet(
	wire.Struct(new(CfReturnGiftCommandServiceImpl), "*"),
	wire.Bind(new(CfReturnGiftCommandService), new(*CfReturnGiftCommandServiceImpl)),
)

func (s *CfReturnGiftCommandServiceImpl) ImportFromWordpressByID(id int) error {
	wpCfReturnGift, err := s.WordpressQueryRepository.FindCfReturnGiftByID(id)
	if err != nil {
		return errors.Wrapf(err, "failed to get wordpress cfReturnGift(id=%d)", id)
	}

	if wpCfReturnGift.Status != wordpress.StatusPublish {
		if err := s.CfReturnGiftCommandRepository.DeleteByID(id); err != nil {
			return errors.Wrapf(err, "failed to delete cfReturnGift(id=%d)", id)
		}

		return serror.New(nil, serror.CodeImportDeleted, "try to import deleted cfReturnGift")
	}

	cfReturnGift, err := s.WordpressService.NewCfReturnGift(wpCfReturnGift)
	if err != nil {
		return errors.Wrap(err, "failed  to initialize cfReturnGift")
	}

	if err := s.CfReturnGiftCommandRepository.Store(cfReturnGift); err != nil {
		return errors.Wrap(err, "failed to store cfReturnGift")
	}

	return nil
}
