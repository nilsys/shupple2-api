package payjp

import (
	"github.com/google/wire"
	"github.com/payjp/payjp-go/v1"
	"github.com/pkg/errors"
	payjp2 "github.com/stayway-corp/stayway-media-api/pkg/domain/repository/payjp"
	"github.com/stayway-corp/stayway-media-api/pkg/util"
)

type CardQueryRepositoryImpl struct {
	PayjpClient *payjp.Service
}

var CardQueryRepositorySet = wire.NewSet(
	wire.Struct(new(CardQueryRepositoryImpl), "*"),
	wire.Bind(new(payjp2.CardQueryRepository), new(*CardQueryRepositoryImpl)),
)

const (
	payjpMaximumLimit = 100
)

func (r *CardQueryRepositoryImpl) Find(customerID, cardID string) (*payjp.CardResponse, error) {
	card, err := r.PayjpClient.Customer.GetCard(customerID, cardID)
	if err != nil {
		return nil, errors.Wrap(err, "failed retrieve card from payjp")
	}

	return card, nil
}

// MEMO: 取得したいcardIDが全て取得出来るまでリクエストを繰り返す
//       1枚でも取得できない場合はNotFoundを返す
func (r *CardQueryRepositoryImpl) FindList(customerID string, cardIDs []string) ([]*payjp.CardResponse, error) {
	rows := make([]*payjp.CardResponse, 0, len(cardIDs))
	var offset int

	for {
		cards, hasNext, err := r.PayjpClient.Customer.ListCard(customerID).Offset(offset).Limit(payjpMaximumLimit).Do()
		if err != nil {
			return nil, errors.Wrap(err, "failed list card from payjp")
		}
		for _, card := range cards {
			if util.ContainsFromStrSlice(cardIDs, card.ID) {
				rows = append(rows, card)
			}
		}
		if !hasNext {
			break
		}
		offset += payjpMaximumLimit
	}

	return rows, nil
}
