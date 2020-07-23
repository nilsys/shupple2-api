package service

import (
	"context"

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

	return s.TransactionService.Do(func(c context.Context) error {
		if err := s.CfReturnGiftCommandRepository.UndeleteByID(c, id); err != nil {
			return errors.Wrapf(err, "failed to undelete cf_return_gift(id=%d)", id)
		}

		if err := s.CfReturnGiftCommandRepository.Store(c, cfReturnGift); err != nil {
			return errors.Wrap(err, "failed to store cf_return_gift")
		}

		return nil
	})
}
