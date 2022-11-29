package dao

import (
	"context"
)

type TransactionManager interface {
	Begin(ctx context.Context) error
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}
