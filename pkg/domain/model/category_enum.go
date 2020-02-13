package model

// categoryのタイプがArea系か判断する
func (categoryType CategoryType) IsAreaKind() bool {
	return categoryType != CategoryTypeTheme
}
