package payjp

import (
	"github.com/google/wire"
	"github.com/payjp/payjp-go/v1"
	"github.com/pkg/errors"
	payjp2 "github.com/stayway-corp/stayway-media-api/pkg/domain/repository/payjp"
)

type CustomerQueryRepositoryImpl struct {
	PayjpClient *payjp.Service
}

var CustomerQueryRepositorySet = wire.NewSet(
	wire.Struct(new(CustomerQueryRepositoryImpl), "*"),
	wire.Bind(new(payjp2.CustomerQueryRepository), new(*CustomerQueryRepositoryImpl)),
)

func (r *CustomerQueryRepositoryImpl) FindCustomer(customerID string) (*payjp.CustomerResponse, error) {
	customer, err := r.PayjpClient.Customer.Retrieve(customerID)
	if err != nil {
		return nil, errors.Wrap(err, "failed retrieve customer")
	}
	return customer, nil
}
