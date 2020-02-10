package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/response"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/dto"
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
func convertPostDetailToOutput(postDetail *dto.PostDetail) *response.Post {
	var areaCategories []string
	var themeCategories []string

	for _, category := range postDetail.Categories {
		if category.Type == model.CategoryTypeArea {
			areaCategories = append(areaCategories, category.Name)
		}
		if category.Type == model.CategoryTypeTheme {
			themeCategories = append(themeCategories, category.Name)
		}
	}

	return &response.Post{
		ID:               postDetail.Post.ID,
		Thumbnail:        postDetail.Post.GenerateThumbnailURL(),
		AreaCategories:   areaCategories,
		ThemeCategories:  themeCategories,
		Title:            postDetail.Post.Title,
		CreatorThumbnail: postDetail.User.GenerateThumbnailURL(),
		CreatorName:      postDetail.User.Name,
		LikeCount:        postDetail.Post.FavoriteCount,
		UpdatedAt:        model.TimeFmtToFrontStr(postDetail.Post.UpdatedAt),
	}
}

// ConvertPostToOutput()のスライスバージョン
func ConvertPostDetailListToOutput(postDetailList []*dto.PostDetail) []*response.Post {
	var responsePosts []*response.Post

	for _, postDetail := range postDetailList {
		responsePosts = append(responsePosts, convertPostDetailToOutput(postDetail))
	}

	return responsePosts
}
