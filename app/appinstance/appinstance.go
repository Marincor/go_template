package appinstance

import (
	"context"
	"net/http"

	"api.default.marincor.pt/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	Config    *config.Config
	DB        *mongo.Database
	DBContext *context.Context
	Server    *http.Server
}

var Data *Application
