package middleware

import (
	"log/slog"
	"net/http"
	"time"

	"api.default.marincor.pt/entity"
)

func Logger() entity.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()

			next.ServeHTTP(w, r)

			elapsedTime := time.Since(startTime)

			slog.Info("http request", slog.String("method", r.Method), slog.String("path", r.URL.Path), slog.String("duration", elapsedTime.String()))
		})
	}
}
