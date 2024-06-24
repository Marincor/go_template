package constants

import (
	"fmt"
	"net/http"
)

const (
	HTTPStatusOK                  = http.StatusOK
	HTTPStatusCreated             = http.StatusCreated
	HTTPStatusAccepted            = http.StatusAccepted
	HTTPStatusNoContent           = http.StatusNoContent
	HTTPStatusBadRequest          = http.StatusBadRequest
	HTTPStatusUnauthorized        = http.StatusUnauthorized
	HTTPStatusNotAcceptable       = http.StatusNotAcceptable
	HTTPStatusTooManyRequests     = http.StatusTooManyRequests
	HTTPStatusInternalServerError = http.StatusInternalServerError
)

var HTTPStatusesOk = []string{
	fmt.Sprintf("%d", HTTPStatusOK),
	fmt.Sprintf("%d", HTTPStatusCreated),
	fmt.Sprintf("%d", HTTPStatusAccepted),
	fmt.Sprintf("%d", HTTPStatusNoContent),
}
