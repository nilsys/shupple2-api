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
		Show(id int, ouser *entity.OptionalUser) (*entity.PostDetailWithHashtagAndIsFavorite, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, *entity.UserRelationFlgMap, error)
		ShowBySlug(slug string, ouser *entity.OptionalUser) (*entity.PostDetailWithHashtagAndIsFavorite, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, *entity.UserRelationFlgMap, error)
		ListByParams(query *query.FindPostListQuery, ouser *entity.OptionalUser) (*entity.PostList, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, *entity.UserRelationFlgMap, error)
		ListFeed(query *query.FindListPaginationQuery, user *entity.User) (*entity.PostList, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, *entity.UserRelationFlgMap, error)
		LitFavorite(targetUserID int, query *query.FindListPaginationQuery, ouser *entity.OptionalUser) (*entity.PostList, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, *entity.UserRelationFlgMap, error)
	}

	PostQueryScenarioImpl struct {
		factory.CategoryIDMapFactory
		service.PostQueryService
		service.UserQueryService
	}
)

var PostQueryScenarioSet = wire.NewSet(
	wire.Struct(new(PostQueryScenarioImpl), "*"),
	wire.Bind(new(PostQueryScenario), new(*PostQueryScenarioImpl)),
)

func (s *PostQueryScenarioImpl) Show(id int, ouser *entity.OptionalUser) (*entity.PostDetailWithHashtagAndIsFavorite, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, *entity.UserRelationFlgMap, error) {
	idRelationFlgMap := &entity.UserRelationFlgMap{}

	post, err := s.PostQueryService.ShowByID(id, ouser)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed find post by id for auth")
	}

	if ouser.Authenticated {
		// 認証されている場合Post.Userをfollow, blockしているかフラグを取得
		idRelationFlgMap, err = s.UserQueryService.RelationFlgMaps(ouser.ID, []int{post.UserID})
		if err != nil {
			return nil, nil, nil, nil, errors.Wrap(err, "failed find is doing flg")
		}
	}

	areaCategoriesMap, themeCategoriesMap, err := s.CategoryIDMapFactory.GenerateCategoryIDMap(post.AreaCategoryIDs(), post.ThemeCategoryIDs())
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed gen category map")
	}

	return post, areaCategoriesMap, themeCategoriesMap, idRelationFlgMap, nil
}

func (s *PostQueryScenarioImpl) ShowBySlug(slug string, ouser *entity.OptionalUser) (*entity.PostDetailWithHashtagAndIsFavorite, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, *entity.UserRelationFlgMap, error) {
	idRelationFlgMap := &entity.UserRelationFlgMap{}

	post, err := s.PostQueryService.ShowBySlug(slug, ouser)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed find post by slug")
	}

	if ouser.Authenticated {
		// 認証されている場合Post.Userをfollow, blockしているかフラグを取得
		idRelationFlgMap, err = s.UserQueryService.RelationFlgMaps(ouser.ID, []int{post.UserID})
		if err != nil {
			return nil, nil, nil, nil, errors.Wrap(err, "failed find is doing flg")
		}
	}

	areaCategoriesMap, themeCategoriesMap, err := s.CategoryIDMapFactory.GenerateCategoryIDMap(post.AreaCategoryIDs(), post.ThemeCategoryIDs())
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed gen category map")
	}

	return post, areaCategoriesMap, themeCategoriesMap, idRelationFlgMap, nil
}

func (s *PostQueryScenarioImpl) ListByParams(query *query.FindPostListQuery, ouser *entity.OptionalUser) (*entity.PostList, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, *entity.UserRelationFlgMap, error) {
	idRelationFlgMap := &entity.UserRelationFlgMap{}

	posts, err := s.PostQueryService.ListByParams(query, ouser)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed find list post by params")
	}

	if ouser.Authenticated {
		// 認証されている場合Post.Userをfollow, blockしているかフラグを取得
		idRelationFlgMap, err = s.UserQueryService.RelationFlgMaps(ouser.ID, posts.UserIDs())
		if err != nil {
			return nil, nil, nil, nil, errors.Wrap(err, "failed find is doing flg")
		}
	}

	areaCategoriesMap, themeCategoriesMap, err := s.CategoryIDMapFactory.GenerateCategoryIDMap(posts.AreaCategoryIDs(), posts.ThemeCategoryIDs())
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed gen category map")
	}

	return posts, areaCategoriesMap, themeCategoriesMap, idRelationFlgMap, nil
}

func (s *PostQueryScenarioImpl) ListFeed(query *query.FindListPaginationQuery, user *entity.User) (*entity.PostList, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, *entity.UserRelationFlgMap, error) {
	posts, err := s.PostQueryService.ListFeed(query, user)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed list feed post")
	}

	// Post.Userをfollow, blockしているかフラグを取得
	idRelationFlgMap, err := s.UserQueryService.RelationFlgMaps(user.ID, posts.UserIDs())
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed find is doing flg")
	}

	areaCategoriesMap, themeCategoriesMap, err := s.CategoryIDMapFactory.GenerateCategoryIDMap(posts.AreaCategoryIDs(), posts.ThemeCategoryIDs())
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed gen category map")
	}

	return posts, areaCategoriesMap, themeCategoriesMap, idRelationFlgMap, nil
}

func (s *PostQueryScenarioImpl) LitFavorite(targetUserID int, query *query.FindListPaginationQuery, ouser *entity.OptionalUser) (*entity.PostList, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, *entity.UserRelationFlgMap, error) {
	idRelationFlgMap := &entity.UserRelationFlgMap{}

	posts, err := s.PostQueryService.ListFavoritePost(targetUserID, query, ouser)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed list favorite post")
	}

	if ouser.Authenticated {
		// 認証されている場合Post.Userをfollow, blockしているかフラグを取得
		idRelationFlgMap, err = s.UserQueryService.RelationFlgMaps(ouser.ID, posts.UserIDs())
		if err != nil {
			return nil, nil, nil, nil, errors.Wrap(err, "failed find is doing flg")
		}
	}

	areaCategoriesMap, themeCategoriesMap, err := s.CategoryIDMapFactory.GenerateCategoryIDMap(posts.AreaCategoryIDs(), posts.ThemeCategoryIDs())
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed gen category map")
	}

	return posts, areaCategoriesMap, themeCategoriesMap, idRelationFlgMap, nil
}
