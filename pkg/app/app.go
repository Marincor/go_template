package app

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"api.default.marincor/adapters/logging"
	"api.default.marincor/clients/iam"
	"api.default.marincor/config"
	"api.default.marincor/config/constants"
	"api.default.marincor/entity"
	"api.default.marincor/pkg/helpers"
	json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type Application struct {
	Config *config.Config
	Server *fiber.App
}

var Inst *Application

func ApplicationInit() {
	configs := config.New()

	iam.New()

	Inst = &Application{
		Config: configs,
		Server: fiber.New(fiber.Config{
			ServerHeader: "Death Star",
			ErrorHandler: customErrorHandler,
			JSONEncoder:  json.Marshal,
			JSONDecoder:  json.Unmarshal,
			Prefork:      constants.Prefork,
		}),
	}
}

func Setup() {
	var err error
	if constants.UseTLS {
		err = Inst.Server.ListenTLS(fmt.Sprintf(":%s", constants.Port), "cert.pem", "key.pem")
	} else {
		err = Inst.Server.Listen(fmt.Sprintf(":%s", constants.Port))
	}

	if errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}

func customErrorHandler(context *fiber.Ctx, err error) error {
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
			Message:     message,
			Reason:      err.Error(),
			StatusCode:  code,
			RequestData: string(context.Body()),
		},
		"critical",
		nil,
	)

	return helpers.CreateResponse(context, errorResponse, code) //nolint: wrapcheck
}
