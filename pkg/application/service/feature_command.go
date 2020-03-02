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
	FeatureCommandService interface {
		ImportFromWordpressByID(wordpressFeatureID int) (*entity.Feature, error)
	}

	FeatureCommandServiceImpl struct {
		FeatureCommandRepository repository.FeatureCommandRepository
		WordpressQueryRepository repository.WordpressQueryRepository
		WordpressService         WordpressService
	}
)

var FeatureCommandServiceSet = wire.NewSet(
	wire.Struct(new(FeatureCommandServiceImpl), "*"),
	wire.Bind(new(FeatureCommandService), new(*FeatureCommandServiceImpl)),
)

func (r *FeatureCommandServiceImpl) ImportFromWordpressByID(id int) (*entity.Feature, error) {
	wpFeatures, err := r.WordpressQueryRepository.FindFeaturesByIDs([]int{id})
	if err != nil || len(wpFeatures) == 0 {
		return nil, serror.NewResourcesNotFoundError(err, "wordpress feature(id=%d)", id)
	}

	if wpFeatures[0].Status != wordpress.StatusPublish {
		if err := r.FeatureCommandRepository.DeleteByID(id); err != nil {
			return nil, errors.Wrapf(err, "failed to delete feature(id=%d)", id)
		}

		return nil, serror.New(nil, serror.CodeImportDeleted, "try to import deleted feature")
	}

	feature, err := r.WordpressService.ConvertFeature(wpFeatures[0])
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert feature")
	}

	if err := r.FeatureCommandRepository.Store(feature); err != nil {
		return nil, errors.Wrap(err, "failed to store feature")
	}

	return feature, nil
}
