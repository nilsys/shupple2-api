package main

import (
	"context"
	"log"

	"github.com/payjp/payjp-go/v1"
	"github.com/pkg/errors"
	repository2 "github.com/stayway-corp/stayway-media-api/pkg/adaptor/infrastructure/repository"
	"github.com/stayway-corp/stayway-media-api/pkg/application/service"
	payjp2 "github.com/stayway-corp/stayway-media-api/pkg/domain/repository/payjp"
)

// pay.jp側のcustomer登録は全て外す
// また、その時点でのCard(stayway側db)も全て論理削除する
type (
	Script struct {
		PayjpClient               *payjp.Service
		CustomerCommandRepository payjp2.CustomerCommandRepository
		repository2.DAO
		service.TransactionService
	}
)

const (
	payjpListLimit = 100
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
		return errors.Wrap(err, "failed list all payjp customer")
	}

	return s.TransactionService.Do(func(ctx context.Context) error {
		if err := s.DeleteAllCard(ctx); err != nil {
			return errors.Wrap(err, "failed delete all card")
		}

		if err := s.UnregisterFromPayjpByIDs(customerIDs(customers)); err != nil {
			return errors.Wrap(err, "failed unregister payjp customer")
		}

		return nil
	})
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

func (s Script) UnregisterFromPayjpByIDs(customerIDs []string) error {
	for _, id := range customerIDs {
		if err := s.PayjpClient.Customer.Delete(id); err != nil {
			return errors.Wrap(err, "failed del payjp customer")
		}
	}
	return nil
}

func (s Script) DeleteAllCard(ctx context.Context) error {
	if err := s.DB(ctx).Exec("UPDATE card SET deleted_at = NOW()").Error; err != nil {
		return errors.Wrap(err, "failed update card.deleted_at")
	}
	return nil
}

func customerIDs(customers []*payjp.CustomerResponse) []string {
	resolve := make([]string, len(customers))
	for i, customer := range customers {
		resolve[i] = customer.ID
	}
	return resolve
}
