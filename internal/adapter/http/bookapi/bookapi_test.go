package bookapi

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookAPI_NewRouter(t *testing.T) {
	t.Run("ok /books/", func(t *testing.T) {
		router := http.NewServeMux()
		router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
		})

		s := NewRouter(router)

		req := httptest.NewRequest(http.MethodGet, "/books/", nil)
		w := httptest.NewRecorder()

		s.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNoContent, w.Result().StatusCode)
	})

	t.Run("ok /books exact match", func(t *testing.T) {
		router := http.NewServeMux()
		router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		s := NewRouter(router)

		req := httptest.NewRequest(http.MethodGet, "/books", nil)
		w := httptest.NewRecorder()

		s.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})
}
