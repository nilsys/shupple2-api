package service

import (
	"github.com/google/wire"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	CategoryQueryService interface {
		ShowBySlug(slug string) (*entity.Category, error)
	}

	CategoryQueryServiceImpl struct {
		Repository repository.CategoryQueryRepository
	}
)

var CategoryQueryServiceSet = wire.NewSet(
	wire.Struct(new(CategoryQueryServiceImpl), "*"),
	wire.Bind(new(CategoryQueryService), new(*CategoryQueryServiceImpl)),
)

func (r *CategoryQueryServiceImpl) ShowBySlug(slug string) (*entity.Category, error) {
	return r.Repository.FindBySlug(slug)
}
