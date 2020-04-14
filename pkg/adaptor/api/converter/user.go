package converter

import (
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/command"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/response"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

/*
 * o -> i
 */
func ConvertListRankinUserParamToQuery(param *param.ListUserRanking) *query.FindUserRankingListQuery {
	fromDate, _ := model.ParseTimeFromFrontStr(param.FromDate)
	toDate, _ := model.ParseTimeFromFrontStr(param.ToDate)

	return &query.FindUserRankingListQuery{
		AreaID:       param.AreaID,
		SubAreaID:    param.SubAreaID,
		SubSubAreaID: param.SubSubAreaID,
		SortBy:       param.SortBy,
		FromDate:     fromDate,
		ToDate:       toDate,
		Limit:        param.GetLimit(),
		Offset:       param.GetOffset(),
	}
}

func ConvertListFollowUserParamToQuery(param *param.ListFollowUser) *query.FindFollowUser {
	return &query.FindFollowUser{
		ID:     param.ID,
		Limit:  param.GetLimit(),
		Offset: param.GetOffset(),
	}
}

func ConvertStoreUserParamToEntity(param *param.StoreUser) *entity.User {
	interests := make([]*entity.UserInterest, len(param.Interests))
	for i, interest := range param.Interests {
		interests[i] = &entity.UserInterest{InterestID: interest}
	}
	return &entity.User{
		UID:          param.UID,
		Name:         param.Name,
		Email:        param.Email,
		Birthdate:    time.Time(param.BirthDate),
		Gender:       param.Gender,
		Profile:      param.Profile,
		URL:          param.URL,
		FacebookURL:  param.FacebookURL,
		InstagramURL: param.InstagramURL,
		TwitterURL:   param.TwitterURL,
		LivingArea:   param.LivingArea,
		Interests:    interests,
	}
}

func ConvertUpdateUserParamToCmd(param *param.UpdateUser) *command.UpdateUser {
	interests := make([]*entity.UserInterest, len(param.Interests))
	for i, interest := range param.Interests {
		interests[i] = &entity.UserInterest{InterestID: interest}
	}
	return &command.UpdateUser{
		Name:         param.Name,
		Email:        param.Email,
		BirthDate:    param.BirthDate,
		Gender:       param.Gender,
		Profile:      param.Profile,
		IconUUID:     param.IconUUID,
		HeaderUUID:   param.HeaderUUID,
		URL:          param.URL,
		FacebookURL:  param.FacebookURL,
		InstagramURL: param.InstagramURL,
		TwitterURL:   param.TwitterURL,
		LivingArea:   param.LivingArea,
		Interests:    interests,
	}
}

func ConvertListFavoriteMediaUserToQuery(param *param.ListFavoriteMediaUser) *query.FindListPaginationQuery {
	return &query.FindListPaginationQuery{
		Limit:  param.GetLimit(),
		Offset: param.GetOffset(),
	}
}

/*
 * i -> o
 */
func ConvertUserRankingToOutput(users []*entity.UserDetail) []*response.RankinUser {
	userRanking := make([]*response.RankinUser, len(users))

	for i, user := range users {
		userRanking[i] = ConvertUserDetailToOutput(user)
	}

	return userRanking
}

func ConvertUsersToFollowUsers(users []*entity.User) []*response.FollowUser {
	followUsers := make([]*response.FollowUser, len(users))
	for i, user := range users {
		followUsers[i] = convertUserToFollowUser(user)
	}
	return followUsers
}

func convertUserToFollowUser(user *entity.User) *response.FollowUser {
	return &response.FollowUser{
		ID:        user.ID,
		Name:      user.Name,
		Thumbnail: user.IconURL(),
	}
}

// UserDetailをランキング一覧で返す型にコンバート
func ConvertUserDetailToOutput(user *entity.UserDetail) *response.RankinUser {
	interests := make([]string, len(user.Interests))
	for i, interest := range user.Interests {
		interests[i] = interest.Name
	}

	return &response.RankinUser{
		ID:        user.ID,
		Name:      user.Name,
		Profile:   user.Profile,
		Thumbnail: user.IconURL(),
		Interest:  interests,
	}
}

func ConvertUserDetailWithCountToOutPut(user *entity.UserDetailWithCount) *response.MyPageUser {
	interests := make([]string, len(user.Interests))
	for i, interest := range user.Interests {
		interests[i] = interest.Name
	}

	return &response.MyPageUser{
		ID:             user.ID,
		UID:            user.UID,
		Name:           user.Name,
		Profile:        user.Profile,
		Birthdate:      model.Date(user.Birthdate),
		Email:          user.Email,
		Gender:         user.Gender,
		Icon:           user.IconURL(),
		Header:         user.HeaderURL(),
		Facebook:       user.FacebookURL,
		Instagram:      user.InstagramURL,
		Twitter:        user.TwitterURL,
		LivingArea:     user.LivingArea,
		PostCount:      user.PostCount + user.ReviewCount,
		FollowingCount: user.FollowingCount,
		FollowedCount:  user.FollowerCount,
		Interest:       interests,
	}
}
