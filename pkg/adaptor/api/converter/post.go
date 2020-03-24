package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/response"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

/*
 * o -> i
 */
func ConvertFindPostListParamToQuery(param *param.ListPostParam) *query.FindPostListQuery {
	return &query.FindPostListQuery{
		UserID:                 param.UserID,
		AreaID:                 param.AreaID,
		SubAreaID:              param.SubAreaID,
		SubSubAreaID:           param.SubSubAreaID,
		MetasearchAreaID:       param.MetasearchAreaID,
		MetasearchSubAreaID:    param.MetasearchSubAreaID,
		MetasearchSubSubAreaID: param.MetasearchSubSubAreaID,
		InnTypeID:              param.InnTypeID,
		InnDiscerningType:      param.InnDiscerningType,
		ThemeID:                param.ThemeID,
		HashTag:                param.HashTag,
		SortBy:                 param.SortBy,
		Keyward:                param.Keyward,
		Limit:                  param.GetLimit(),
		OffSet:                 param.GetOffSet(),
	}
}

func ConvertListFeedPostParamToQuery(param *param.ListFeedPostParam) *query.FindListPaginationQuery {
	return &query.FindListPaginationQuery{
		Limit:  param.GetLimit(),
		Offset: param.GetOffSet(),
	}
}

/*
 * i -> o
 */
// ConvertPostToOutput()のスライスバージョン
func ConvertPostDetailListToOutput(queryPostList *entity.PostDetailList) response.PostList {
	responsePosts := make([]*response.Post, len(queryPostList.Posts))

	for i, queryPost := range queryPostList.Posts {
		responsePosts[i] = ConvertQueryPostToOutput(queryPost)
	}

	return response.PostList{
		TotalNumber: queryPostList.TotalNumber,
		Posts:       responsePosts,
	}
}

func ConvertPostListToOutput(posts []*entity.PostDetail) []*response.Post {
	responsePosts := make([]*response.Post, len(posts))

	for i, post := range posts {
		responsePosts[i] = ConvertQueryPostToOutput(post)
	}

	return responsePosts
}

// outputの構造体へconvert
func ConvertQueryPostToOutput(queryPost *entity.PostDetail) *response.Post {
	var areaCategories = make([]response.Category, 0)
	var themeCategories = make([]response.Category, 0)

	for _, category := range queryPost.Categories {
		if category.Type.IsAreaKind() {
			areaCategories = append(areaCategories, response.NewCategory(category.ID, category.Name, category.Type))
		} else {
			themeCategories = append(themeCategories, response.NewCategory(category.ID, category.Name, category.Type))
		}
	}

	return &response.Post{
		ID:              queryPost.ID,
		Thumbnail:       queryPost.Thumbnail,
		Title:           queryPost.Title,
		AreaCategories:  areaCategories,
		ThemeCategories: themeCategories,
		Creator: response.Creator{
			Thumbnail: queryPost.User.GenerateThumbnailURL(),
			Name:      queryPost.User.Name,
		},
		LikeCount: queryPost.FavoriteCount,
		Views:     queryPost.Views,
		CreatedAt: model.TimeResponse(queryPost.CreatedAt),
		UpdatedAt: model.TimeResponse(queryPost.UpdatedAt),
	}
}

func ConvertPostDetailWithHashtagToOutput(post *entity.PostDetailWithHashtag) *response.PostShow {
	var areaCategories = make([]response.Category, 0)
	var themeCategories = make([]response.Category, 0)
	var hashtags = make([]response.Hashtag, len(post.Hashtag))
	var bodies = make([]response.PostBody, len(post.Bodies))

	for _, category := range post.Categories {
		if category.Type.IsAreaKind() {
			areaCategories = append(areaCategories, response.NewCategory(category.ID, category.Name, category.Type))
		} else {
			themeCategories = append(themeCategories, response.NewCategory(category.ID, category.Name, category.Type))
		}
	}
	for i, hashtag := range post.Hashtag {
		hashtags[i] = response.NewHashtag(hashtag.ID, hashtag.Name)
	}
	for i, body := range post.Bodies {
		bodies[i] = response.NewPostBody(body.Page, body.Body)
	}

	return &response.PostShow{
		ID:              post.ID,
		Thumbnail:       post.PostTiny.Thumbnail,
		Title:           post.Title,
		Body:            bodies,
		TOC:             post.TOC,
		FavoriteCount:   post.FavoriteCount,
		FacebookCount:   post.FacebookCount,
		TwitterCount:    post.TwitterCount,
		Views:           post.Views,
		Creator:         response.NewCreator(post.User.ID, post.User.GenerateThumbnailURL(), post.User.Name, post.User.Profile),
		AreaCategories:  areaCategories,
		ThemeCategories: themeCategories,
		Hashtags:        hashtags,
		CreatedAt:       model.TimeFmtToFrontStr(post.CreatedAt),
		UpdatedAt:       model.TimeFmtToFrontStr(post.UpdatedAt),
	}
}
