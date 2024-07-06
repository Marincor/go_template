package middleware

import (
	"net/http"

	"api.default.marincor.pt/config/constants"
	"api.default.marincor.pt/entity"
	"api.default.marincor.pt/pkg/helpers"
)

func ContentType() entity.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if len(r.Header["Accept"]) == 0 {
				helpers.CreateResponse(w, &entity.ErrorResponse{
					Message:    "Content Not Accepted",
					StatusCode: http.StatusNotAcceptable,
				}, http.StatusNotAcceptable)

				return
			}

			for _, header := range r.Header["Accept"] {
				if !helpers.Contains(constants.AllowedContentTypes, header) {
					helpers.CreateResponse(w, &entity.ErrorResponse{
						Message:    "Content Not Accepted",
						StatusCode: http.StatusNotAcceptable,
					}, http.StatusNotAcceptable)

					return
				}
			}

			next.ServeHTTP(w, r)

		})
	}
}
