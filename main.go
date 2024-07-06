package main

import (
	"api.default.marincor.pt/app/appinstance"
	"api.default.marincor.pt/pkg/app"
)

func main() {
	app.ApplicationInit()
	defer func() {
		if err := appinstance.Data.DB.Client().Disconnect(*appinstance.Data.DBContext); err != nil {
			panic(err)
		}
	}()

	// Listening to Server
	app.Setup()
}
