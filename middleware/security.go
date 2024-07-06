package middleware

import (
	"net/http"

	"api.default.marincor.pt/entity"
)

func SecurityHeaders() entity.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("X-XSS-Protection", "1; mode=block")
			w.Header().Add("X-Content-Type-Options", "nosniff")
			w.Header().Add("X-Frame-Options", "Deny")
			w.Header().Add("Cache-Control", "no-store")
			w.Header().Add("Content-Security-Policy", "frame-ancestors 'none'")
			w.Header().Add("Content-Security-Policy", "default-src 'none'")
			w.Header().Add("Referrer-Policy", "no-referrer")
			w.Header().Add("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")

			next.ServeHTTP(w, r)
		})
	}
}
