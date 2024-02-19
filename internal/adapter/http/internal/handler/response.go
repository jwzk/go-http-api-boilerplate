package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/jwzk/go-http-api-boilerplate/internal/domain/model"
	"github.com/jwzk/go-http-api-boilerplate/internal/domain/port"
)

func JsonResponse(l port.Logger, w http.ResponseWriter, data any, err error) {
	statusCode := getErrStatusCode(l, data, err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data == nil {
		return
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(fmt.Errorf("write json response: %w", err))
	}
}

func getErrStatusCode(l port.Logger, data any, err error) int {
	switch {
	case errors.Is(err, model.ErrBadRequest):
		return http.StatusBadRequest
	case errors.Is(err, model.ErrNotFound):
		return http.StatusNotFound
	case err != nil:
		l.Errorf("internal server error response: %w", err)
		return http.StatusInternalServerError
	case data == nil:
		return http.StatusNoContent
	default:
		return http.StatusOK
	}
}
