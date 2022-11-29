package mysql

import (
	"database/sql"
	"fmt"

	"github.com/William9923/clean-transaction/internal/conf"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func Init() error {
	cfg := conf.GetConfig()
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		cfg.MySQL.User,
		cfg.MySQL.Pass,
		cfg.MySQL.Host,
		cfg.MySQL.Port,
		cfg.MySQL.DBName,
	)

	var err error

	db, err = sql.Open("mysql", connStr)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	if cfg.MySQL.MaxIdle > 0 {
		db.SetMaxIdleConns(cfg.MySQL.MaxIdle)
	}

	if cfg.MySQL.MaxOpen > 0 {
		db.SetMaxOpenConns(cfg.MySQL.MaxOpen)
	}

	return nil
}

func DB() *sql.DB {
	return db
}
