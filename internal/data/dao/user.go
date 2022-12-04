package dao

import (
	"context"
	"database/sql"

	"github.com/William9923/clean-transaction/internal/data/model"
)

type UserDAO interface {
	GetUsersInTransfer(ctx context.Context, userIDs [2]uint64) ([2]model.User, error)
	GetUser(ctx context.Context, userID uint64) (model.User, error)

	UpdateUser(ctx context.Context, user model.User) error
	InsertUser(ctx context.Context, user model.User) error
}

type UserTxDAO interface {
	GetUser(ctx context.Context, userID uint64, tx *sql.Tx) (model.User, error)
	UpdateUser(ctx context.Context, user model.User, tx *sql.Tx) error
	InsertUser(ctx context.Context, user model.User, tx *sql.Tx) error
}
