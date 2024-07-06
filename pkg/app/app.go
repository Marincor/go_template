package app

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"

	"api.default.marincor.pt/app/appinstance"
	"api.default.marincor.pt/clients/mongodb"
	"api.default.marincor.pt/config"
	"api.default.marincor.pt/config/constants"
	"api.default.marincor.pt/pkg/route"
)

func ApplicationInit() {
	configs := config.New()

	db, ctx := mongodb.Connect(configs.DbHost, configs.DbName)

	appinstance.Data = &appinstance.Application{
		DB:        db,
		DBContext: ctx,
		Config:    configs,
		Server: &http.Server{
			Addr: fmt.Sprintf(":%s", constants.Port),

			TLSConfig: &tls.Config{
				ServerName: constants.ServerName,
			},
			ErrorLog: log.New(os.Stderr, fmt.Sprintf("[%s] ", constants.ServerName), log.LstdFlags),
		},
	}

	appinstance.Data.Server.Handler = route.Routes()

}

func Setup() {
	if !constants.UseTLS {
		log.Fatal(appinstance.Data.Server.ListenAndServe())

		return
	}

	log.Fatal(appinstance.Data.Server.ListenAndServeTLS("cert.pem", "key.pem"))

}
