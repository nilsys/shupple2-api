package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/repository"
)

type (
	CfProjectCommandRepositoryImpl struct {
		DAO
	}
)

var CfProjectCommandRepositorySet = wire.NewSet(
	wire.Struct(new(CfProjectCommandRepositoryImpl), "*"),
	wire.Bind(new(repository.CfProjectCommandRepository), new(*CfProjectCommandRepositoryImpl)),
)

func (r *CfProjectCommandRepositoryImpl) StoreSupportComment(c context.Context, comment *entity.CfProjectSupportCommentTable) error {
	if err := r.DB(c).Save(comment).Error; err != nil {
		return errors.Wrap(err, "failed store cf_project_support_comment")
	}
	return nil
}

func (r *CfProjectCommandRepositoryImpl) IncrementSupportCommentCount(c context.Context, id int) error {
	if err := r.DB(c).Exec("UPDATE cf_project SET support_comment_count=support_comment_count+1 WHERE id = ?", id).Error; err != nil {
		return errors.Wrap(err, "failed to increment cf_project.support_comment_count")
	}
	return nil
}
