package factory

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/dto"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	PostCategoryFactory interface {
		NewPostCategoryFromPost(post *entity.Post) (*dto.PostAndCategories, error)
	}

	PostCategoryFactoryImpl struct {
		CategoryQueryRepository repository.CategoryQueryRepository
	}
)

var PostCategoryFactorySet = wire.NewSet(
	wire.Struct(new(PostCategoryFactoryImpl), "*"),
	wire.Bind(new(PostCategoryFactory), new(*PostCategoryFactoryImpl)),
)

// PostからCategoryを参照し、PostCategoryを作成
func (factory *PostCategoryFactoryImpl) NewPostCategoryFromPost(post *entity.Post) (*dto.PostAndCategories, error) {
	categories, err := factory.CategoryQueryRepository.FindByIDs(post.GetCategoryIDs()...)
	if err != nil {
		return nil, errors.Wrap(err, "failed get categories")
	}

	return &dto.PostAndCategories{
		Post:       post,
		Categories: categories,
	}, nil
}
