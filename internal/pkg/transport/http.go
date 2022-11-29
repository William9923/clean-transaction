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

	// // global middlewares
	// e.Use(
	// 	middleware.ContextMiddleware(),
	// 	middleware.Logger(),
	// 	echoMiddleware.BodyDumpWithConfig(echoMiddleware.BodyDumpConfig{
	// 		Skipper: middleware.BodyDumpSkipper,
	// 		Handler: middleware.BodyDumpHandler,
	// 	}),
	// 	middleware.Recover(),
	// 	middleware.CORS(),
	// 	middleware.Headers(),
	// )
	//
	// e.HTTPErrorHandler = ErrHandler{E: e}.Handler
	//
	// e.Debug = conf.Server().GetLoglevel() == "DEBUG"
	srv := server{
		e:    e,
		port: ":8888",
		// port:         conf.Server().GetPort(),
		// readTimeout:  time.Duration(conf.Server().GetReadTimeout()) * time.Millisecond,
		// writeTimeout: time.Duration(conf.Server().GetWriteTimeout()) * time.Millisecond,
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
