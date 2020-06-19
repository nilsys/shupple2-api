package payjp

import (
	"github.com/google/wire"
	"github.com/payjp/payjp-go/v1"
	"github.com/pkg/errors"
	payjp2 "github.com/stayway-corp/stayway-media-api/pkg/domain/repository/payjp"
)

type CardQueryRepositoryImpl struct {
	PayjpClient *payjp.Service
}

var CardQueryRepositorySet = wire.NewSet(
	wire.Struct(new(CardQueryRepositoryImpl), "*"),
	wire.Bind(new(payjp2.CardQueryRepository), new(*CardQueryRepositoryImpl)),
)

func (r *CardQueryRepositoryImpl) Find(customerID, cardID string) (*payjp.CardResponse, error) {
	card, err := r.PayjpClient.Customer.GetCard(customerID, cardID)
	if err != nil {
		return nil, errors.Wrap(err, "failed retrieve card")
	}

	return card, nil
}
