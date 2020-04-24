package repository

import "github.com/stayway-corp/stayway-media-api/pkg/domain/entity/wordpress"

type (
	WordpressQueryRepository interface {
		FindUsersByIDs(ids []int) ([]*wordpress.User, error)
		FindPostsByIDs(ids []int) ([]*wordpress.Post, error)
		FindLocationsByIDs(ids []int) ([]*wordpress.Location, error)
		FindCategoriesByIDs(ids []int) ([]*wordpress.Category, error)
		FindLocationCategoriesByIDs(ids []int) ([]*wordpress.LocationCategory, error)
		FindComicsByIDs(ids []int) ([]*wordpress.Comic, error)
		FindFeaturesByIDs(ids []int) ([]*wordpress.Feature, error)
		FindVlogsByIDs(ids []int) ([]*wordpress.Vlog, error)
		FindMediaByIDs(ids []int) ([]*wordpress.Media, error)
		FindPostTagsByIDs(ids []int) ([]*wordpress.PostTag, error)
		DownloadAvatar(avatarURL string) (*wordpress.MediaBody, error)
	}
)
