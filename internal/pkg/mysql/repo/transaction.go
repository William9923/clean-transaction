package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/William9923/clean-transaction/internal/data/dao"
	internal_mysql "github.com/William9923/clean-transaction/internal/pkg/mysql"
)

type transactionManager struct {
	db *sql.DB
}

func TransactionManager() dao.TransactionManager {
	return transactionManager{
		db: internal_mysql.DB(),
	}
}

func (repo transactionManager) Begin(ctx context.Context) (context.Context, error) {
	tx := internal_mysql.ExtractTx(ctx)
	if tx != nil {
		return ctx, errors.New("ctx already had a transaction")
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return ctx, err
	}

	ctxWithTrx := internal_mysql.InjectTx(ctx, tx)
	return ctxWithTrx, nil
}

func (repo transactionManager) Commit(ctx context.Context) error {
	tx := internal_mysql.ExtractTx(ctx)
	if tx == nil {
		return errors.New("missing transaction from context")
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (transactionManager) Rollback(ctx context.Context) error {
	tx := internal_mysql.ExtractTx(ctx)
	if tx == nil {
		return errors.New("missing transaction from context")
	}

	if err := tx.Rollback(); err != nil {
		return err
	}

	return nil
}
