package service

import "context"

type TransactionService interface {
	Do(f func(ctx context.Context) error) error
}
