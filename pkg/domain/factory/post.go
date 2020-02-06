package factory

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/dto"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	// PostDetail生成ファクトリー
	PostDetailFactory interface {
		NewPostCategoryFromPost(post *entity.Post) (*dto.PostDetail, error)
	}

	// PostDetail生成ファクトリー実装
	PostDetailFactoryImpl struct {
		CategoryQueryRepository repository.CategoryQueryRepository
		UserQueryRepository     repository.UserQueryRepository
	}
)

var PostDetailFactorySet = wire.NewSet(
	wire.Struct(new(PostDetailFactoryImpl), "*"),
	wire.Bind(new(PostDetailFactory), new(*PostDetailFactoryImpl)),
)

// Postから[]Category, Userを参照し、PostCategoryを作成
func (factory *PostDetailFactoryImpl) NewPostCategoryFromPost(post *entity.Post) (*dto.PostDetail, error) {
	categories, err := factory.CategoryQueryRepository.FindByIDs(post.GetCategoryIDs()...)
	if err != nil {
		return nil, errors.Wrap(err, "failed get categories")
	}

	user, err := factory.UserQueryRepository.FindByID(post.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "failed get user")
	}

	return &dto.PostDetail{
		Post:       post,
		Categories: categories,
		User:       user,
	}, nil
}
