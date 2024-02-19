package http

import (
	"log/slog"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// Chain creates a middleware chain
func Chain(h http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

// Recoverer middleware recovers from panics and logs the error
func Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				slog.ErrorContext(r.Context(), "panic recovered in http handler", "error", err)
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(`{"error": "Internal Server Error"}`))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func NewServer(
	server *http.Server,
	apiRouter http.Handler,
	middlewares ...func(http.Handler) http.Handler,
) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("."))
	})

	tracedRouter := otelhttp.NewHandler(apiRouter, "book-api")
	Mount(mux, "/api", tracedRouter)

	allMiddlewares := []func(http.Handler) http.Handler{
		Recoverer,
	}
	allMiddlewares = append(allMiddlewares, middlewares...)

	server.Handler = Chain(mux, allMiddlewares...)

	return server
}
