package transfer

import (
	"context"
	"errors"

	"github.com/William9923/clean-transaction/internal/data/dao"
	"github.com/William9923/clean-transaction/internal/data/model"
)

type DoTransferParam struct {
	FromUserID uint64
	ToUserID   uint64
	Amount     int32
}

type TransferServiceParam struct {
	TransactionManager dao.TransactionManager
	TransferLogsRepo   dao.TransferLogsDAO
	UserRepo           dao.UserDAO
}

type TransferService interface {
	Transfer(ctx context.Context, param DoTransferParam) error
	TransferV2(ctx context.Context, param DoTransferParam) error
}

type transferService struct {
	transactionManager dao.TransactionManager
	transferLogsRepo   dao.TransferLogsDAO
	userRepo           dao.UserDAO
}

func InitTransferService(params TransferServiceParam) TransferService {
	return &transferService{
		transactionManager: params.TransactionManager,
		transferLogsRepo:   params.TransferLogsRepo,
		userRepo:           params.UserRepo,
	}
}

func (s *transferService) Transfer(ctx context.Context, param DoTransferParam) error {

	var needRollback bool = false

	ctxWithTrx, err := s.transactionManager.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if needRollback {
			s.transactionManager.Rollback(ctxWithTrx)
		}
	}()

	users, err := s.userRepo.GetUsersInTransfer(ctxWithTrx, [2]uint64{param.FromUserID, param.ToUserID})
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

	err = s.transferLogsRepo.CreateTransferLogs(ctxWithTrx, fromUser, toUser, param.Amount)
	if err != nil {
		needRollback = true
		return err
	}

	if err = s.depositUserBalance(ctxWithTrx, toUser, param.Amount); err != nil {
		needRollback = true
		return err
	}

	if err = s.withdrawUserBalance(ctxWithTrx, fromUser, param.Amount); err != nil {
		needRollback = true
		return err
	}

	if err := s.transactionManager.Commit(ctxWithTrx); err != nil {
		needRollback = true
		return err
	}

	return nil
}

func (s *transferService) TransferV2(ctx context.Context, param DoTransferParam) error {
	return s.transactionManager.WithinTransaction(ctx, func(trxCtx context.Context) error {
		return s.transfer(trxCtx, param)
	})
}

func (s *transferService) transfer(ctx context.Context, param DoTransferParam) error {
	users, err := s.userRepo.GetUsersInTransfer(ctx, [2]uint64{param.FromUserID, param.ToUserID})
	if err != nil {
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

	err = s.transferLogsRepo.CreateTransferLogs(ctx, fromUser, toUser, param.Amount)
	if err != nil {
		return err
	}

	if err = s.depositUserBalance(ctx, toUser, param.Amount); err != nil {
		return err
	}

	if err = s.withdrawUserBalance(ctx, fromUser, param.Amount); err != nil {
		return err
	}
	return nil
}

// NOTE: if there are 2 session, in which user 1 transfer to user 2 & user 2 transfer to user 1, then it is possible to have deadlock with below code
func (s *transferService) transferWithPossibleRace(ctx context.Context, param DoTransferParam) error {

	fromUser, err := s.userRepo.GetUser(ctx, param.FromUserID)
	if err != nil {
		return err
	}

	toUser, err := s.userRepo.GetUser(ctx, param.ToUserID)
	if err != nil {
		return err
	}

	err = s.transferLogsRepo.CreateTransferLogs(ctx, fromUser, toUser, param.Amount)
	if err != nil {
		return err
	}

	if err = s.depositUserBalance(ctx, toUser, param.Amount); err != nil {
		return err
	}

	if err = s.withdrawUserBalance(ctx, fromUser, param.Amount); err != nil {
		return err
	}
	return nil
}

func (s *transferService) depositUserBalance(ctx context.Context, user model.User, amount int32) error {
	if amount <= 0 {
		return errors.New("invalid amount to withdraw")
	}

	user.Balance += amount
	if err := s.userRepo.UpdateUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *transferService) withdrawUserBalance(ctx context.Context, user model.User, amount int32) error {
	if user.Balance < amount {
		return errors.New("invalid balance to withdraw")
	}

	user.Balance -= amount
	if err := s.userRepo.UpdateUser(ctx, user); err != nil {
		return err
	}

	return nil
}
