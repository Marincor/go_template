package entity

type LogDetails struct {
	Message    string      `json:"message"`
	Reason     string      `json:"reason"`
	StatusCode int         `json:"status_code"`
	Request    interface{} `json:"request"`
	Response   interface{} `json:"response"`
	User       string      `json:"token"`
	RemoteIP   string      `json:"ipaddress"`
	Method     string      `json:"method"`
	URLpath    string      `json:"route"`
	SessionID  string      `json:"sessid"`
}
