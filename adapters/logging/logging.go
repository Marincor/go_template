package logging

import (
	"log"
	"os"

	"api.default.marincor.pt/entity"
)

func Log(details *entity.LogDetails) {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	logger.Println(
		details,
	)
}
