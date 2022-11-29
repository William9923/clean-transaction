package repo

import (
	"context"
	"database/sql"

	"github.com/William9923/clean-transaction/internal/data/dao"
	"github.com/William9923/clean-transaction/internal/data/model"
)

type userRepo struct {
	db *sql.DB
}

func UserRepo() dao.UserDAO {
	return userRepo{} // TODO: inject with MySQL dependencies
}

// DepositUserBalance implements dao.UserDAO
func (userRepo) DepositUserBalance(ctx context.Context, user model.User, amount int32) error {
	panic("unimplemented")
}

// GetMultipleUsers implements dao.UserDAO
func (userRepo) GetMultipleUsers(ctx context.Context, userIDs []uint64) ([]model.User, error) {
	panic("unimplemented")
}

// GetUser implements dao.UserDAO
func (userRepo) GetUser(ctx context.Context, userID uint64) (model.User, error) {
	panic("unimplemented")
}
