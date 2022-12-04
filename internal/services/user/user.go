package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/William9923/clean-transaction/internal/data/dao"
	"github.com/William9923/clean-transaction/internal/data/model"
)

type ModifyUserParam struct {
	UserID     uint64
	NewBalance int32
}

type UserServiceParam struct {
	TransactionManager dao.TransactionManager
	UserRepo           dao.UserDAO
}

type UserService interface {
	ModifyUserBalance(ctx context.Context, param ModifyUserParam) error
	ModifyUserBalanceV2(ctx context.Context, param ModifyUserParam) error
}

type userService struct {
	transactionManager dao.TransactionManager
	userRepo           dao.UserDAO
}

func InitUserService(params UserServiceParam) UserService {
	return &userService{
		transactionManager: params.TransactionManager,
		userRepo:           params.UserRepo,
	}
}

func isUserNotFound(err error) bool {
	return err != nil && errors.Is(err, sql.ErrNoRows)
}

// NOTE: Transaction with the possibility of deadlock
func (s *userService) ModifyUserBalance(ctx context.Context, param ModifyUserParam) error {
	return s.transactionManager.WithinTransaction(ctx, func(trxCtx context.Context) error {

		var found bool = true
		user, err := s.userRepo.GetUser(trxCtx, param.UserID)
		if err != nil {
			if isUserNotFound(err) {
				found = false
			} else {
				return err
			}
		}

		if found {
			user.Balance = param.NewBalance
			if err := s.userRepo.UpdateUser(trxCtx, user); err != nil {
				return err
			}
		} else {
			newUser := model.User{
				Name:    "New User",
				Balance: param.NewBalance,
			}
			if err := s.userRepo.InsertUser(trxCtx, newUser); err != nil {
				return err
			}
		}
		return nil
	})
}

// NOTE: Transaction with deadlock free
func (s *userService) ModifyUserBalanceV2(ctx context.Context, param ModifyUserParam) error {

	var found bool = true
	if err := s.updateExistingUserBalance(ctx, param); err != nil {
		if isUserNotFound(err) {
			found = false
		} else {
			return err
		}
	}

	if !found {
		if err := s.insertNonExistingUserBalance(ctx, param); err != nil {
			return err
		}
	}

	return nil
}

func (s *userService) updateExistingUserBalance(ctx context.Context, param ModifyUserParam) error {
	return s.transactionManager.WithinTransaction(ctx, func(trxCtx context.Context) error {
		user, err := s.userRepo.GetUser(trxCtx, param.UserID)
		if err != nil {
			return err
		}

		user.Balance = param.NewBalance
		if err := s.userRepo.UpdateUser(trxCtx, user); err != nil {
			return err
		}
		return nil
	})
}

func (s *userService) insertNonExistingUserBalance(ctx context.Context, param ModifyUserParam) error {

	return s.transactionManager.WithinTransaction(ctx, func(trxCtx context.Context) error {
		newUser := model.User{
			Name:    "New User",
			Balance: param.NewBalance,
		}
		if err := s.userRepo.InsertUser(trxCtx, newUser); err != nil {
			return err
		}
		return nil
	})
}
