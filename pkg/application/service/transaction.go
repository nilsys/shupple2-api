package service

import (
	"context"
)

type TransactionService interface {
	Do(f func(ctx context.Context) error) error
}

type TransactionServiceForTest struct{}

func (s TransactionServiceForTest) Do(f func(ctx context.Context) error) error {
	return f(context.Background())
}
