package dao

import (
	"context"

	"github.com/William9923/clean-transaction/internal/data/model"
)

type TransferLogsDAO interface {
	CreateTransferLogs(ctx context.Context, user1, user2 model.User, amount int32) error
}
