package payjp

import (
	"github.com/payjp/payjp-go/v1"
)

type (
	CustomerCommandRepository interface {
		StoreCustomer(customerID string, email string) error
	}

	CustomerQueryRepository interface {
		FindCustomer(customerID string) (*payjp.CustomerResponse, error)
	}
)
