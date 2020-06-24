package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type (
	CfProjectCommandRepository interface {
		StoreSupportComment(c context.Context, comment *entity.CfProjectSupportCommentTable) error
		IncrementSupportCommentCount(c context.Context, id int) error
	}

	CfProjectQueryRepository interface {
		Lock(c context.Context, id int) (*entity.CfProject, error)
	}
)
