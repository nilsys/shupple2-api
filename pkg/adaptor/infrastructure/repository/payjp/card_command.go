package payjp

import (
	"github.com/google/wire"
	"github.com/payjp/payjp-go/v1"
	payjp2 "github.com/stayway-corp/stayway-media-api/pkg/domain/repository/payjp"
)

type CardCommandRepositoryImpl struct {
	PayjpClient *payjp.Service
}

var CardCommandRepositorySet = wire.NewSet(
	wire.Struct(new(CardCommandRepositoryImpl), "*"),
	wire.Bind(new(payjp2.CardCommandRepository), new(*CardCommandRepositoryImpl)),
)

func (r *CardCommandRepositoryImpl) Register(customerID, cardToken string) (*payjp.CardResponse, error) {
	card, err := r.PayjpClient.Customer.AddCardToken(customerID, cardToken)
	if err != nil {
		return nil, handleError(err, "failed store card")
	}

	return card, nil
}

func (r *CardCommandRepositoryImpl) Delete(customerID, cardID string) error {
	err := r.PayjpClient.Customer.DeleteCard(customerID, cardID)
	if err != nil {
		return handleError(err, "failed delete card")
	}

	return nil
}
