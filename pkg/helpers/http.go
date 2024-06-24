package helpers

import (
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

type Request struct {
	Header        interface{}       `json:"header"`
	Proto         string            `json:"proto"`
	ContentLength int64             `json:"content_length"`
	Host          string            `json:"host"`
	RemoteAddr    string            `json:"remote_addr"`
	Method        string            `json:"method"`
	URL           string            `json:"route"`
	Body          interface{}       `json:"body"`
	QueryParams   url.Values        `json:"query_params"`
	Params        map[string]string `json:"params"`
}

func FromHTTPRequest(_ *fiber.Ctx) *Request {
	return &Request{}
}

func CreateResponse(context *fiber.Ctx, payload interface{}, status ...int) {
	returnStatus := http.StatusOK
	if len(status) > 0 {
		returnStatus = status[0]
	}

	context.Status(returnStatus).JSON(payload) //nolint: errcheck
}
