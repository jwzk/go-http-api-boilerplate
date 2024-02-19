package http

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

// JSON sends a JSON response with the given status code.
func JSON(ctx context.Context, w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		slog.WarnContext(ctx, "failed to write json response", "error", err)
	}
}
