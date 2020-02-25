package param

// 検索キーワード(単数)
type Keyward struct {
	Value string `query:"q" validate:"required"`
}
