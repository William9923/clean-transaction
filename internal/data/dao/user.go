package dao

import (
	"context"

	"github.com/William9923/clean-transaction/internal/data/model"
)

type UserDAO interface {
	GetUsersInTransfer(ctx context.Context, userIDs [2]uint64) ([2]model.User, error)
	GetUser(ctx context.Context, userID uint64) (model.User, error)

	UpdateUser(ctx context.Context, user model.User) error
}
