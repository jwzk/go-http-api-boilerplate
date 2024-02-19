package main

import (
	"context"
	"errors"
	"flag"
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
	"github.com/jwzk/go-http-api-boilerplate/pkg/logger"
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

	// Logger
	l := logger.New(*loggerLevel)
	defer l.Sync()

	// DAO
	bookDAO := bookmemorydao.NewBookDAO()

	// Router
	apiRouter := bookapi.NewRouter(
		book.NewBookRouter(
			l,
			bookusecase.NewGetBook(l, bookDAO),
			bookusecase.NewGetBooks(l, bookDAO),
			bookusecase.NewCreateBook(l, bookDAO),
			bookusecase.NewUpdateBook(l, bookDAO),
			bookusecase.NewDeleteBook(l, bookDAO),
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
		l.GetMiddleware(),
		func(next http.Handler) http.Handler {
			return http.TimeoutHandler(next, writeTimeout, "timeout")
		},
	)

	// Run Server
	go func() {
		l.Infof("http server listening on :%s", *httpPort)

		err := s.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			l.Fatalf("http server listen: %v", err)
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
		l.Fatalf("http server shutdown: %v", err)
	}

	l.Info("http server shutdown complete")
}
