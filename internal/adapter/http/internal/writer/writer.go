package writer

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/jwzk/go-http-api-boilerplate/internal/domain/model"
	pkghttp "github.com/jwzk/go-http-api-boilerplate/pkg/http"
)

// JSON maps domain errors to HTTP status codes and writes a JSON response.
func JSON(ctx context.Context, w http.ResponseWriter, data any, err error) {
	statusCode := getErrStatusCode(ctx, data, err)
	pkghttp.JSON(ctx, w, statusCode, data)
}

func getErrStatusCode(ctx context.Context, data any, err error) int {
	switch {
	case errors.Is(err, model.ErrBadRequest):
		return http.StatusBadRequest
	case errors.Is(err, model.ErrNotFound):
		return http.StatusNotFound
	case err != nil:
		slog.ErrorContext(ctx, "internal server error response", "error", err)
		return http.StatusInternalServerError
	case data == nil:
		return http.StatusNoContent
	default:
		return http.StatusOK
	}
}
