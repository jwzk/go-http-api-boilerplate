package http

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSON(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		w := httptest.NewRecorder()
		data := map[string]string{"foo": "bar"}

		JSON(context.Background(), w, http.StatusOK, data)

		assert.Equal(t, "application/json", w.Result().Header.Get("Content-Type"))
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.JSONEq(t, `{"foo":"bar"}`, w.Body.String())
	})

	t.Run("nil data", func(t *testing.T) {
		w := httptest.NewRecorder()

		JSON(context.Background(), w, http.StatusNoContent, nil)

		assert.Equal(t, "application/json", w.Result().Header.Get("Content-Type"))
		assert.Equal(t, http.StatusNoContent, w.Result().StatusCode)
		assert.Empty(t, w.Body.String())
	})

	t.Run("unmarshallable data", func(t *testing.T) {
		w := httptest.NewRecorder()
		data := make(chan int) // Cannot be marshalled to JSON

		JSON(context.Background(), w, http.StatusOK, data)

		assert.Equal(t, "application/json", w.Result().Header.Get("Content-Type"))
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		// Body shouldn't contain the channel
		assert.Empty(t, w.Body.String())
	})
}

// errorWriter simulates a broken http.ResponseWriter that fails on Write
type errorWriter struct {
	http.ResponseWriter
}

func (ew *errorWriter) Write(b []byte) (int, error) {
	return 0, errors.New("write error")
}

func TestJSON_WriteError(t *testing.T) {
	w := httptest.NewRecorder()
	ew := &errorWriter{ResponseWriter: w}
	data := map[string]string{"foo": "bar"}

	// This should trigger the slog.WarnContext line
	JSON(context.Background(), ew, http.StatusOK, data)

	assert.Equal(t, "application/json", w.Result().Header.Get("Content-Type"))
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
}
