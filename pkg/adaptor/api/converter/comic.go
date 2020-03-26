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
	responseComics := make([]*response.Comic, len(comics))

	for i, comic := range comics {
		responseComics[i] = convertComicToOutput(comic)
	}

	return responseComics
}

func ConvertQueryComicOutput(queryComic *entity.QueryComic) *response.ShowComic {
	return &response.ShowComic{
		ID:        queryComic.Comic.ID,
		Slug:      queryComic.Comic.Slug,
		Title:     queryComic.Comic.Title,
		Thumbnail: queryComic.Comic.Thumbnail,
		Body:      queryComic.Comic.Body,
		CreatedAt: model.TimeResponse(queryComic.Comic.CreatedAt),
		Creator:   response.NewCreator(queryComic.User.ID, queryComic.User.Name, queryComic.User.GenerateThumbnailURL(), queryComic.User.Name, queryComic.User.Profile),
	}
}

// outputの構造体へconvert
func convertComicToOutput(comic *entity.Comic) *response.Comic {
	return &response.Comic{
		ID:        comic.ID,
		Slug:      comic.Slug,
		Title:     comic.Title,
		Thumbnail: comic.Thumbnail,
	}
}
