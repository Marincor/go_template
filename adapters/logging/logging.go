package logging

import (
	"context"
	"time"

	"api.default.marincor.pt/clients/google/logging"
	"api.default.marincor.pt/config/constants"
	"api.default.marincor.pt/entity"
)

func Log(details *entity.LogDetails, severity string, resourceLabels *map[string]string) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(constants.DefaultContextTimeout))
	defer cancel()

	go logging.Log(ctx, details, severity, resourceLabels)
}
