package logging

import (
	"api.default.marincor/clients/google/logging"
	"api.default.marincor/entity"
)

func Log(details *entity.LogDetails, severity string, resourceLabels *map[string]string) {
	go logging.Log(details, severity, resourceLabels)
}
