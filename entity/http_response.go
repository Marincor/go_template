package entity

type SuccessfulResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type SuccessListResponse struct {
	Data  interface{} `json:"data"`
	Count int         `json:"count"`
}

type ErrorResponse struct {
	Message     string `json:"error"`
	Description string `json:"description,omitempty"`
	StatusCode  int    `json:"status_code,omitempty"`
}
