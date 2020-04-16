package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/input"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model/query"
)

// i/oの構造体からレポジトリで使用するクエリ構造体へconvert
func ConvertShowComicListParamToQuery(param *input.ShowComicListParam) *query.FindListPaginationQuery {
	return &query.FindListPaginationQuery{
		Limit:  param.GetLimit(),
		Offset: param.GetOffSet(),
	}
}

// ConvertComicToOutput()のスライスバージョン
func ConvertComicListToOutput(comics *entity.ComicList) output.ComicList {
	responseComics := make([]*output.Comic, len(comics.Comics))

	for i, comic := range comics.Comics {
		responseComics[i] = convertComicToOutput(comic)
	}

	return output.ComicList{
		TotalNumber: comics.TotalNumber,
		Comics:      responseComics,
	}
}

func ConvertQueryComicOutput(queryComic *entity.QueryComic) *output.ShowComic {
	return &output.ShowComic{
		ID:        queryComic.Comic.ID,
		Slug:      queryComic.Comic.Slug,
		Title:     queryComic.Comic.Title,
		Thumbnail: queryComic.Comic.Thumbnail,
		Body:      queryComic.Comic.Body,
		CreatedAt: model.TimeResponse(queryComic.Comic.CreatedAt),
		Creator:   output.NewCreatorFromUser(queryComic.User),
	}
}

// outputの構造体へconvert
func convertComicToOutput(comic *entity.Comic) *output.Comic {
	return &output.Comic{
		ID:        comic.ID,
		Slug:      comic.Slug,
		Title:     comic.Title,
		Thumbnail: comic.Thumbnail,
	}
}
