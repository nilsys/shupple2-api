package service

import (
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
		WordpressService         WordpressService
	}
)

var VlogCommandServiceSet = wire.NewSet(
	wire.Struct(new(VlogCommandServiceImpl), "*"),
	wire.Bind(new(VlogCommandService), new(*VlogCommandServiceImpl)),
)

func (r *VlogCommandServiceImpl) ImportFromWordpressByID(id int) (*entity.Vlog, error) {
	wpVlogs, err := r.WordpressQueryRepository.FindVlogsByIDs([]int{id})
	if err != nil || len(wpVlogs) == 0 {
		return nil, serror.NewResourcesNotFoundError(err, "wordpress vlog(id=%d)", id)
	}

	if wpVlogs[0].Status != wordpress.StatusPublish {
		if err := r.VlogCommandRepository.DeleteByID(id); err != nil {
			return nil, errors.Wrapf(err, "failed to delete vlog(id=%d)", id)
		}

		return nil, serror.New(nil, serror.CodeImportDeleted, "try to import deleted vlog")
	}

	vlog, err := r.WordpressService.ConvertVlog(wpVlogs[0])
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert vlog")
	}

	if err := r.VlogCommandRepository.Store(vlog); err != nil {
		return nil, errors.Wrap(err, "failed to store vlog")
	}

	return vlog, nil
}
