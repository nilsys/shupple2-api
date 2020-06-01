package input

// 漫画一覧取得APIパラメータ
type (
	ShowComicListParam struct {
		Page      int `query:"page"`
		PerPage   int `query:"perPage"`
		ExcludeID int `query:"excludeId"`
	}

	ShowComicParam struct {
		ID int `param:"id" validate:"required"`
	}

	FavoriteComicParam struct {
		ComicID int `param:"id" validate:"required"`
	}
)

const showComicListDefaultPerPage = 4

// PerPageがクエリで飛んで来なかった場合、デフォルト値である4を返す
func (showListParam ShowComicListParam) GetLimit() int {
	if showListParam.PerPage == 0 {
		return showComicListDefaultPerPage
	}
	return showListParam.PerPage
}

// offSetを返す(sqlで使う想定)
func (showListParam ShowComicListParam) GetOffSet() int {
	if showListParam.Page == 1 || showListParam.Page == 0 {
		return 0
	}
	return showListParam.GetLimit() * (showListParam.Page - 1)
}
