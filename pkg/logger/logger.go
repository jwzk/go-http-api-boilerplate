package logger

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

type Logger struct {
	*zap.SugaredLogger
	logger *zap.Logger
}

func New(level string) *Logger {
	var l *zap.Logger

	if level == "debug" {
		l = zap.Must(zap.NewDevelopment())
	} else if level == "test" {
		l = zap.NewNop()
	} else {
		l = zap.Must(zap.NewProduction())
	}

	return &Logger{SugaredLogger: l.Sugar(), logger: l}
}

func (l Logger) GetMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			res := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t := time.Now()
			defer func() {
				l.logger.Info(fmt.Sprintf("%s %s %s", r.Method, r.URL.Path, r.Proto),
					zap.String("method", r.Method),
					zap.String("path", r.URL.Path),
					zap.Int("status", res.Status()),
					zap.String("ip", r.RemoteAddr),
					zap.String("id", middleware.GetReqID(r.Context())),
					zap.Duration("lat", time.Since(t)),
					zap.Int("size", res.BytesWritten()),
				)
			}()

			next.ServeHTTP(res, r)
		})
	}
}
