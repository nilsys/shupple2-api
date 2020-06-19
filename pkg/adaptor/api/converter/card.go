package converter

import (
	"github.com/payjp/payjp-go/v1"
	"github.com/stayway-corp/stayway-media-api/pkg/adaptor/api/output"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/model"
)

func (c Converters) ConvertCardToOutput(card *payjp.CardResponse) *output.Card {
	return &output.Card{
		ID:      card.ID,
		Last4:   card.Last4,
		Expired: model.CardExpiredFromMonthAndYear(card.ExpMonth, card.ExpYear),
	}
}
