package repository

import "github.com/uma-co82/shupple2-api/pkg/domain/entity"

type (
	UserQueryRepository interface {
		FindByFirebaseID(id string) (entity.UserTiny, error)
	}
)
