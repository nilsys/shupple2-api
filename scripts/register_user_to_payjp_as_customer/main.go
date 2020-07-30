package main

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/payjp/payjp-go/v1"
	"github.com/pkg/errors"
	"github.com/stayway-corp/stayway-media-api/pkg/domain/entity"
	payjp2 "github.com/stayway-corp/stayway-media-api/pkg/domain/repository/payjp"
)

type (
	Script struct {
		PayjpClient               *payjp.Service
		DB                        *gorm.DB
		CustomerCommandRepository payjp2.CustomerCommandRepository
	}
)

const (
	payjpListLimit = 100
	listUserLimit  = 100
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	script, err := InitializeScript("./config.yaml")
	if err != nil {
		return errors.Wrap(err, "failed init script")
	}
	return script.Run()
}

func (s Script) Run() error {
	customers, err := s.AllPayjpCustomer()
	if err != nil {
		return errors.Wrap(err, "failed list payjp customer")
	}

	customerIDs := s.AllPayjpCustomerIDMap(customers)

	lastID := 0
	for {
		var users []*entity.UserTiny
		if err := s.DB.Where("id > ?", lastID).Limit(listUserLimit).Find(&users).Error; err != nil {
			return errors.Wrap(err, "failed list user")
		}
		if len(users) == 0 {
			break
		}

		for _, user := range users {
			if _, ok := customerIDs[user.PayjpCustomerID()]; ok {
				continue
			}
			if err := s.CustomerCommandRepository.StoreCustomer(user.PayjpCustomerID(), user.Email); err != nil {
				return errors.Wrap(err, "failed create customer to payjp")
			}
		}

		lastID += users[len(users)-1].ID
	}

	return nil
}

func (s Script) AllPayjpCustomer() ([]*payjp.CustomerResponse, error) {
	var customers []*payjp.CustomerResponse
	offset := 0
	for {
		partCustomers, hasMore, err := s.PayjpClient.Customer.List().Offset(offset).Limit(payjpListLimit).Do()
		if err != nil {
			return nil, errors.Wrap(err, "failed list payjp customer")
		}
		customers = append(customers, partCustomers...)
		if !hasMore {
			break
		}
		offset += payjpListLimit
	}
	return customers, nil
}

func (s Script) AllPayjpCustomerIDMap(customers []*payjp.CustomerResponse) map[string]struct{} {
	resolve := make(map[string]struct{}, len(customers))
	for _, customer := range customers {
		resolve[customer.ID] = struct{}{}
	}
	return resolve
}
