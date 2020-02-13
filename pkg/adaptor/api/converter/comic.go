package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/param"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/response"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

// i/oの構造体からレポジトリで使用するクエリ構造体へconvert
func ConvertShowComicListParamToQuery(param *param.ShowComicListParam) *query.FindListPaginationQuery {
	return &query.FindListPaginationQuery{
		Limit:  param.GetLimit(),
		Offset: param.GetOffSet(),
	}
}

// ConvertComicToOutput()のスライスバージョン
func ConvertComicListToOutput(comics []*entity.Comic) []*response.Comic {
	// MEMO: 代入しないと0件の時にフロントにnullが返る
	responseComics := []*response.Comic{}

	for _, comic := range comics {
		responseComics = append(responseComics, convertComicToOutput(comic))
	}

	return responseComics
}

func ConvertQueryComicOutput(queryComic *entity.QueryComic) *response.ShowComic {
	return &response.ShowComic{
		ID:        queryComic.Comic.ID,
		Slug:      queryComic.Comic.Slug,
		Title:     queryComic.Comic.Title,
		Thumbnail: queryComic.Comic.GenerateThumbnailURL(),
		Body:      queryComic.Comic.Body,
		CreatedAt: model.TimeFmtToFrontStr(queryComic.Comic.CreatedAt),
		Creator:   response.NewCreator(queryComic.User.GenerateThumbnailURL(), queryComic.User.Name, queryComic.User.Profile),
	}
}

// outputの構造体へconvert
func convertComicToOutput(comic *entity.Comic) *response.Comic {
	return &response.Comic{
		ID:        comic.ID,
		Slug:      comic.Slug,
		Title:     comic.Title,
		Thumbnail: comic.GenerateThumbnailURL(),
	}
}
