package entity

type (
	Card struct {
		ID     int `gorm:"primary_key"`
		UserID int
		CardID string
		Times
	}
)

func NewCard(userID int, cardID string) *Card {
	return &Card{
		UserID: userID,
		CardID: cardID,
	}
}
