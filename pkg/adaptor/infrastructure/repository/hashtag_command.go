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

// hashtag.nameで検索、無ければInsert、あれば返す
func (r *HashtagCommandRepositoryImpl) FirstOrCreate(hashtag *entity.Hashtag) (*entity.Hashtag, error) {
	var row entity.Hashtag

	if err := r.DB(context.TODO()).Where("name = ?", hashtag.Name).Attrs(hashtag).FirstOrCreate(&row).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find or create hashtag")
	}
	return &row, nil
}

// hashtag_idとcategory_idの組み合わせが無ければInsert、あれば何もしない
func (r *HashtagCommandRepositoryImpl) StoreHashtagCategory(c context.Context, hashtagCategory *entity.HashtagCategory) error {
	if err := r.DB(c).Where(hashtagCategory).Attrs(hashtagCategory).FirstOrCreate(hashtagCategory).Error; err != nil {
		return errors.Wrap(err, "failed to find or create hashtag_category")
	}
	return nil
}

func (r HashtagCommandRepositoryImpl) IncrementScoreByID(c context.Context, id int) error {
	if err := r.DB(c).Exec("UPDATE hashtag SET score=score+1 WHERE id = ?", id).Error; err != nil {
		return errors.Wrap(err, "failed increment review score")
	}
	return nil
}

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
