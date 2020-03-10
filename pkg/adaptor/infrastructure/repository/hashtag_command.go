package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type HashtagCommandRepositoryImpl struct {
	DAO
}

var HashtagCommandRepositorySet = wire.NewSet(
	wire.Struct(new(HashtagCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.HashtagCommandRepository), new(*HashtagCommandRepositoryImpl)),
)

func (r *HashtagCommandRepositoryImpl) Store(hashtag *entity.Hashtag) error {
	return errors.Wrap(r.DB(context.TODO()).Save(hashtag).Error, "failed to save hashtag")
}

func (r *HashtagCommandRepositoryImpl) IncrementPostCountByPostID(c context.Context, postID int) error {
	const stmt = "UPDATE hashtag h JOIN post_hashtag ph ON h.id = ph.hashtag_id SET post_count = post_count + 1, score = score + 1 WHERE ph.post_id = ?"
	return errors.Wrap(r.DB(c).Exec(stmt, postID).Error, "failed to increment hashtag.post_count")
}

func (r *HashtagCommandRepositoryImpl) DecrementPostCountByPostID(c context.Context, postID int) error {
	const stmt = "UPDATE hashtag h JOIN post_hashtag ph ON h.id = ph.hashtag_id SET post_count = post_count - 1, score = score - 1 WHERE ph.post_id = ?"
	return errors.Wrap(r.DB(c).Exec(stmt, postID).Error, "failed to decrement hashtag.post_count")
}
