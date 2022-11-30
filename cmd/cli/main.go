package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/William9923/clean-transaction/internal/conf"
	internal_mysql "github.com/William9923/clean-transaction/internal/pkg/mysql"
	"github.com/William9923/clean-transaction/internal/pkg/mysql/repo"
	"github.com/William9923/clean-transaction/internal/services/transfer"
	"github.com/labstack/gommon/log"
)

func main() {

	cfgpath := flag.String("configpath", "./conf/config.toml", "path to config file")
	fromTargetUser := flag.Int64("from", 1, "user to transfer from...")
	toTargetUser := flag.Int64("to", 2, "user to transfer to...")
	amount := flag.Int64("amount", 0, "amount to transfer")
	flag.Parse()

	if err := conf.Load(*cfgpath); err != nil {
		panic(fmt.Errorf("error parsing config. %w", err))
	}

	if err := internal_mysql.Init(); err != nil {
		panic(fmt.Errorf("error init mysql. %w", err))
	}

	service := transfer.InitTransferService(transfer.TransferServiceParam{
		TransactionManager: repo.TransactionManager(),
		TransferLogsRepo:   repo.TransferLogRepo(),
		UserRepo:           repo.UserRepo(),
	})

	ctx := context.Background()
	err := service.Transfer(ctx, transfer.DoTransferParam{
		FromUserID: uint64(*fromTargetUser),
		ToUserID:   uint64(*toTargetUser),
		Amount:     int32(*amount),
	})
	if err != nil {
		log.Error(err)
	}
}
