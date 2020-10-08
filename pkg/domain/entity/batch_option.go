package entity

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	BatchOption struct {
		Name  model.BatchOptionName
		Value string
	}
)
