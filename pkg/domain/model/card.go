package model

import "fmt"

// 04/2024のカードの有効期限フォーマット
// pay.jpのレスポンスで使用する事を想定
func CardExpiredFromMonthAndYear(month, year int) string {
	return fmt.Sprintf("%d/%d", month, year)
}
