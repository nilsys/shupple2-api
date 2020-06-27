package service

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	UserQueryService interface {
		ShowByUID(uid string, ouser entity.OptionalUser) (*entity.UserDetailWithMediaCount, error)
		ShowByID(id int) (*entity.UserDetailWithMediaCount, error)
		ShowByMigrationCode(code string) (*entity.UserDetailWithMediaCount, error)
		ShowUserRanking(query *query.FindUserRankingListQuery) ([]*entity.UserDetail, error)
		ListRecommendFollowUser(interestIDs []int) ([]*entity.UserTable, error)
		ListFollowing(query *query.FindFollowUser) ([]*entity.User, error)
		ListFollowed(query *query.FindFollowUser) ([]*entity.User, error)
		ListFavoritePostUser(postID int, user *entity.OptionalUser, query *query.FindListPaginationQuery) ([]*entity.User, error)
		ListFavoriteReviewUser(reviewID int, user *entity.OptionalUser, query *query.FindListPaginationQuery) ([]*entity.User, error)
		IsExistByPhoneNumber(number string) (bool, error)
	}

	UserQueryServiceImpl struct {
		repository.UserQueryRepository
	}
)

var UserQueryServiceSet = wire.NewSet(
	wire.Struct(new(UserQueryServiceImpl), "*"),
	wire.Bind(new(UserQueryService), new(*UserQueryServiceImpl)),
)

func (s *UserQueryServiceImpl) ShowByUID(uid string, ouser entity.OptionalUser) (*entity.UserDetailWithMediaCount, error) {
	userTable, err := s.UserQueryRepository.FindByUID(uid)
	if err != nil {
		return nil, errors.Wrap(err, "failed find user by uid")
	}

	if ouser.Authenticated {
		user, err := s.UserQueryRepository.FindUserDetailWithCountByID(userTable.ID)
		if err != nil {
			return nil, errors.Wrap(err, "failed find user by id")
		}
		user.IsFollow, err = s.UserQueryRepository.IsFollow(userTable.ID, ouser.ID)
		if err != nil {
			return nil, errors.Wrap(err, "failed find user_following")
		}

		return user, nil
	}

	return s.UserQueryRepository.FindUserDetailWithCountByID(userTable.ID)
}

func (s *UserQueryServiceImpl) ShowByID(id int) (*entity.UserDetailWithMediaCount, error) {
	return s.UserQueryRepository.FindUserDetailWithCountByID(id)
}

func (s *UserQueryServiceImpl) ShowByMigrationCode(code string) (*entity.UserDetailWithMediaCount, error) {
	userTable, err := s.UserQueryRepository.FindByMigrationCode(code)
	if err != nil {
		return nil, errors.Wrap(err, "failed find user by migration_code")
	}
	return s.UserQueryRepository.FindUserDetailWithCountByID(userTable.ID)
}

func (s *UserQueryServiceImpl) ShowUserRanking(query *query.FindUserRankingListQuery) ([]*entity.UserDetail, error) {
	return s.UserQueryRepository.FindUserRankingListByParams(query)
}

func (s *UserQueryServiceImpl) ListRecommendFollowUser(interestIDs []int) ([]*entity.UserTable, error) {
	return s.UserQueryRepository.FindRecommendFollowUserList(interestIDs)
}

func (s *UserQueryServiceImpl) ListFollowing(query *query.FindFollowUser) ([]*entity.User, error) {
	return s.UserQueryRepository.FindFollowingByID(query)
}

func (s *UserQueryServiceImpl) ListFollowed(query *query.FindFollowUser) ([]*entity.User, error) {
	return s.UserQueryRepository.FindFollowedByID(query)
}

func (s *UserQueryServiceImpl) ListFavoritePostUser(postID int, user *entity.OptionalUser, query *query.FindListPaginationQuery) ([]*entity.User, error) {
	if user.IsAuthorized() {
		return s.UserQueryRepository.FindFavoritePostUserByUserID(postID, user.ID, query)
	}

	return s.UserQueryRepository.FindFavoritePostUser(postID, query)
}

func (s *UserQueryServiceImpl) ListFavoriteReviewUser(reviewID int, user *entity.OptionalUser, query *query.FindListPaginationQuery) ([]*entity.User, error) {
	if user.IsAuthorized() {
		return s.UserQueryRepository.FindFavoriteReviewUserByUserID(reviewID, user.ID, query)
	}

	return s.UserQueryRepository.FindFavoriteReviewUser(reviewID, query)
}

func (s *UserQueryServiceImpl) IsExistByPhoneNumber(number string) (bool, error) {
	cognitoUser, err := s.UserQueryRepository.FindConfirmedUserTypeByPhoneNumberFromCognito(number)
	if err != nil {
		return false, errors.Wrap(err, "failed find from cognito")
	}
	if cognitoUser == nil {
		return false, nil
	}

	// stayway側に登録されているか
	return s.UserQueryRepository.IsExistByCognitoUserName(*cognitoUser.Username)
}
