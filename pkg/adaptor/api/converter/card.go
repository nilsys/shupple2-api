package converter

import (
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
)

func (c Converters) ConvertCardToOutput(card *entity.Card) *output.Card {
	return &output.Card{
		ID:      card.ID,
		Last4:   card.Last4,
		Expired: card.Expired,
	}
}
