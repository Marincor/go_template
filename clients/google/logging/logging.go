package logging

import (
	"context"

	"api.default.marincor.pt/config/constants"
	"api.default.marincor.pt/entity"
)

func Log(_ context.Context, message *entity.LogDetails, severity string, resourceLabels *map[string]string) {
	labels := map[string]string{"service": constants.MainServiceName}
	if resourceLabels != nil {
		for k, v := range *resourceLabels {
			labels[k] = v
		}
	}

	// TODO: IMPLEMENT LOGGING 
	// gcpLogging.Log(
	// 	constants.GcpProjectID,
	// 	constants.MainLoggerName,
	// 	message,
	// 	severity,
	// 	"api",
	// 	labels,
	// )
}
