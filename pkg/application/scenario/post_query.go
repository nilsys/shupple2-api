package scenario

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/factory"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	PostQueryScenario interface {
		Show(id int, ouser *entity.OptionalUser) (*entity.PostDetailWithHashtagAndIsFavorite, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, map[int]bool, error)
		ShowBySlug(slug string, ouser *entity.OptionalUser) (*entity.PostDetailWithHashtagAndIsFavorite, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, map[int]bool, error)
		ListByParams(query *query.FindPostListQuery, ouser *entity.OptionalUser) (*entity.PostList, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, map[int]bool, error)
		ListFeed(query *query.FindListPaginationQuery, user *entity.User) (*entity.PostList, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, map[int]bool, error)
		LitFavorite(targetUserID int, query *query.FindListPaginationQuery, ouser *entity.OptionalUser) (*entity.PostList, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, map[int]bool, error)
	}

	PostQueryScenarioImpl struct {
		factory.CategoryIDMapFactory
		service.PostQueryService
		repository.UserQueryRepository
	}
)

var PostQueryScenarioSet = wire.NewSet(
	wire.Struct(new(PostQueryScenarioImpl), "*"),
	wire.Bind(new(PostQueryScenario), new(*PostQueryScenarioImpl)),
)

func (s *PostQueryScenarioImpl) Show(id int, ouser *entity.OptionalUser) (*entity.PostDetailWithHashtagAndIsFavorite, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, map[int]bool, error) {
	var idIsFollowMap map[int]bool

	post, err := s.PostQueryService.ShowByID(id, ouser)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed find post by id for auth")
	}

	if ouser.Authenticated {
		// 認証されている場合Post.Userをfollowしているかフラグを取得
		idIsFollowMap, err = s.UserQueryRepository.IsFollowing(ouser.ID, []int{post.UserID})
		if err != nil {
			return nil, nil, nil, nil, errors.Wrap(err, "failed list user_following")
		}
	}

	areaCategoriesMap, themeCategoriesMap, err := s.CategoryIDMapFactory.GenerateCategoryIDMap(post.AreaCategoryIDs(), post.ThemeCategoryIDs())
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed gen category map")
	}

	return post, areaCategoriesMap, themeCategoriesMap, idIsFollowMap, nil
}

func (s *PostQueryScenarioImpl) ShowBySlug(slug string, ouser *entity.OptionalUser) (*entity.PostDetailWithHashtagAndIsFavorite, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, map[int]bool, error) {
	var idIsFollowMap map[int]bool

	post, err := s.PostQueryService.ShowBySlug(slug, ouser)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed find post by slug")
	}

	if ouser.Authenticated {
		// 認証されている場合Post.Userをfollowしているかフラグを取得
		idIsFollowMap, err = s.UserQueryRepository.IsFollowing(ouser.ID, []int{post.UserID})
		if err != nil {
			return nil, nil, nil, nil, errors.Wrap(err, "failed list user_following")
		}
	}

	areaCategoriesMap, themeCategoriesMap, err := s.CategoryIDMapFactory.GenerateCategoryIDMap(post.AreaCategoryIDs(), post.ThemeCategoryIDs())
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed gen category map")
	}

	return post, areaCategoriesMap, themeCategoriesMap, idIsFollowMap, nil
}

func (s *PostQueryScenarioImpl) ListByParams(query *query.FindPostListQuery, ouser *entity.OptionalUser) (*entity.PostList, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, map[int]bool, error) {
	var idIsFollowMap map[int]bool

	posts, err := s.PostQueryService.ListByParams(query, ouser)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed find list post by params")
	}

	if ouser.Authenticated {
		// 認証されている場合Post.Userをfollowしているかフラグを取得
		idIsFollowMap, err = s.UserQueryRepository.IsFollowing(ouser.ID, posts.UserIDs())
		if err != nil {
			return nil, nil, nil, nil, errors.Wrap(err, "failed list user_following")
		}
	}

	areaCategoriesMap, themeCategoriesMap, err := s.CategoryIDMapFactory.GenerateCategoryIDMap(posts.AreaCategoryIDs(), posts.ThemeCategoryIDs())
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed gen category map")
	}

	return posts, areaCategoriesMap, themeCategoriesMap, idIsFollowMap, nil
}

func (s *PostQueryScenarioImpl) ListFeed(query *query.FindListPaginationQuery, user *entity.User) (*entity.PostList, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, map[int]bool, error) {
	var idIsFollowMap map[int]bool

	posts, err := s.PostQueryService.ListFeed(query, user)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed list feed post")
	}

	// 認証されている場合Post.Userをfollowしているかフラグを取得
	idIsFollowMap, err = s.UserQueryRepository.IsFollowing(user.ID, posts.UserIDs())
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed list user_following")
	}

	areaCategoriesMap, themeCategoriesMap, err := s.CategoryIDMapFactory.GenerateCategoryIDMap(posts.AreaCategoryIDs(), posts.ThemeCategoryIDs())
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed gen category map")
	}

	return posts, areaCategoriesMap, themeCategoriesMap, idIsFollowMap, nil
}

func (s *PostQueryScenarioImpl) LitFavorite(targetUserID int, query *query.FindListPaginationQuery, ouser *entity.OptionalUser) (*entity.PostList, map[int]*entity.AreaCategory, map[int]*entity.ThemeCategory, map[int]bool, error) {
	var idIsFollowMap map[int]bool

	posts, err := s.PostQueryService.ListFavoritePost(targetUserID, query, ouser)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed list favorite post")
	}

	if ouser.Authenticated {
		// 認証されている場合Post.Userをfollowしているかフラグを取得
		idIsFollowMap, err = s.UserQueryRepository.IsFollowing(ouser.ID, posts.UserIDs())
		if err != nil {
			return nil, nil, nil, nil, errors.Wrap(err, "failed list user_following")
		}
	}

	areaCategoriesMap, themeCategoriesMap, err := s.CategoryIDMapFactory.GenerateCategoryIDMap(posts.AreaCategoryIDs(), posts.ThemeCategoryIDs())
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed gen category map")
	}

	return posts, areaCategoriesMap, themeCategoriesMap, idIsFollowMap, nil
}
