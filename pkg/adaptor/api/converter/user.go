package converter

import (
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/command"

	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

/*
 * o -> i
 */
func (c Converters) ConvertListRankinUserParamToQuery(param *input.ListUserRanking) *query.FindUserRankingListQuery {
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

func (c Converters) ConvertListFollowUserParamToQuery(param *input.ListFollowUser) *query.FindFollowUser {
	return &query.FindFollowUser{
		ID:     param.ID,
		Limit:  param.GetLimit(),
		Offset: param.GetOffset(),
	}
}

func (c Converters) ConvertStoreUserParamToEntity(param *input.StoreUser) *entity.User {
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

func (c Converters) ConvertUpdateUserParamToCmd(param *input.UpdateUser) *command.UpdateUser {
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

func (c Converters) ConvertListFavoriteMediaUserToQuery(param *input.ListFavoriteMediaUser) *query.FindListPaginationQuery {
	return &query.FindListPaginationQuery{
		Limit:  param.GetLimit(),
		Offset: param.GetOffset(),
	}
}

/*
 * i -> o
 */
func (c Converters) ConvertUserRankingToOutput(users []*entity.UserDetail) []*output.RankinUser {
	userRanking := make([]*output.RankinUser, len(users))

	for i, user := range users {
		userRanking[i] = c.ConvertUserDetailToOutput(user)
	}

	return userRanking
}

func (c Converters) ConvertUsersToUserSummaryList(users []*entity.User) []*output.UserSummary {
	followUsers := make([]*output.UserSummary, len(users))
	for i, user := range users {
		followUsers[i] = c.convertUserToUserSummary(user)
	}
	return followUsers
}

func (c Converters) convertUserToUserSummary(user *entity.User) *output.UserSummary {
	return &output.UserSummary{
		ID:      user.ID,
		UID:     user.UID,
		Name:    user.Name,
		IconURL: user.AvatarURL(c.filesURL()),
	}
}

// UserDetailをランキング一覧で返す型にコンバート
func (c Converters) ConvertUserDetailToOutput(user *entity.UserDetail) *output.RankinUser {
	interests := make([]string, len(user.Interests))
	for i, interest := range user.Interests {
		interests[i] = interest.Name
	}

	return &output.RankinUser{
		ID:        user.ID,
		UID:       user.UID,
		Name:      user.Name,
		Profile:   user.Profile,
		Thumbnail: user.AvatarURL(c.filesURL()),
		Interests: interests,
	}
}

func (c Converters) ConvertUserDetailWithCountToOutPut(user *entity.UserDetailWithMediaCount) *output.MyPageUser {
	interests := make([]string, len(user.Interests))
	for i, interest := range user.Interests {
		interests[i] = interest.Name
	}

	return &output.MyPageUser{
		ID:             user.ID,
		UID:            user.UID,
		Name:           user.Name,
		Profile:        user.Profile,
		Birthdate:      model.Date(user.Birthdate),
		Email:          user.Email,
		Gender:         user.Gender,
		Icon:           user.AvatarURL(c.filesURL()),
		Header:         user.HeaderURL(c.filesURL()),
		FacebookURL:    user.FacebookURL,
		InstagramURL:   user.InstagramURL,
		TwitterURL:     user.TwitterURL,
		URL:            user.URL,
		LivingArea:     user.LivingArea,
		PostCount:      user.PostCount + user.ReviewCount,
		FollowingCount: user.FollowingCount,
		FollowedCount:  user.FollowerCount,
		Interests:      interests,
	}
}

func (c Converters) NewCreatorFromUser(user *entity.User) output.Creator {
	return output.NewCreator(
		user.ID, user.UID, user.AvatarURL(c.filesURL()), user.Name, user.Profile,
		user.FacebookURL, user.InstagramURL, user.TwitterURL, user.YoutubeURL, user.URL,
	)
}
