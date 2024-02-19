package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jwzk/go-http-api-boilerplate/internal/domain/model"
	"github.com/jwzk/go-http-api-boilerplate/pkg/logger"
)

func TestHandler_JsonResponse(t *testing.T) {
	t.Run("empty data", func(t *testing.T) {
		w := httptest.NewRecorder()

		JsonResponse(w, http.StatusBadRequest, nil)

		assert.Equal(t, "application/json", w.Result().Header["Content-Type"][0])
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("with data", func(t *testing.T) {
		w := httptest.NewRecorder()

		JsonResponse(w, http.StatusOK, "test")

		assert.Equal(t, "application/json", w.Result().Header["Content-Type"][0])
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)

		var res string
		json.NewDecoder(w.Result().Body).Decode(&res)

		assert.Equal(t, "test", res)
	})
}

func TestHandler_JsonResponseError(t *testing.T) {
	logger := logger.New("test")

	t.Run("statusCodeErr", func(t *testing.T) {
		w := httptest.NewRecorder()

		JsonResponseError(logger, w, model.ErrBadRequest)

		assert.Equal(t, "application/json", w.Result().Header["Content-Type"][0])
		assert.Equal(t, model.ErrBadRequest.Code, w.Result().StatusCode)
	})

	t.Run("global error", func(t *testing.T) {
		w := httptest.NewRecorder()

		JsonResponseError(logger, w, errors.New("test error"))

		assert.Equal(t, "application/json", w.Result().Header["Content-Type"][0])
		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}
