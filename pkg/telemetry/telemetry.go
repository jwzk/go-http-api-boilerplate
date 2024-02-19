package telemetry

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/felixge/httpsnoop"
	"go.opentelemetry.io/otel/trace"
)

// traceHandler is a custom slog.Handler that automatically extracts
// OpenTelemetry trace_id and span_id from the context and appends
// them to the log record for distributed tracing correlation (e.g. Datadog).
type traceHandler struct {
	slog.Handler
}

// Handle adds tracing attributes to the log record if a valid span exists in the context.
func (h traceHandler) Handle(ctx context.Context, r slog.Record) error {
	spanCtx := trace.SpanContextFromContext(ctx)
	if spanCtx.IsValid() {
		r.AddAttrs(
			slog.String("trace_id", spanCtx.TraceID().String()),
			slog.String("span_id", spanCtx.SpanID().String()),
		)
	}
	return h.Handler.Handle(ctx, r)
}

// Init configures the global slog logger with JSON formatting and tracing correlation.
func Init(level string) {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	if level == "debug" {
		opts.Level = slog.LevelDebug
	} else if level == "test" {
		opts.Level = slog.LevelError
	}

	// Create a standard JSON handler
	jsonHandler := slog.NewJSONHandler(os.Stdout, opts)

	// Wrap it with our trace extracting handler
	wrappedHandler := traceHandler{Handler: jsonHandler}

	// Set it as the global default logger
	slog.SetDefault(slog.New(wrappedHandler))
}

// AccessLogger is an HTTP middleware that logs request access details
func AccessLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := httpsnoop.CaptureMetrics(next, w, r)

		slog.InfoContext(r.Context(), fmt.Sprintf("%s %s %s", r.Method, r.URL.Path, r.Proto),
			"method", r.Method,
			"path", r.URL.Path,
			"status", m.Code,
			"ip", r.RemoteAddr,
			"lat", m.Duration.String(),
			"size", m.Written,
		)
	})
}
