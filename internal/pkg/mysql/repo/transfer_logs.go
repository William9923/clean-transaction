package repo

import (
	"context"
	"database/sql"

	"github.com/William9923/clean-transaction/internal/data/constant"
	"github.com/William9923/clean-transaction/internal/data/dao"
	"github.com/William9923/clean-transaction/internal/data/model"
	internal_mysql "github.com/William9923/clean-transaction/internal/pkg/mysql"
	internal_time "github.com/William9923/clean-transaction/internal/pkg/time"
	"github.com/go-sql-driver/mysql"
)

const (
	queryInsertTransferLog = `
    INSERT INTO transfer_log_tab (
      user_id_1 ,
      user_id_2 ,
      amount,
      created_at ,
      status
    ) 
    VALUES (?, ?, ?, ?, ?)
  `
)

type transferLogRepo struct {
	db *sql.DB
}

func TransferLogRepo() dao.TransferLogsDAO {
	return transferLogRepo{
		db: internal_mysql.DB(),
	}
}

func (repo transferLogRepo) CreateTransferLogs(ctx context.Context, user1 model.User, user2 model.User, amount int32) error {
	tx := internal_mysql.ExtractTx(ctx)

	vals := []interface{}{
		user1.UserID,
		user2.UserID,
		amount,
		internal_time.Now().UnixMilli(),
		constant.TRANSFER_VALID,
	}

	var err error
	if tx == nil {
		_, err = repo.db.ExecContext(ctx, queryInsertTransferLog, vals...)
	} else {
		_, err = tx.ExecContext(ctx, queryInsertTransferLog, vals...)
	}

	if err != nil {
		if errMySQL, ok := err.(*mysql.MySQLError); ok {
			return internal_mysql.GetMysqlSpecificError(int(errMySQL.Number), err)
		}
		return err
	}

	return nil
}
