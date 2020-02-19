package entity

type (
	// table: review
	Review struct {
		ID            int            `json:"id" gorm:"column:id"`
		UserID        int            `json:"-" gorm:"column:user_id"`
		TouristSpotID int            `json:"-" gorm:"column:tourist_spot_id"`
		InnID         int            `json:"innId" gorm:"column:inn_id"`
		Score         int            `json:"score" gorm:"column:score"`
		MediaCount    int            `json:"-" gorm:"column:media_count"`
		Body          string         `json:"body" gorm:"column:body"`
		FavoriteCount int            `json:"favoriteCount" gorm:"column:favorite_count"`
		Medias        []*ReviewMedia `json:"media" gorm:"foreignkey:ReviewID"`
	}

	// table: review_media
	ReviewMedia struct {
		UUID string `json:"uuid" gorm:"column:id"`
		Mime int    `json:"mime" gorm:"column:mime_type"`
		// TODO: 仮置き
		URL      string `json:"url" gorm:"column:priority"`
		ReviewID int    `json:"-" gorm:"review_id"`
	}
)
