package appinstance

import (
	"database/sql"

	"net/http"

	"api.default.marincor.pt/config"
)

type Application struct {
	Config *config.Config
	DB     *sql.DB
	Server *http.Server
}

var Data *Application
