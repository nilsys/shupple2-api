package entity

type (
	Comic struct {
		ID        int `gorm:"primary_key"`
		UserID    int
		Slug      string
		Thumbnail string
		Title     string
		Body      string
		Times
	}

	ComicList struct {
		TotalNumber int
		Comics      []*Comic
	}

	QueryComic struct {
		Comic
		User *User `gorm:"foreignkey:UserID"`
	}

	UserFavoriteComic struct {
		UserID  int
		ComicID int
	}
)

// テーブル名
func (queryComic *QueryComic) TableName() string {
	return "comic"
}

func NewUserFavoriteComic(userID, comicID int) *UserFavoriteComic {
	return &UserFavoriteComic{
		UserID:  userID,
		ComicID: comicID,
	}
}
