package middleware

import (
	"io"
	"net/http"
	"time"

	"api.default.marincor.pt/adapters/logging"
	"api.default.marincor.pt/config/constants"
	"api.default.marincor.pt/entity"
)

func Logger() entity.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()

			next.ServeHTTP(w, r)

			elapsedTime := time.Since(startTime)

			statusCode := w.Header().Get(constants.StatusCodeContextKey)
			w.Header().Del(constants.StatusCodeContextKey)

			bodyBytes, _ := io.ReadAll(r.Body)

			go logging.Log(&entity.LogDetails{
				Message:  "Logger",
				RemoteIP: r.RemoteAddr,
				Request: map[string]interface{}{
					"body":     string(bodyBytes),
					"query":    r.URL.RawQuery,
					"path":     r.URL.Path,
					"duration": elapsedTime.String(),
				},
				Response: map[string]interface{}{
					"status_code": statusCode,
				},
			})
		})
	}
}
