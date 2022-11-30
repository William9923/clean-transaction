package dao

import (
	"context"
)

type TransactionFn func(context.Context) error

type TransactionManager interface {
	WithinTransaction(ctx context.Context, fn TransactionFn) error
	Begin(ctx context.Context) (context.Context, error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}
