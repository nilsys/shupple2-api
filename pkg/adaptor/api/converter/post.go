package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

/*
 * o -> i
 */
func ConvertFindPostListParamToQuery(param *input.ListPostParam) *query.FindPostListQuery {
	return &query.FindPostListQuery{
		UserID:                 param.UserID,
		AreaID:                 param.AreaID,
		SubAreaID:              param.SubAreaID,
		SubSubAreaID:           param.SubSubAreaID,
		ChildAreaID:            param.ChildAreaID,
		ChildSubAreaID:         param.ChildSubAreaID,
		ChildSubSubAreaID:      param.ChildSubSubAreaID,
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

func ConvertListFeedPostParamToQuery(param *input.ListFeedPostParam) *query.FindListPaginationQuery {
	return &query.FindListPaginationQuery{
		Limit:  param.GetLimit(),
		Offset: param.GetOffSet(),
	}
}

/*
 * i -> o
 */
func ConvertPostDetailListToOutput(posts []*entity.PostDetail) []*output.Post {
	responsePosts := make([]*output.Post, len(posts))

	for i, queryPost := range posts {
		responsePosts[i] = ConvertQueryPostToOutput(queryPost)
	}

	return responsePosts
}

func ConvertQueryPostToOutput(queryPost *entity.PostDetail) *output.Post {
	return &output.Post{
		ID:              queryPost.ID,
		Thumbnail:       queryPost.Thumbnail,
		Title:           queryPost.Title,
		AreaCategories:  ConvertAreaCategoriesToOutput(queryPost.AreaCategories),
		ThemeCategories: ConvertThemeCategoriesToOutput(queryPost.ThemeCategories),
		Creator:         output.NewCreatorFromUser(queryPost.User),
		FavoriteCount:   queryPost.FavoriteCount,
		Views:           queryPost.Views,
		HideAds:         queryPost.HideAds,
		CreatedAt:       model.TimeResponse(queryPost.CreatedAt),
		UpdatedAt:       model.TimeResponse(queryPost.UpdatedAt),
	}
}

func ConvertPostDetailWithHashtagAndIsFavoriteToOutput(post *entity.PostDetailWithHashtagAndIsFavorite) *output.PostShow {
	var hashtags = make([]*output.Hashtag, len(post.Hashtag))
	var bodies = make([]*output.PostBody, len(post.Bodies))

	for i, hashtag := range post.Hashtag {
		hashtags[i] = output.NewHashtag(hashtag.ID, hashtag.Name)
	}
	for i, body := range post.Bodies {
		bodies[i] = output.NewPostBody(body.Page, body.Body)
	}

	return &output.PostShow{
		ID:              post.ID,
		Thumbnail:       post.PostTiny.Thumbnail,
		Title:           post.Title,
		Body:            bodies,
		TOC:             post.TOC,
		FavoriteCount:   post.FavoriteCount,
		FacebookCount:   post.FacebookCount,
		TwitterCount:    post.TwitterCount,
		Views:           post.Views,
		SEOTitle:        post.SEOTitle,
		SEODescription:  post.SEODescription,
		HideAds:         post.HideAds,
		IsFavorited:     post.IsFavorite,
		Creator:         output.NewCreatorFromUser(post.User),
		AreaCategories:  ConvertAreaCategoriesToOutput(post.AreaCategories),
		ThemeCategories: ConvertThemeCategoriesToOutput(post.ThemeCategories),
		Hashtags:        hashtags,
		CreatedAt:       model.TimeFmtToFrontStr(post.CreatedAt),
		UpdatedAt:       model.TimeFmtToFrontStr(post.UpdatedAt),
	}
}

func ConvertPostDetailWithHashtagToOutput(post *entity.PostDetailWithHashtag) *output.PostShow {
	var hashtags = make([]*output.Hashtag, len(post.Hashtag))
	var bodies = make([]*output.PostBody, len(post.Bodies))

	for i, hashtag := range post.Hashtag {
		hashtags[i] = output.NewHashtag(hashtag.ID, hashtag.Name)
	}
	for i, body := range post.Bodies {
		bodies[i] = output.NewPostBody(body.Page, body.Body)
	}

	return &output.PostShow{
		ID:              post.ID,
		Thumbnail:       post.PostTiny.Thumbnail,
		Title:           post.Title,
		Body:            bodies,
		TOC:             post.TOC,
		FavoriteCount:   post.FavoriteCount,
		FacebookCount:   post.FacebookCount,
		TwitterCount:    post.TwitterCount,
		Views:           post.Views,
		SEOTitle:        post.SEOTitle,
		SEODescription:  post.SEODescription,
		HideAds:         post.HideAds,
		Creator:         output.NewCreatorFromUser(post.User),
		AreaCategories:  ConvertAreaCategoriesToOutput(post.AreaCategories),
		ThemeCategories: ConvertThemeCategoriesToOutput(post.ThemeCategories),
		Hashtags:        hashtags,
		CreatedAt:       model.TimeFmtToFrontStr(post.CreatedAt),
		UpdatedAt:       model.TimeFmtToFrontStr(post.UpdatedAt),
	}
}

func ConvertPostListToOutput(list *entity.PostList) *output.PostList {
	posts := ConvertPostListTiniesToOutput(list.Posts)

	return &output.PostList{
		TotalNumber: list.TotalNumber,
		Posts:       posts,
	}
}

func ConvertPostListTiniesToOutput(list []*entity.PostListTiny) []*output.Post {
	res := make([]*output.Post, len(list))

	for i, tiny := range list {
		res[i] = ConvertPostListTinyToOutput(tiny)
	}

	return res
}

func ConvertPostListTinyToOutput(post *entity.PostListTiny) *output.Post {
	return &output.Post{
		ID:              post.ID,
		Thumbnail:       post.Thumbnail,
		Title:           post.Title,
		AreaCategories:  ConvertAreaCategoriesToOutput(post.AreaCategories),
		ThemeCategories: ConvertThemeCategoriesToOutput(post.ThemeCategories),
		Creator:         output.NewCreatorFromUser(post.User),
		FavoriteCount:   post.FavoriteCount,
		Views:           post.Views,
		HideAds:         post.HideAds,
		IsFavorite:      post.IsFavorite,
		CreatedAt:       model.TimeResponse(post.CreatedAt),
		UpdatedAt:       model.TimeResponse(post.UpdatedAt),
	}
}
