package middleware

import (
	"net/http"

	"api.default.marincor.pt/config/constants"
	"api.default.marincor.pt/entity"
	"api.default.marincor.pt/pkg/helpers"
)

func Cors() entity.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", constants.AllowedOrigins) // change this later
			w.Header().Set("Access-Control-Allow-Headers", constants.AllowedHeaders)
			w.Header().Set("Access-Control-Allow-Methods", constants.AllowedMethods)

			if r.Method == "OPTIONS" {
				helpers.CreateResponse(w, nil)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
