package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/jwzk/go-http-api-boilerplate/internal/domain/model"
	"github.com/jwzk/go-http-api-boilerplate/internal/domain/port"
)

func JsonResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func JsonResponseError(l port.Logger, w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	var statusCodeErr *model.StatusError
	if errors.As(err, &statusCodeErr) {
		w.WriteHeader(statusCodeErr.Code)
		return
	}

	l.Errorf("internal response error: %w", err)
	w.WriteHeader(http.StatusInternalServerError)
}
