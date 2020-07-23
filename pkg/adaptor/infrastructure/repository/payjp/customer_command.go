package payjp

import (
	"github.com/google/wire"
	"github.com/payjp/payjp-go/v1"
	payjp2 "github.com/stayway-corp/stayway-media-api/pkg/domain/repository/payjp"
)

type CustomerCommandRepositoryImpl struct {
	PayjpClient *payjp.Service
}

var CustomerCommandRepositorySet = wire.NewSet(
	wire.Struct(new(CustomerCommandRepositoryImpl), "*"),
	wire.Bind(new(payjp2.CustomerCommandRepository), new(*CustomerCommandRepositoryImpl)),
)

func (r *CustomerCommandRepositoryImpl) StoreCustomer(customerID string, email string) error {
	_, err := r.PayjpClient.Customer.Create(payjp.Customer{
		Email: email,
		ID:    customerID,
	})
	if err != nil {
		return handleError(err, "failed create customer")
	}

	return nil
}
