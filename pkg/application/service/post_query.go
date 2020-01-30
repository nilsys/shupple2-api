package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/dto"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/factory"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	PostQueryService interface {
		FindByID(id int) (*entity.Post, error)
		// TODO: 命名
		FindByParams(query *query.FindPostListQuery) ([]*dto.PostAndCategories, error)
	}

	PostQueryServiceImpl struct {
		PostQueryRepository repository.PostQueryRepository
		PostCategoryFactory factory.PostCategoryFactory
	}
)

var PostQueryServiceSet = wire.NewSet(
	wire.Struct(new(PostQueryServiceImpl), "*"),
	wire.Bind(new(PostQueryService), new(*PostQueryServiceImpl)),
)

func (r *PostQueryServiceImpl) FindByID(id int) (*entity.Post, error) {
	post, err := r.PostQueryRepository.FindByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get post")
	}

	return post, nil
}

func (r *PostQueryServiceImpl) FindByParams(query *query.FindPostListQuery) ([]*dto.PostAndCategories, error) {
	var postAndCategoriesList []*dto.PostAndCategories

	posts, err := r.PostQueryRepository.FindByParams(query)
	if err != nil {
		return nil, errors.Wrap(err, "failed find post by params")
	}

	for _, post := range posts {
		postAndCategories, err := r.PostCategoryFactory.NewPostCategoryFromPost(post)
		if err != nil {
			return nil, errors.Wrap(err, "failed generate post and categories")
		}

		postAndCategoriesList = append(postAndCategoriesList, postAndCategories)
	}

	return postAndCategoriesList, nil
}
