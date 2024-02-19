package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewServer(
	server *http.Server,
	apiRouter chi.Router,
	middlewares ...func(http.Handler) http.Handler,
) *http.Server {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middlewares...)
	r.Use(middleware.Heartbeat("/ping"))

	r.Mount("/api", apiRouter)
	server.Handler = r

	return server
}
