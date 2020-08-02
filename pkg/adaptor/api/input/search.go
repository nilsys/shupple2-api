package input

import (
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

// 検索キーワード(単数)
type Keyward struct {
	Value          string               `query:"q" validate:"required"`
	SuggestionType model.SuggestionType `query:"type"`
}

func (keyward Keyward) SearchResult(byKeyward func(string) (*entity.SearchSuggestions, error), byType func(string, model.SuggestionType) (*entity.SearchSuggestions, error)) (*entity.SearchSuggestions, error) {
	if keyward.SuggestionType != 0 {
		return byType(keyward.Value, keyward.SuggestionType)
	}

	return byKeyward(keyward.Value)
}
