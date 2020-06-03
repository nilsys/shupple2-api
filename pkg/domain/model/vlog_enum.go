package model

// queryのORDERで使用する値を返す
// デフォルトは新着順を表すNEW
// MEMO: 現状は新着順しかないが、ゆくゆく人気順が追加される
func (sortBy MediaSortBy) GetVlogOrderQuery() string {
	switch sortBy {
	default:
		return "vlog.created_at desc"
	}
}
