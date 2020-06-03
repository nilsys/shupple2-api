package entity

type (
	Comic struct {
		ID            int `gorm:"primary_key"`
		UserID        int
		Slug          string
		Thumbnail     string
		Title         string
		Body          string
		FavoriteCount int
		Times
	}

	ComicWithIsFavorite struct {
		Comic
		IsFavorite bool
	}

	ComicList struct {
		TotalNumber int
		Comics      []*ComicWithIsFavorite
	}

	ComicDetail struct {
		ComicWithIsFavorite
		User *User `gorm:"foreignkey:UserID"`
	}

	UserFavoriteComic struct {
		UserID  int
		ComicID int
	}
)

// テーブル名
func (queryComic *ComicDetail) TableName() string {
	return "comic"
}

func NewUserFavoriteComic(userID, comicID int) *UserFavoriteComic {
	return &UserFavoriteComic{
		UserID:  userID,
		ComicID: comicID,
	}
}

func (c *ComicWithIsFavorite) TableName() string {
	return "comic"
}
