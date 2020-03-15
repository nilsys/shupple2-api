package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

// Review参照系レポジトリ実装
type ReviewCommandRepositoryImpl struct {
	DAO
}

var ReviewCommandRepositorySet = wire.NewSet(
	wire.Struct(new(ReviewCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.ReviewCommandRepository), new(*ReviewCommandRepositoryImpl)),
)

// TODO: updateの時にSaveの挙動確認
func (r *ReviewCommandRepositoryImpl) StoreReview(c context.Context, review *entity.Review) error {
	if err := r.DB(c).Save(review).Error; err != nil {
		return errors.Wrap(err, "failed store review")
	}
	return nil
}
