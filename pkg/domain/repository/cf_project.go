package repository

import (
	"context"

	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

type (
	CfProjectQueryRepository interface {
		LockCfProjectListByIDs(c context.Context, ids []int) (*entity.CfProjectList, error)
	}
)
