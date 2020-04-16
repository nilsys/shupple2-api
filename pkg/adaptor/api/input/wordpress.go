package input

import "github.com/stayway-corp/stayway-media-api/pkg/domain/entity/wordpress"

type ImportWordpressEntityParam struct {
	EntityType wordpress.EntityType `json:"type"`
	ID         int                  `json:"id"`
}
