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
	FeatureCommandService interface {
		ImportFromWordpressByID(wordpressFeatureID int) (*entity.Feature, error)
	}

	FeatureCommandServiceImpl struct {
		FeatureCommandRepository repository.FeatureCommandRepository
		WordpressQueryRepository repository.WordpressQueryRepository
		WordpressService
		TransactionService
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

	var feature *entity.Feature
	err = r.TransactionService.Do(func(c context.Context) error {
		feature, err = r.FeatureCommandRepository.Lock(c, id)
		if err != nil {
			if !serror.IsErrorCode(err, serror.CodeNotFound) {
				return errors.Wrap(err, "failed to get feature")
			}
			feature = &entity.Feature{}
		}

		if err := r.WordpressService.PatchFeature(feature, wpFeatures[0]); err != nil {
			return errors.Wrap(err, "failed  to patch feature")
		}

		if err := r.FeatureCommandRepository.Store(c, feature); err != nil {
			return errors.Wrap(err, "failed to store feature")
		}

		return nil
	})

	return feature, nil
}
