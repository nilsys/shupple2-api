package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type (
	HashtagQueryRepository interface {
		FindByNames(names []string) (map[string]*entity.Hashtag, error)
		FindRecommendList(areaID, subAreaID, subSubAreaID, limit int) ([]*entity.Hashtag, error)
		// name部分一致検索
		SearchByName(name string) ([]*entity.Hashtag, error)
	}

	HashtagCommandRepository interface {
		FirstOrCreate(hashtag *entity.Hashtag) (*entity.Hashtag, error)
		IncrementScoreByID(c context.Context, id int) error
		Store(hashtag *entity.Hashtag) error
		IncrementPostCountByPostID(c context.Context, postID int) error
		DecrementPostCountByPostID(c context.Context, postID int) error
		StoreHashtagFollow(following *entity.UserFollowHashtag) error
		DeleteHashtagFollow(userID, hashtagID int) error
	}
)
