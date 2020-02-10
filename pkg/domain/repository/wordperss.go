package repository

import "github.com/stayway-corp/stayway-media-api/pkg/domain/entity/wordpress"

type (
	WordpressQueryRepository interface {
		FindPostsByIDs(ids []int) ([]*wordpress.Post, error)
		FindLocationsByIDs(ids []int) ([]*wordpress.Location, error)
		FindCategoriesByIDs(ids []int) ([]*wordpress.Category, error)
		FindLocationCategoriesByIDs(ids []int) ([]*wordpress.LocationCategory, error)
		FindComicsByIDs(ids []int) ([]*wordpress.Comic, error)
		FindFeaturesByIDs(ids []int) ([]*wordpress.Feature, error)
		FindVlogsByIDs(ids []int) ([]*wordpress.Vlog, error)
		FindUserByID(id int) (*wordpress.User, error)
		DownloadAvatar(avatarURL string) ([]byte, error)
	}
)
