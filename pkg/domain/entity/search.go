package entity

import "github.com/stayway-corp/stayway-media-api/pkg/domain/model"

type (
	SearchSuggetion struct {
		ID   int                  `json:"id"`
		Type model.SuggestionType `json:"type"`
		Name string               `json:"name"`
	}
)
