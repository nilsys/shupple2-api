package scenario

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/factory"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

type (
	PostQueryScenario interface {
		ShowByID(id int, ouser entity.OptionalUser) (*entity.PostDetailWithHashtagAndIsFavorite, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, error)
		ShowBySlug(slug string, ouser entity.OptionalUser) (*entity.PostDetailWithHashtagAndIsFavorite, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, error)
		ListByParams(query *query.FindPostListQuery, ouser entity.OptionalUser) (*entity.PostList, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, error)
		ListFeed(targetUserID int, query *query.FindListPaginationQuery, ouser entity.OptionalUser) (*entity.PostList, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, error)
		LitFavorite(targetUserID int, query *query.FindListPaginationQuery, ouser entity.OptionalUser) (*entity.PostList, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, error)
	}

	PostQueryScenarioImpl struct {
		factory.CategoryIDMapFactory
		service.PostQueryService
	}
)

var PostQueryScenarioSet = wire.NewSet(
	wire.Struct(new(PostQueryScenarioImpl), "*"),
	wire.Bind(new(PostQueryScenario), new(*PostQueryScenarioImpl)),
)

func (s *PostQueryScenarioImpl) ShowByID(id int, ouser entity.OptionalUser) (*entity.PostDetailWithHashtagAndIsFavorite, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, error) {
	var (
		post *entity.PostDetailWithHashtagAndIsFavorite
		err  error
	)

	if ouser.Authenticated {
		post, err = s.PostQueryService.ShowQueryByIDForAuth(id, ouser.ID)
	} else {
		post, err = s.PostQueryService.ShowQueryByID(id)
	}
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed find post by id")
	}

	areaCategoriesMap, themeCategoriesMap, err := s.CategoryIDMapFactory.GenerateCategoryIDMap(post.AreaCategoryIDs(), post.ThemeCategoryIDs())
	if err != nil {
		return nil, areaCategoriesMap, themeCategoriesMap, errors.Wrap(err, "failed gen category map")
	}
	return post, areaCategoriesMap, themeCategoriesMap, nil
}

func (s *PostQueryScenarioImpl) ShowBySlug(slug string, ouser entity.OptionalUser) (*entity.PostDetailWithHashtagAndIsFavorite, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, error) {
	var (
		post *entity.PostDetailWithHashtagAndIsFavorite
		err  error
	)
	if ouser.Authenticated {
		post, err = s.PostQueryService.ShowQueryBySlugForAuth(slug, ouser.ID)
	} else {
		post, err = s.PostQueryService.ShowQueryBySlug(slug)
	}

	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed find post by slug")
	}

	areaCategoriesMap, themeCategoriesMap, err := s.CategoryIDMapFactory.GenerateCategoryIDMap(post.AreaCategoryIDs(), post.ThemeCategoryIDs())
	if err != nil {
		return nil, areaCategoriesMap, themeCategoriesMap, errors.Wrap(err, "failed gen category map")
	}
	return post, areaCategoriesMap, themeCategoriesMap, nil
}

func (s *PostQueryScenarioImpl) ListByParams(query *query.FindPostListQuery, ouser entity.OptionalUser) (*entity.PostList, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, error) {
	var (
		posts *entity.PostList
		err   error
	)

	if ouser.Authenticated {
		posts, err = s.PostQueryService.ListByParamsForAuth(query, ouser.ID)
	} else {
		posts, err = s.PostQueryService.ListByParams(query)
	}
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed find list post by params")
	}

	areaCategoriesMap, themeCategoriesMap, err := s.CategoryIDMapFactory.GenerateCategoryIDMap(posts.AreaCategoryIDs(), posts.ThemeCategoryIDs())
	if err != nil {
		return nil, areaCategoriesMap, themeCategoriesMap, errors.Wrap(err, "failed gen category map")
	}
	return posts, areaCategoriesMap, themeCategoriesMap, nil
}

func (s *PostQueryScenarioImpl) ListFeed(targetUserID int, query *query.FindListPaginationQuery, ouser entity.OptionalUser) (*entity.PostList, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, error) {
	var (
		posts *entity.PostList
		err   error
	)
	if ouser.Authenticated {
		posts, err = s.PostQueryService.ListFeedForAuth(ouser.ID, targetUserID, query)
	} else {
		posts, err = s.PostQueryService.ListFeed(targetUserID, query)
	}
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed list feed post")
	}
	areaCategoriesMap, themeCategoriesMap, err := s.CategoryIDMapFactory.GenerateCategoryIDMap(posts.AreaCategoryIDs(), posts.ThemeCategoryIDs())
	if err != nil {
		return nil, areaCategoriesMap, themeCategoriesMap, errors.Wrap(err, "failed gen category map")
	}
	return posts, areaCategoriesMap, themeCategoriesMap, nil
}

func (s *PostQueryScenarioImpl) LitFavorite(targetUserID int, query *query.FindListPaginationQuery, ouser entity.OptionalUser) (*entity.PostList, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, error) {
	var (
		posts *entity.PostList
		err   error
	)
	if ouser.Authenticated {
		posts, err = s.PostQueryService.ListFavoritePostForAuth(ouser.ID, targetUserID, query)
	} else {
		posts, err = s.PostQueryService.ListFavoritePost(targetUserID, query)
	}
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed list favorite post")
	}
	areaCategoriesMap, themeCategoriesMap, err := s.CategoryIDMapFactory.GenerateCategoryIDMap(posts.AreaCategoryIDs(), posts.ThemeCategoryIDs())
	if err != nil {
		return nil, areaCategoriesMap, themeCategoriesMap, errors.Wrap(err, "failed gen category map")
	}
	return posts, areaCategoriesMap, themeCategoriesMap, nil
}
