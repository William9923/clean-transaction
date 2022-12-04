package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sync"

	"github.com/William9923/clean-transaction/internal/conf"
	internal_mysql "github.com/William9923/clean-transaction/internal/pkg/mysql"
	"github.com/William9923/clean-transaction/internal/pkg/mysql/repo"
	"github.com/William9923/clean-transaction/internal/services/user"
	"github.com/labstack/gommon/log"
)

func main() {

	cfgpath := flag.String("configpath", "./conf/config.toml", "path to config file")
	if err := conf.Load(*cfgpath); err != nil {
		panic(fmt.Errorf("error parsing config. %w", err))
	}

	if err := internal_mysql.Init(); err != nil {
		panic(fmt.Errorf("error init mysql. %w", err))
	}

	service := user.InitUserService(user.UserServiceParam{
		TransactionManager: repo.TransactionManager(),
		UserRepo:           repo.UserRepo(),
	})

	ctx := context.Background()
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			batchJobs(ctx, service, 1, 1000)
		}()
	}

	wg.Wait()

}

func batchJobs(ctx context.Context, service user.UserService, rangeStart uint32, rangeEnd uint32) {
	for i := rangeStart; i < rangeEnd; i++ {
		userID := uint64(rand.Int31n(10) + 1)
		if err := service.ModifyUserBalance(ctx, user.ModifyUserParam{
			UserID:     userID,
			NewBalance: rand.Int31n(1000) * 1000,
		}); err != nil {
			log.Errorf("failed..., for User ID: %d, with err: %+v", userID, err)
		}
	}
}

func fileExist(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func createFile(name string) error {
	fo, err := os.Create(name)
	if err != nil {
		return err
	}
	defer func() {
		fo.Close()
	}()
	return nil
}
