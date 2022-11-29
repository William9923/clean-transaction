package dao

import (
	"context"

	"github.com/William9923/clean-transaction/internal/data/model"
)

type UserDAO interface {
	GetMultipleUsers(ctx context.Context, userIDs []uint64) ([]model.User, error)
	GetUser(ctx context.Context, userID uint64) (model.User, error)

	DepositUserBalance(ctx context.Context, user model.User, amount int32) error
	WithdrawUserBalance(ctx context.Context, user model.User, amount int32) error
}
