package input

import (
	"time"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
)

type (
	ShowByUIDParam struct {
		UID string `param:"uid" validate:"required"`
	}

	ShowByIDParam struct {
		ID int `param:"id" validate:"required"`
	}

	ShowByMigrationCodeParam struct {
		MigrationCode string `param:"code" validate:"required"`
	}

	ListUserRanking struct {
		AreaID       int              `query:"areaId"`
		SubAreaID    int              `query:"subAreaId"`
		SubSubAreaID int              `query:"subSubAreaID"`
		SortBy       model.UserSortBy `query:"sortBy"`
		FromDate     model.Date       `query:"fromDate"`
		ToDate       model.Date       `query:"toDate"`
		PerPage      int              `query:"perPage"`
		Page         int              `query:"page"`
	}

	ListRecommendFollowUser struct {
		InterestIDs []int `query:"interestId" validate:"required"`
	}

	ListFollowUser struct {
		ID      int `param:"id" validate:"required"`
		PerPage int `query:"perPage"`
		Page    int `query:"page"`
	}

	StoreUser struct {
		Name          string       `json:"name"`
		CognitoToken  string       `json:"cognitoToken"`
		MigrationCode *string      `json:"migrationCode"`
		UID           string       `json:"uid"`
		Email         string       `json:"email"`
		BirthDate     model.Date   `json:"birthDate"`
		Gender        model.Gender `json:"gender"`
		Profile       string       `json:"profile"`
		URL           string       `json:"url"`
		FacebookURL   string       `json:"facebookUrl"`
		InstagramURL  string       `json:"instagramUrl"`
		TwitterURL    string       `json:"twitterUrl"`
		YoutubeURL    string       `json:"youtubeUrl"`
		LivingArea    string       `json:"livingArea"`
		Interests     []int        `json:"interests"`
	}

	UpdateUser struct {
		Name         string       `json:"name"`
		Email        string       `json:"email"`
		BirthDate    model.Date   `json:"birthDate"`
		Gender       model.Gender `json:"gender"`
		Profile      string       `json:"profile"`
		IconUUID     string       `json:"iconUuid"`
		HeaderUUID   string       `json:"headerUuid"`
		URL          string       `json:"url"`
		FacebookURL  string       `json:"facebookUrl"`
		InstagramURL string       `json:"instagramUrl"`
		TwitterURL   string       `json:"twitterUrl"`
		YoutubeURL   string       `json:"youtubeUrl"`
		LivingArea   string       `json:"livingArea"`
		Interests    []int        `json:"interests"`
	}

	FollowParam struct {
		ID int `param:"id" validate:"required"`
	}

	// 記事、レビューにいいねしているユーザー一覧
	ListFavoriteMediaUser struct {
		MediaID int `param:"id" validate:"required"`
		PerPage int `query:"perPage"`
		Page    int `query:"page"`
	}
)

const getUsersDefaultPerPage = 30

// いずれのクエリも飛んで来なかった場合エラーを返す
func (param *ListUserRanking) Validate() error {
	// いずれのクエリも飛んで来ない場合
	if param.AreaID == 0 && param.SubAreaID == 0 && param.SubSubAreaID == 0 && param.SortBy == 0 {
		return serror.New(nil, serror.CodeInvalidParam, "Invalid show user ranking list input")
	}

	// RECOMMENDの時はdateの指定が必要無い
	if param.SortBy == model.UserSortByRECOMMEND {
		// AreaID,SubAreaID,SubSubAreaIDのいずれか2つ以上指定されている場合
		if (param.AreaID != 0 && param.SubAreaID != 0) || (param.AreaID != 0 && param.SubSubAreaID != 0) || (param.SubAreaID != 0 && param.SubSubAreaID != 0) {
			return serror.New(nil, serror.CodeInvalidParam, "Invalid show user ranking list search target input")
		}
		return nil
	}

	// AreaID,SubAreaID,SubSubAreaIDのいずれか2つ以上指定されている場合
	if (param.AreaID != 0 && param.SubAreaID != 0) || (param.AreaID != 0 && param.SubSubAreaID != 0) || (param.SubAreaID != 0 && param.SubSubAreaID != 0) {
		return serror.New(nil, serror.CodeInvalidParam, "Invalid show user ranking list search target input")
	}

	// RANKINGの時はdateの指定必須
	if time.Time(param.ToDate).IsZero() || time.Time(param.FromDate).IsZero() {
		return serror.New(nil, serror.CodeInvalidParam, "required from&to date when RANKING")
	}

	return nil
}

// PerPageがクエリで飛んで来なかった場合、デフォルト値である10を返す
func (param *ListUserRanking) GetLimit() int {
	if param.PerPage == 0 {
		return getUsersDefaultPerPage
	}
	return param.PerPage
}

// offsetを返す(sqlで使う想定)
func (param *ListUserRanking) GetOffset() int {
	if param.Page == 1 || param.Page == 0 {
		return 0
	}
	return param.GetLimit() * (param.Page - 1)
}

// PerPageがクエリで飛んで来なかった場合、デフォルト値である10を返す
func (param *ListFollowUser) GetLimit() int {
	if param.PerPage == 0 {
		return getUsersDefaultPerPage
	}
	return param.PerPage
}

// offsetを返す(sqlで使う想定)
func (param *ListFollowUser) GetOffset() int {
	if param.Page == 1 || param.Page == 0 {
		return 0
	}
	return param.GetLimit() * (param.Page - 1)
}

func (p *ListFavoriteMediaUser) GetLimit() int {
	if p.PerPage == 0 {
		return getUsersDefaultPerPage
	}
	return p.PerPage
}

func (p *ListFavoriteMediaUser) GetOffset() int {
	if p.Page == 1 || p.Page == 0 {
		return 0
	}
	return p.GetLimit() * (p.Page - 1)
}
