package transfer

import (
	"context"

	"github.com/William9923/clean-transaction/internal/data/dao"
	"github.com/William9923/clean-transaction/internal/data/model"
)

type DoTransferParam struct {
	FromUserID uint64
	ToUserID   uint64
	Amount     int32
}

type TransferServiceParam struct{}

type TransferService interface {
	Transfer(ctx context.Context, param DoTransferParam) error
}

type transferService struct {
	TransactionManager dao.TransactionManager
	TransferLogsRepo   dao.TransferLogsDAO
	UserRepo           dao.UserDAO
}

func InitTransferService(params TransferServiceParam) TransferService {
	return &transferService{}
}

func (s *transferService) Transfer(ctx context.Context, param DoTransferParam) error {

	var needRollback bool = false

	if err := s.TransactionManager.Begin(ctx); err != nil {
		return err
	}
	defer func() {
		if needRollback {
			s.TransactionManager.Rollback(ctx)
		}
	}()

	users, err := s.UserRepo.GetMultipleUsers(ctx, []uint64{param.FromUserID, param.ToUserID})
	if err != nil {
		needRollback = true
		return err
	}

	var fromUser model.User
	var toUser model.User
	for _, user := range users {
		if user.UserID == param.FromUserID {
			fromUser = user
		}

		if user.UserID == param.ToUserID {
			toUser = user
		}
	}

	err = s.TransferLogsRepo.CreateTransferLogs(ctx, fromUser, toUser, param.Amount)
	if err != nil {
		needRollback = true
		return err
	}

	if err = s.UserRepo.DepositUserBalance(ctx, toUser, param.Amount); err != nil {
		needRollback = true
		return err
	}

	if err = s.UserRepo.WithdrawUserBalance(ctx, fromUser, param.Amount); err != nil {
		needRollback = true
		return err
	}

	if err := s.TransactionManager.Commit(ctx); err != nil {
		needRollback = true
		return err
	}

	return nil
}
