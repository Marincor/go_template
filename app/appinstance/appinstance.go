package appinstance

import (
	"database/sql"

	"api.default.marincor.pt/config"
	"github.com/gofiber/fiber/v2"
)

type Application struct {
	Config *config.Config
	DB     *sql.DB
	Server *fiber.App
}

var Data *Application
