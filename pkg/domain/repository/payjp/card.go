package payjp

import (
	"github.com/payjp/payjp-go/v1"
)

type (
	CardCommandRepository interface {
		Register(customerID, cardToken string) (*payjp.CardResponse, error)
		Delete(customerID, cardID string) error
	}

	CardQueryRepository interface {
		Find(customerID, cardID string) (*payjp.CardResponse, error)
	}
)
