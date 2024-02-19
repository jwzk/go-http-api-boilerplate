package bookapi

import (
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/jwzk/go-http-api-boilerplate/internal/domain/port"
)

func NewHTTPServer(
	l port.Logger,
	port string,
	loggerMiddleware func(http.Handler) http.Handler,
	bookRouter chi.Router,
) *http.Server {
	// Init router
	r := chi.NewRouter()
	apiRouter := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)
	r.Use(loggerMiddleware)
	r.Use(middleware.Recoverer)
	r.Use(func(next http.Handler) http.Handler {
		return http.TimeoutHandler(next, time.Millisecond*250, "")
	})
	r.Use(middleware.Heartbeat("/ping"))

	// Router
	apiRouter.Mount("/books", bookRouter)
	r.Mount("/api", apiRouter)

	// Config server
	return &http.Server{
		Addr:         net.JoinHostPort("", port),
		ReadTimeout:  time.Millisecond * 250,
		WriteTimeout: time.Millisecond * 500,
		IdleTimeout:  time.Second * 10,
		Handler:      r,
	}
}
