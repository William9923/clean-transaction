package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/William9923/clean-transaction/internal/data/dao"
	"github.com/William9923/clean-transaction/internal/data/model"
	internal_mysql "github.com/William9923/clean-transaction/internal/pkg/mysql"
	internal_time "github.com/William9923/clean-transaction/internal/pkg/time"
	"github.com/go-sql-driver/mysql"
)

const (
	queryUpdateUser = `
    UPDATE user_tab
      SET name = ?, balance = ?, updated_at = ?
    WHERE user_id = ?
  `

	queryGetUser = `
    SELECT 
      user_id,
      name,
      balance,
      created_at,
      updated_at
    FROM user_tab
    WHERE user_id = ?
  `

	queryGetUsersInTrx = `
    SELECT 
      user_id,
      name,
      balance,
      created_at,
      updated_at
    FROM user_tab
    WHERE user_id IN (?,?) 
  `
)

type userRepo struct {
	db *sql.DB
}

func UserRepo() dao.UserDAO {
	return userRepo{
		db: internal_mysql.DB(),
	}
}

func (repo userRepo) GetUser(ctx context.Context, userID uint64) (model.User, error) {

	tx := internal_mysql.ExtractTx(ctx)
	val := []interface{}{
		userID,
	}

	var query string = queryGetUser
	var row *sql.Row
	if tx == nil {
		row = repo.db.QueryRowContext(ctx, query, val...)
	} else {
		query += " FOR UPDATE"
		row = tx.QueryRowContext(ctx, query, val...)
	}

	var result model.User
	errDB := row.Scan(
		&result.UserID,
		&result.Name,
		&result.Balance,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

	if errDB != nil {
		if errors.Is(errDB, sql.ErrNoRows) {
			return model.User{}, errors.New("invalid user, no entry for that user ID")
		}
		if errMySQL, ok := errDB.(*mysql.MySQLError); ok {
			return model.User{}, internal_mysql.GetMysqlSpecificError(int(errMySQL.Number), errDB)
		}
		return model.User{}, errDB
	}
	return result, nil
}

func (repo userRepo) GetUsersInTransfer(ctx context.Context, userIDs [2]uint64) ([2]model.User, error) {
	tx := internal_mysql.ExtractTx(ctx)
	val := []interface{}{
		userIDs[0],
		userIDs[1],
	}

	var query string = queryGetUsersInTrx
	var rows *sql.Rows
	var err error
	if tx == nil {
		rows, err = repo.db.QueryContext(ctx, query, val...)
	} else {
		query += " FOR UPDATE"
		rows, err = tx.QueryContext(ctx, query, val...)
	}
	if err != nil {
		return [2]model.User{}, err
	}

	var results [2]model.User
	idx := 0
	for rows.Next() {
		var result model.User

		err := rows.Scan(
			&result.UserID,
			&result.Name,
			&result.Balance,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
		if err != nil {
			return [2]model.User{}, err
		}
		results[idx] = result
		idx += 1
		if idx > len(results) {
			break
		}
	}

	return results, nil
}

func (repo userRepo) UpdateUser(ctx context.Context, user model.User) error {

	tx := internal_mysql.ExtractTx(ctx)
	queryValues := []interface{}{
		user.Name,
		user.Balance,
		internal_time.Now().UnixMilli(),
		user.UserID,
	}

	var errDB error
	if tx == nil {
		_, errDB = repo.db.ExecContext(ctx, queryUpdateUser, queryValues...)
	} else {
		_, errDB = tx.ExecContext(ctx, queryUpdateUser, queryValues...)
	}

	if errDB != nil {
		if errMySQL, ok := errDB.(*mysql.MySQLError); ok {
			return internal_mysql.GetMysqlSpecificError(int(errMySQL.Number), errDB)
		}
		return errDB
	}

	return nil
}
