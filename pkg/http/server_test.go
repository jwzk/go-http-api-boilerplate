package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestBookAPI_NewServer(t *testing.T) {
	var (
		expectMiddleware bool
		middleware       = func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expectMiddleware = true
				next.ServeHTTP(w, r)
			})
		}
	)

	t.Run("ok config", func(t *testing.T) {
		s := NewServer(
			&http.Server{
				Addr:         ":4142",
				ReadTimeout:  time.Millisecond * 100,
				WriteTimeout: time.Millisecond * 200,
				IdleTimeout:  time.Millisecond * 300,
			},
			chi.NewRouter(),
		)

		assert.Equal(t, ":4142", s.Addr)
		assert.Equal(t, time.Millisecond*100, s.ReadTimeout)
		assert.Equal(t, time.Millisecond*200, s.WriteTimeout)
		assert.Equal(t, time.Millisecond*300, s.IdleTimeout)
	})

	t.Run("ok middleware", func(t *testing.T) {
		s := NewServer(&http.Server{}, chi.NewRouter(), middleware)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()

		s.Handler.ServeHTTP(w, req)

		assert.True(t, expectMiddleware)
	})

	t.Run("ok not found", func(t *testing.T) {
		s := NewServer(&http.Server{}, chi.NewRouter())

		req := httptest.NewRequest(http.MethodGet, "/not-found", nil)
		w := httptest.NewRecorder()

		s.Handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
	})

	t.Run("ok /ping", func(t *testing.T) {
		s := NewServer(&http.Server{}, chi.NewRouter())

		req := httptest.NewRequest(http.MethodGet, "/ping", nil)
		w := httptest.NewRecorder()

		s.Handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})
}
