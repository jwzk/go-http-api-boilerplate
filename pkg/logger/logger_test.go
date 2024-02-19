package logger

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger_New(t *testing.T) {
	t.Run("case debug", func(t *testing.T) {
		l := New("debug")
		assert.NotNil(t, l)
	})

	t.Run("case test", func(t *testing.T) {
		l := New("test")
		assert.NotNil(t, l)
	})

	t.Run("other cases", func(t *testing.T) {
		l := New("abcd")
		assert.NotNil(t, l)
	})
}

func TestLogger_GetMiddleware(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		l := New("test")
		middleware := l.GetMiddleware()
		assert.NotNil(t, middleware)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()

		middleware(handler).ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})
}
