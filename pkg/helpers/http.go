package helpers

import (
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

type Request struct {
	Method           string          `json:"method"`
	URL              string          `json:"url"`
	Proto            string          `json:"proto"`
	Header           interface{}     `json:"header"`
	Body             interface{}     `json:"body"`
	QueryParams      interface{}     `json:"query_params"`
	ContentLength    int64           `json:"content_length"`
	TransferEncoding []string        `json:"transfer_encoding"`
	Host             string          `json:"host"`
	Form             url.Values      `json:"form"`
	PostForm         url.Values      `json:"post_form"`
	MultipartForm    *multipart.Form `json:"multipart_form"`
	RemoteAddr       string          `json:"remote_addr"`
}

func FromHTTPRequest(context *fiber.Ctx) *Request {
	multipartform, _ := context.MultipartForm()
	req := &Request{
		Method:        context.Method(),
		URL:           context.OriginalURL(),
		Proto:         context.Protocol(),
		Header:        context.GetReqHeaders(),
		ContentLength: int64(context.Request().Header.ContentLength()),
		Host:          context.Hostname(),
		MultipartForm: multipartform,
		RemoteAddr:    context.IP(),
		Body:          string(context.Body()),
		QueryParams:   context.Context().QueryArgs().QueryString(),
	}

	return req
}

func CreateResponse(context *fiber.Ctx, payload interface{}, status ...int) error {
	returnStatus := http.StatusOK
	if len(status) > 0 {
		returnStatus = status[0]
	}

	return context.Status(returnStatus).JSON(payload)
}
