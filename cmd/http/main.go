package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/William9923/clean-transaction/api"
	"github.com/William9923/clean-transaction/internal/pkg/transport"
)

func main() {
	transports := []transport.Transport{}
	srv := transport.NewServer()
	api.HealthCheck(srv.Engine())
	transports = append(transports, srv)
	stopFn := transport.TransportController(transports...)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)

	sig := <-quit
	log.Printf(fmt.Sprintf("exiting. received signal: %s", sig.String()))

	stopFn(time.Duration(30) * time.Second)
}
