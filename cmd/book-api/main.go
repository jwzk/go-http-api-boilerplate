package main

import (
	"context"
	"errors"
	"flag"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	bookmemorydao "github.com/jwzk/go-http-api-boilerplate/internal/adapter/dao/book/memory"
	"github.com/jwzk/go-http-api-boilerplate/internal/adapter/http/bookapi"
	"github.com/jwzk/go-http-api-boilerplate/internal/adapter/http/bookapi/book"
	bookusecase "github.com/jwzk/go-http-api-boilerplate/internal/domain/usecase/book"
	xhttp "github.com/jwzk/go-http-api-boilerplate/pkg/http"
	"github.com/jwzk/go-http-api-boilerplate/pkg/telemetry"
)

const (
	readTimeout  = time.Millisecond * 250
	writeTimeout = time.Millisecond * 30
	idleTimeout  = time.Second * 10
)

func main() {
	// Flag
	loggerLevel := flag.String("level", "info", "logger level (debug|info)")
	httpPort := flag.String("port", "4100", "http api port")
	flag.Parse()

	// Initialize Telemetry and Logger
	telemetry.Init(*loggerLevel)

	// DAO
	bookDAO := bookmemorydao.NewBookDAO()

	// Router
	apiRouter := bookapi.NewRouter(
		book.NewBookRouter(
			bookusecase.NewGetBook(bookDAO),
			bookusecase.NewGetBooks(bookDAO),
			bookusecase.NewCreateBook(bookDAO),
			bookusecase.NewUpdateBook(bookDAO),
			bookusecase.NewDeleteBook(bookDAO),
		))

	// HTTP Server
	s := xhttp.NewServer(
		&http.Server{
			Addr:         net.JoinHostPort("", *httpPort),
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout + time.Millisecond*50,
			IdleTimeout:  idleTimeout,
		},
		apiRouter,
		xhttp.CORS,
		telemetry.AccessLogger,
		func(next http.Handler) http.Handler {
			return http.TimeoutHandler(next, writeTimeout, "timeout")
		},
	)

	// Run Server
	go func() {
		slog.Info("http server listening", "port", *httpPort)

		err := s.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			slog.Error("http server listen failed", "error", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Wait shutdown signal
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := s.Shutdown(ctx)
	if err != nil {
		slog.Error("http server shutdown failed", "error", err)
		os.Exit(1)
	}

	slog.Info("http server shutdown complete")
}
