package entity

import "net/http"

type (
	CustomResponseWriter struct {
		http.ResponseWriter
		StatusCode int
		Response   interface{}
	}
)
