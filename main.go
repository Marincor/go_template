package main

import "api.default.marincor/pkg/app"

func main() {
	app.ApplicationInit()
	app.Inst.Server = route()

	// Listening to Server
	app.Setup()
}
