package logging

import (
	"api.default.marincor/config/constants"
	"api.default.marincor/entity"
	gcpLogging "github.com/INDICO-INNOVATION/gcp_logging_easycall"
)

func Log(message *entity.LogDetails, severity string, resourceLabels *map[string]string) {
	logMessage := &gcpLogging.Logger{
		User:         message.User,
		Message:      message.Message,
		Reason:       message.Reason,
		RemoteIp:     message.RemoteIP,
		Method:       message.Method,
		Urlpath:      message.URLpath,
		StatusCode:   message.StatusCode,
		RequestData:  message.RequestData,
		ResponseData: message.ResponseData,
		SessionId:    message.SessionID,
	}

	labels := map[string]string{"service": constants.MainServiceName}
	if resourceLabels != nil {
		for k, v := range *resourceLabels {
			labels[k] = v
		}
	}

	gcpLogging.Log(
		constants.GcpProjectID,
		constants.MainLoggerName,
		logMessage,
		severity,
		"api",
		labels,
	)
}
