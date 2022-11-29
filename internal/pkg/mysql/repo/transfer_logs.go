package repo

import (
	"context"
	"database/sql"

	"github.com/William9923/clean-transaction/internal/data/dao"
	"github.com/William9923/clean-transaction/internal/data/model"
)

type transferLogRepo struct {
	db *sql.DB
}

func TransferLogRepo() dao.TransferLogsDAO {
	return transferLogRepo{} // TODO: inject with MySQL dependencies
}

// CreateTransferLogs implements dao.TransferLogsDAO
func (transferLogRepo) CreateTransferLogs(ctx context.Context, user1 model.User, user2 model.User, amount int32) error {
	panic("unimplemented")
}
