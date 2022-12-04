package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/William9923/clean-transaction/internal/conf"
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

func (repo transactionManager) WithinTransaction(ctx context.Context, fn dao.TransactionFn) error {
	var needRollback bool = false

	ctxWithTrx, err := repo.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if needRollback {
			repo.Rollback(ctxWithTrx)
		}
	}()

	if err := fn(ctxWithTrx); err != nil {
		needRollback = true
		return err
	}

	if err := repo.Commit(ctxWithTrx); err != nil {
		needRollback = true
		return err
	}

	return nil
}

func (repo transactionManager) Begin(ctx context.Context) (context.Context, error) {
	tx := internal_mysql.ExtractTx(ctx)
	if tx != nil {
		return ctx, errors.New("ctx already had a transaction")
	}

	isolationLevel, err := getSqlTrxLevel(conf.GetConfig().MySQL.IsolationLevel)
	if err != nil {
		return ctx, err
	}

	opts := &sql.TxOptions{
		ReadOnly:  false,
		Isolation: isolationLevel,
	}

	tx, err = repo.db.BeginTx(ctx, opts)
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

func getSqlTrxLevel(level string) (sql.IsolationLevel, error) {
	switch level {
	case "default", "repeatable-read": // NOTE: default of mysql -> repeatable-read
		return sql.LevelRepeatableRead, nil
	case "read-committed":
		return sql.LevelReadCommitted, nil
	case "read-uncommitted":
		return sql.LevelReadUncommitted, nil
	case "serializable":
		return sql.LevelSerializable, nil
	default:
		return sql.LevelDefault, errors.New("unrecognized mysql trx level")
	}
}
