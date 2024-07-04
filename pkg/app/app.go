package app

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"api.default.marincor.pt/adapters/logging"
	"api.default.marincor.pt/app/appinstance"
	"api.default.marincor.pt/config"
	"api.default.marincor.pt/config/constants"
	"api.default.marincor.pt/entity"
	"api.default.marincor.pt/pkg/helpers"
	json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func ApplicationInit() {
	configs := config.New()

	if constants.Environment != constants.Test {
		// iam.New()
	}

	appinstance.Data = &appinstance.Application{
		Config: configs,
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%s", constants.Port),
			Handler: CustomErrorHandler{},
			TLSConfig: &tls.Config{
				ServerName: constants.ServerName,
			},
			ErrorLog: log.New(os.Stderr, fmt.Sprintf("[%s] ", constants.ServerName), log.LstdFlags),
		},
	}

	// TODO: CHECK
	// appinstance.Data.DB = postgres.Connect(appinstance.Data.Config.DBString)
}

func Setup() {
	if !constants.UseTLS {
		// err = appinstance.Data.Server.Listen(fmt.Sprintf(":%s", constants.Port))
		log.Fatal(appinstance.Data.Server.ListenAndServe())
	} else {
		// err = appinstance.Data.Server.ListenAndServeTLS(fmt.Sprintf(":%s", constants.Port), "cert.pem", "key.pem")
		certPem, err := os.ReadFile("cert.pem")
		if err != nil {
			log.Fatalf("Failed to read CA certificate: %v", err)
		}

		keyPem, err := os.ReadFile("key.pem")
		if err != nil {
			log.Fatalf("Failed to read Key certificate: %v", err)
		}

		fmt.Println((certPem))

		log.Fatal(appinstance.Data.Server.ListenAndServeTLS(string(certPem), string(keyPem)))
	}

}

type CustomErrorHandler struct{}

// ServeHTTP implementa a interface http.Handler.
func (ceh CustomErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Simule um erro (por exemplo, página não encontrada).
	http.Error(w, "Página não encontrada", http.StatusNotFound)
}

func customErrorHandlerss(ctx *fiber.Ctx, err error) error {
	var code int = fiber.StatusInternalServerError
	var capturedError *fiber.Error
	message := "unknown error"

	if errors.As(err, &capturedError) {
		code = capturedError.Code
		if code == fiber.StatusNotFound {
			message = "route not found"
		}
	}

	var errorResponse *entity.ErrorResponse

	erro := json.Unmarshal([]byte(err.Error()), &errorResponse)
	if erro != nil {
		errorResponse = &entity.ErrorResponse{
			Message:     message,
			StatusCode:  code,
			Description: err.Error(),
		}
	}

	go logging.Log(
		&entity.LogDetails{
			Message:  message,
			Method:   ctx.Method(),
			Reason:   err.Error(),
			RemoteIP: ctx.IP(),
			Request: map[string]interface{}{
				"body":       string(ctx.BodyRaw()),
				"query":      ctx.Queries(),
				"url_params": ctx.AllParams(),
			},
			StatusCode: code,
			URLpath:    ctx.Path(),
		},
		"critical",
		nil,
	)

	helpers.CreateResponse(ctx, errorResponse, code) //nolint: wrapcheck

	return nil
}

func Log(ctx *fiber.Ctx) error {
	logSeverity := ctx.Locals(constants.LogSeverityKey)

	payload := new(entity.LogDetails)
	bytedata, _ := helpers.Marshal(ctx.Locals(constants.LogDataKey))
	helpers.Unmarshal(bytedata, &payload) //nolint: errcheck

	if logSeverity == nil {
		logSeverity = "debug"
	}

	body := map[string]interface{}{}
	helpers.Unmarshal(ctx.BodyRaw(), &body) //nolint: errcheck

	request := map[string]interface{}{
		"body":       body,
		"query":      ctx.Queries(),
		"url_params": ctx.AllParams(),
	}

	severity := fmt.Sprintf("%v", logSeverity)

	logging.Log(&entity.LogDetails{
		Message:    payload.Message,
		StatusCode: payload.StatusCode,
		Reason:     payload.Reason,
		Response:   payload.Response,
		Request:    request,
		Method:     ctx.Method(),
		RemoteIP:   ctx.IP(),
		URLpath:    ctx.Path(),
	}, severity, nil)

	return nil
}
