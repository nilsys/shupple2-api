package model

// queryのORDERで使用する値を返す
// デフォルトは新着順を表すNEW
// MEMO: 現状は新着順しかないが、ゆくゆく人気順が追加される
func (sortBy SortBy) GetVlogOrderQuery() string {
	switch sortBy {
	default:
		return "created_at desc"
	}
}
