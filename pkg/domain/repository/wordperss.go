package repository

import "github.com/stayway-corp/stayway-media-api/pkg/domain/entity/wordpress"

type (
	WordpressQueryRepository interface {
		FindUserByID(id int) (*wordpress.User, error)
		FindPostByID(id int) (*wordpress.Post, error)
		FindLocationByID(id int) (*wordpress.Location, error)
		FindCategoryByID(id int) (*wordpress.Category, error)
		FindLocationCategoryByID(id int) (*wordpress.LocationCategory, error)
		FindComicByID(id int) (*wordpress.Comic, error)
		FindFeatureByID(id int) (*wordpress.Feature, error)
		FindVlogByID(id int) (*wordpress.Vlog, error)
		FindMediaByID(id int) (*wordpress.Media, error)
		FindPostTagsByIDs(ids []int) ([]*wordpress.PostTag, error)
		FindCfProjectByID(id int) (*wordpress.CfProject, error)
		FindCfReturnGiftByID(id int) (*wordpress.CfReturnGift, error)
		FindCfReturnGiftsByCfProjectID(id int) ([]*wordpress.CfReturnGift, error)
		FetchMediaBodyByID(id int) (*wordpress.MediaBody, error)
		FetchResource(url string) (*wordpress.MediaBody, error)
	}
)
