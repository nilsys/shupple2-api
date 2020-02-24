package repository

import "github.com/stayway-corp/stayway-media-api/pkg/domain/entity/wordpress"

type (
	WordpressQueryRepository interface {
		FindUsersByIDs(ids []int) ([]*wordpress.User, error)
		FindPostsByIDs(ids []int) ([]*wordpress.Post, error)
		FindLocationsByIDs(ids []int) ([]*wordpress.Location, error)
		FindCategoriesByIDs(ids []int) ([]*wordpress.Category, error)
		FindCategoriesByParentID(parentID, limit int) ([]*wordpress.Category, error)
		FindLocationCategoriesByIDs(ids []int) ([]*wordpress.LocationCategory, error)
		FindComicsByIDs(ids []int) ([]*wordpress.Comic, error)
		FindFeaturesByIDs(ids []int) ([]*wordpress.Feature, error)
		FindVlogsByIDs(ids []int) ([]*wordpress.Vlog, error)
		DownloadAvatar(avatarURL string) ([]byte, error)
	}
)
