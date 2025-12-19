package middleware

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

const LoggerKey string = "logger"

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		reqIP := r.Header.Get("X-Forwarded-For")
		if reqIP == "" {
			host, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				reqIP = r.RemoteAddr
			} else {
				reqIP = host
			}
		}

		reqId := middleware.GetReqID(r.Context())
		reqLogger := slog.With(
			slog.String("requestId", reqId),
			slog.String("requestIP", reqIP),
			slog.String("method", r.Method),
			slog.String("url", r.URL.Path),
		)

		ctx := context.WithValue(r.Context(), LoggerKey, reqLogger)
		wrapWriter := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		next.ServeHTTP(wrapWriter, r.WithContext(ctx))

		httpStatus := wrapWriter.Status()
		attrs := []any{
			slog.Int("status", httpStatus),
			slog.Duration("duration", time.Since(start)),
		}

		if httpStatus >= 500 {
			reqLogger.Error("server error", attrs...)
		} else {
			reqLogger.Info("request processed", attrs...)
		}
	})
}

func GetLogger(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(LoggerKey).(*slog.Logger); ok {
		return logger
	}

	return slog.Default()
}
