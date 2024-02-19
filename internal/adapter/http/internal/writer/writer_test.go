package writer

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jwzk/go-http-api-boilerplate/internal/domain/model"
)

func TestWriter_JSON(t *testing.T) {
	t.Run("empty data", func(t *testing.T) {
		w := httptest.NewRecorder()

		JSON(context.Background(), w, nil, nil)

		assert.Equal(t, "application/json", w.Result().Header["Content-Type"][0])
		assert.Equal(t, http.StatusNoContent, w.Result().StatusCode)
	})

	t.Run("with data", func(t *testing.T) {
		w := httptest.NewRecorder()

		JSON(context.Background(), w, "test", nil)

		assert.Equal(t, "application/json", w.Result().Header["Content-Type"][0])
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)

		var res string
		err := json.NewDecoder(w.Result().Body).Decode(&res)
		assert.NoError(t, err)

		assert.Equal(t, "test", res)
	})

	t.Run("error_with_data", func(t *testing.T) {
		w := httptest.NewRecorder()
		data := make(chan int) // Unmarshallable data to trigger json error
		err := errors.New("test error")
		JSON(context.Background(), w, data, err)
	})
}

func TestWriter_getErrStatusCode(t *testing.T) {
	tests := []struct {
		name        string
		data        any
		err         error
		expectedRes int
	}{
		{
			"statusBadRequest",
			nil,
			model.ErrBadRequest,
			http.StatusBadRequest,
		},
		{
			"statusNotFound",
			nil,
			model.ErrNotFound,
			http.StatusNotFound,
		},
		{
			"statusInternalServerError",
			nil,
			errors.New("test error"),
			http.StatusInternalServerError,
		},
		{
			"statusNoContent",
			nil,
			nil,
			http.StatusNoContent,
		},
		{
			"statusOK",
			"test",
			nil,
			http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := getErrStatusCode(context.Background(), tt.data, tt.err)
			assert.Equal(t, tt.expectedRes, res)
		})
	}
}
