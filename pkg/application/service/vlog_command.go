package service

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity/wordpress"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	VlogCommandService interface {
		ImportFromWordpressByID(wordpressVlogID int) (*entity.Vlog, error)
	}

	VlogCommandServiceImpl struct {
		VlogCommandRepository    repository.VlogCommandRepository
		WordpressQueryRepository repository.WordpressQueryRepository
		WordpressService
		TransactionService
	}
)

var VlogCommandServiceSet = wire.NewSet(
	wire.Struct(new(VlogCommandServiceImpl), "*"),
	wire.Bind(new(VlogCommandService), new(*VlogCommandServiceImpl)),
)

func (r *VlogCommandServiceImpl) ImportFromWordpressByID(id int) (*entity.Vlog, error) {
	wpVlog, err := r.WordpressQueryRepository.FindVlogByID(id)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get wordpress vlog(id=%d)", id)
	}

	if wpVlog.Status != wordpress.StatusPublish {
		if err := r.VlogCommandRepository.DeleteByID(id); err != nil {
			return nil, errors.Wrapf(err, "failed to delete vlog(id=%d)", id)
		}

		return nil, serror.New(nil, serror.CodeImportDeleted, "try to import deleted vlog")
	}

	var vlog *entity.Vlog
	err = r.TransactionService.Do(func(c context.Context) error {
		vlog, err = r.VlogCommandRepository.Lock(c, id)
		if err != nil {
			if !serror.IsErrorCode(err, serror.CodeNotFound) {
				return errors.Wrap(err, "failed to get vlog")
			}
			vlog = &entity.Vlog{}
		}

		if err := r.WordpressService.PatchVlog(vlog, wpVlog); err != nil {
			return errors.Wrap(err, "failed  to patch vlog")
		}

		if err := r.VlogCommandRepository.Store(c, vlog); err != nil {
			return errors.Wrap(err, "failed to store vlog")
		}

		return nil
	})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return vlog, nil
}
