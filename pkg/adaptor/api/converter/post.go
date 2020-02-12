package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/response"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

// i/oの構造体からレポジトリで使用するクエリ構造体へconvert
func ConvertFindPostListParamToQuery(param *param.ShowPostListParam) *query.FindPostListQuery {
	sortBy, _ := model.ParseSortBy(param.SortBy)
	return &query.FindPostListQuery{
		AreaID:       param.AreaID,
		SubAreaID:    param.SubAreaID,
		SubSubAreaID: param.SubSubAreaID,
		ThemeID:      param.ThemeID,
		HashTag:      param.HashTag,
		SortBy:       sortBy,
		Limit:        param.GetLimit(),
		OffSet:       param.GetOffSet(),
	}
}

// outputの構造体へconvert
func convertPostToOutput(postDetail *entity.Post) *response.Post {
	var areaCategories []response.Category
	var themeCategories []response.Category

	for _, category := range postDetail.Categories {
		if category.Type == model.CategoryTypeArea || category.Type == model.CategoryTypeSubArea || category.Type == model.CategoryTypeSubSubArea {
			areaCategories = append(areaCategories, response.NewCategory(category.ID, category.Name))
		}
		if category.Type == model.CategoryTypeTheme {
			themeCategories = append(themeCategories, response.NewCategory(category.ID, category.Name))
		}
	}

	return &response.Post{
		ID:              postDetail.ID,
		Thumbnail:       postDetail.GenerateThumbnailURL(),
		AreaCategories:  areaCategories,
		ThemeCategories: themeCategories,
		Title:           postDetail.Title,
		Creator: response.Creator{
			Thumbnail: postDetail.User.GenerateThumbnailURL(),
			Name:      postDetail.User.Name,
		},
		LikeCount: postDetail.FavoriteCount,
		UpdatedAt: model.TimeFmtToFrontStr(postDetail.UpdatedAt),
	}
}

// ConvertPostToOutput()のスライスバージョン
func ConvertPostToOutput(postDetailList []*entity.Post) []*response.Post {
	var responsePosts []*response.Post

	for _, postDetail := range postDetailList {
		responsePosts = append(responsePosts, convertPostToOutput(postDetail))
	}

	return responsePosts
}
