package entity

type LogDetails struct {
	User         string      `json:"token"`
	Message      string      `json:"message"`
	Reason       string      `json:"reason"`
	RemoteIP     string      `json:"ipaddress"`
	Method       string      `json:"method"`
	URLpath      string      `json:"route"`
	StatusCode   int         `json:"status_code"`
	RequestData  interface{} `json:"request_data"`
	ResponseData interface{} `json:"response_data"`
	SessionID    string      `json:"sessid"`
}
