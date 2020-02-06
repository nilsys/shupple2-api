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
	// Post参照系サービス
	PostQueryService interface {
		ShowByID(id int) (*entity.Post, error)
		ShowListByParams(query *query.FindPostListQuery) ([]*dto.PostDetail, error)
	}

	// Post参照系サービス実装
	PostQueryServiceImpl struct {
		PostQueryRepository repository.PostQueryRepository
		PostCategoryFactory factory.PostDetailFactory
	}
)

var PostQueryServiceSet = wire.NewSet(
	wire.Struct(new(PostQueryServiceImpl), "*"),
	wire.Bind(new(PostQueryService), new(*PostQueryServiceImpl)),
)

func (r *PostQueryServiceImpl) ShowByID(id int) (*entity.Post, error) {
	post, err := r.PostQueryRepository.FindByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get post")
	}

	return post, nil
}

func (r *PostQueryServiceImpl) ShowListByParams(query *query.FindPostListQuery) ([]*dto.PostDetail, error) {
	var postAndCategoriesList []*dto.PostDetail

	posts, err := r.PostQueryRepository.FindListByParams(query)
	if err != nil {
		return nil, errors.Wrap(err, "failed find post by params")
	}

	// MEMO: factoryの方に移動させた方がすっきりする(factoryがカオスになるが)
	for _, post := range posts {
		postAndCategories, err := r.PostCategoryFactory.NewPostCategoryFromPost(post)
		if err != nil {
			return nil, errors.Wrap(err, "failed generate post and categories")
		}
		postAndCategoriesList = append(postAndCategoriesList, postAndCategories)
	}

	return postAndCategoriesList, nil
}
