package payjp

import (
	"github.com/google/wire"
	"github.com/payjp/payjp-go/v1"
	"github.com/pkg/errors"
	payjp2 "github.com/stayway-corp/stayway-media-api/pkg/domain/repository/payjp"
)

type ChargeCommandRepositoryImpl struct {
	PayjpClient *payjp.Service
}

var ChargeCommandRepositorySet = wire.NewSet(
	wire.Struct(new(ChargeCommandRepositoryImpl), "*"),
	wire.Bind(new(payjp2.ChargeCommandRepository), new(*ChargeCommandRepositoryImpl)),
)

func (r *ChargeCommandRepositoryImpl) Create(customerID string, cardID string, amount int) (*payjp.ChargeResponse, error) {
	charge, err := r.PayjpClient.Charge.Create(amount, payjp.Charge{
		Currency:       "jpy",
		CustomerID:     customerID,
		CustomerCardID: cardID,
		Capture:        false,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed create pay")
	}

	return charge, nil
}

func (r *ChargeCommandRepositoryImpl) Capture(chargeID string) error {
	_, err := r.PayjpClient.Charge.Capture(chargeID)
	if err != nil {
		return errors.Wrap(err, "failed capture pay")
	}

	return nil
}
