package telemetry

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/trace"
)

func TestInit(t *testing.T) {
	// Should not panic for different log levels
	Init("debug")
	Init("test")
	Init("info")
	Init("unknown")
}

func TestTraceHandler_Handle(t *testing.T) {
	var buf bytes.Buffer
	jsonHandler := slog.NewJSONHandler(&buf, &slog.HandlerOptions{Level: slog.LevelDebug})
	wrappedHandler := traceHandler{Handler: jsonHandler}
	logger := slog.New(wrappedHandler)

	ctx := context.Background()

	// 1. Context without trace
	logger.InfoContext(ctx, "hello world")

	var logMap map[string]any
	err := json.Unmarshal(buf.Bytes(), &logMap)
	assert.NoError(t, err)
	assert.NotContains(t, logMap, "trace_id")

	buf.Reset()

	// 2. Context with valid trace
	// Note: Creating a truly valid fake span context requires setting random IDs
	traceID, _ := trace.TraceIDFromHex("0123456789abcdef0123456789abcdef")
	spanID, _ := trace.SpanIDFromHex("0123456789abcdef")

	spanCtx := trace.NewSpanContext(trace.SpanContextConfig{
		TraceID:    traceID,
		SpanID:     spanID,
		TraceFlags: trace.FlagsSampled,
	})

	ctxWithTrace := trace.ContextWithSpanContext(ctx, spanCtx)

	logger.InfoContext(ctxWithTrace, "hello world with trace")

	err = json.Unmarshal(buf.Bytes(), &logMap)
	assert.NoError(t, err)
	assert.Equal(t, "0123456789abcdef0123456789abcdef", logMap["trace_id"])
	assert.Equal(t, "0123456789abcdef", logMap["span_id"])
}

func TestAccessLogger(t *testing.T) {
	// We want to test that the middleware wraps the handler and allows it to execute
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		_, _ = w.Write([]byte("test body"))
	})

	middleware := AccessLogger(nextHandler)

	req := httptest.NewRequest(http.MethodGet, "/test-path", nil)
	w := httptest.NewRecorder()

	// Execute the middleware
	middleware.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusAccepted, w.Result().StatusCode)
	assert.Equal(t, "test body", w.Body.String())
}
