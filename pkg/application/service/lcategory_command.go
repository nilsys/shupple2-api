package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	LcategoryCommandService interface {
		ImportFromWordpressByID(wordpressLcategoryID int) (*entity.Lcategory, error)
	}

	LcategoryCommandServiceImpl struct {
		LcategoryCommandRepository repository.LcategoryCommandRepository
		WordpressQueryRepository   repository.WordpressQueryRepository
		WordpressService           WordpressService
	}
)

var LcategoryCommandServiceSet = wire.NewSet(
	wire.Struct(new(LcategoryCommandServiceImpl), "*"),
	wire.Bind(new(LcategoryCommandService), new(*LcategoryCommandServiceImpl)),
)

func (r *LcategoryCommandServiceImpl) ImportFromWordpressByID(id int) (*entity.Lcategory, error) {
	wpLcategories, err := r.WordpressQueryRepository.FindLocationCategoriesByIDs([]int{id})
	if err != nil || len(wpLcategories) == 0 {
		return nil, serror.NewResourcesNotFoundError(err, "wordpress lcategory(id=%d)", id)
	}

	lcategory := r.WordpressService.ConvertLcategory(wpLcategories[0])
	if err := r.LcategoryCommandRepository.Store(lcategory); err != nil {
		return nil, errors.Wrap(err, "failed to store lcategory")
	}

	return lcategory, nil
}
