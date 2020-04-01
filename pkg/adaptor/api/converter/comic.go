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
func ConvertComicListToOutput(comics *entity.ComicList) response.ComicList {
	responseComics := make([]*response.Comic, len(comics.Comics))

	for i, comic := range comics.Comics {
		responseComics[i] = convertComicToOutput(comic)
	}

	return response.ComicList{
		TotalNumber: comics.TotalNumber,
		Comics:      responseComics,
	}
}

func ConvertQueryComicOutput(queryComic *entity.QueryComic) *response.ShowComic {
	return &response.ShowComic{
		ID:        queryComic.Comic.ID,
		Slug:      queryComic.Comic.Slug,
		Title:     queryComic.Comic.Title,
		Thumbnail: queryComic.Comic.Thumbnail,
		Body:      queryComic.Comic.Body,
		CreatedAt: model.TimeResponse(queryComic.Comic.CreatedAt),
		Creator:   response.NewCreatorFromUser(queryComic.User),
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
