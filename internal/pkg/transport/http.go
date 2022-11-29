package transport

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type server struct {
	e            *echo.Echo
	port         string
	readTimeout  time.Duration
	writeTimeout time.Duration
}

func NewServer() server {
	e := echo.New()

	srv := server{
		e:            e,
		port:         ":8888",
		readTimeout:  500 * time.Millisecond,
		writeTimeout: 500 * time.Millisecond,
	}

	return srv
}

func (srv server) Engine() *echo.Echo {
	return srv.e
}

func (srv server) Start() {
	s := &http.Server{
		Addr:         srv.port,
		ReadTimeout:  srv.readTimeout,
		WriteTimeout: srv.writeTimeout,
	}

	if err := srv.e.StartServer(s); err != nil {
		srv.e.Logger.Error(err)
		srv.e.Logger.Info("Shutting down the server")
		os.Exit(1)
	}
}

func (srv server) Stop() {
	ctx := context.Background()
	if err := srv.e.Shutdown(ctx); err != nil {
		srv.e.Logger.Fatal(err)
	}
}
