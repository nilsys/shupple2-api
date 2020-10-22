package repository

import (
	"context"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/uma-co82/shupple2-api/pkg/application/service"
	"github.com/uma-co82/shupple2-api/pkg/domain/model"
)

type TransactionServiceImpl struct {
	DB *gorm.DB
}

var TransactionServiceSet = wire.NewSet(
	wire.Struct(new(TransactionServiceImpl), "*"),
	wire.Bind(new(service.TransactionService), new(*TransactionServiceImpl)),
)

func (s TransactionServiceImpl) Do(f func(context.Context) error) error {
	return Transaction(s.DB, func(tx *gorm.DB) error {
		return f(context.WithValue(context.Background(), model.ContextKeyTransaction, tx))
	})
}
