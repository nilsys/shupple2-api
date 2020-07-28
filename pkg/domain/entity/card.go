package entity

type (
	// CardテーブルはDeletedAtを持つが、gormのUnscoped()が関係参照に適用されない為deleted_atは明示しない
	Card struct {
		ID      int `gorm:"primary_key"`
		UserID  int
		CardID  string
		Last4   string
		Expired string
		TimesWithoutDeletedAt
	}
)

func NewCard(userID int, cardID, last4, expired string) *Card {
	return &Card{
		UserID:  userID,
		CardID:  cardID,
		Last4:   last4,
		Expired: expired,
	}
}
