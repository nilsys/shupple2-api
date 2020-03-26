package param

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/serror"
)

type (
	ListUserRanking struct {
		AreaID       int              `query:"areaId"`
		SubAreaID    int              `query:"subAreaId"`
		SubSubAreaID int              `query:"subSubAreaID"`
		SortBy       model.UserSortBy `query:"sortBy"`
		FromDate     string           `query:"fromDate" validate:"required"`
		ToDate       string           `query:"toDate" validate:"required"`
		PerPage      int              `query:"perPage"`
		Page         int              `query:"page"`
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
		PinnedPostID  int          `json:"pinnedPostId"`
		Interests     []int        `json:"interests"`
		IconUUID      string       `json:"iconUuid"`
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
	if param.AreaID == 0 && param.SubAreaID == 0 && param.SubSubAreaID == 0 {
		return serror.New(nil, serror.CodeInvalidParam, "Invalid show user ranking list param")
	}

	// AreaID,SubAreaID,SubSubAreaIDのいずれか2つ以上指定されている場合
	if (param.AreaID != 0 && param.SubAreaID != 0) || (param.AreaID != 0 && param.SubSubAreaID != 0) || (param.SubAreaID != 0 && param.SubSubAreaID != 0) {
		return serror.New(nil, serror.CodeInvalidParam, "Invalid show user ranking list search target param")
	}

	if _, err := model.ParseTimeFromFrontStr(param.FromDate); err != nil {
		return serror.New(nil, serror.CodeInvalidParam, "Invalid show user rankin list date format")
	}
	if _, err := model.ParseTimeFromFrontStr(param.ToDate); err != nil {
		return serror.New(nil, serror.CodeInvalidParam, "Invalid show user rankin list date format")
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
	return param.GetLimit()*(param.Page-1) + 1
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
	return param.GetLimit()*(param.Page-1) + 1
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
	return p.GetLimit()*(p.Page-1) + 1
}
