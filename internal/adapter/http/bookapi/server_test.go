package bookapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"

	"github.com/jwzk/go-http-api-boilerplate/pkg/logger"
)

func TestBookAPI_NewHTTPServer(t *testing.T) {
	logger := logger.New("test")
	middleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}

	t.Run("ok config", func(t *testing.T) {
		s := NewHTTPServer(logger, "4140", middleware, chi.NewRouter())
		assert.Equal(t, ":4140", s.Addr)
		assert.Equal(t, time.Millisecond*250, s.ReadTimeout)
		assert.Equal(t, time.Millisecond*500, s.WriteTimeout)
		assert.Equal(t, time.Second*10, s.IdleTimeout)
	})

	t.Run("ok /ping", func(t *testing.T) {
		s := NewHTTPServer(logger, "4140", middleware, chi.NewRouter())

		req := httptest.NewRequest(http.MethodGet, "/ping", nil)
		w := httptest.NewRecorder()

		s.Handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})

	t.Run("ok not found", func(t *testing.T) {
		s := NewHTTPServer(logger, "4140", middleware, chi.NewRouter())

		req := httptest.NewRequest(http.MethodGet, "/error", nil)
		w := httptest.NewRecorder()

		s.Handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
	})

	t.Run("ok /books", func(t *testing.T) {
		router := chi.NewRouter().Route("/books", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNoContent)
			})
		})

		s := NewHTTPServer(logger, "4140", middleware, router)

		req := httptest.NewRequest(http.MethodGet, "/api/books/", nil)
		w := httptest.NewRecorder()

		s.Handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNoContent, w.Result().StatusCode)
	})
}
