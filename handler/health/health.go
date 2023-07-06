package health

import (
	"net/http"

	"api.default.marincor/app/usecases/health"
	"api.default.marincor/entity"
	"api.default.marincor/pkg/helpers"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	uc *health.Usecase
}

func Handle() *Handler {
	return &Handler{
		uc: health.New(),
	}
}

func (handler *Handler) Check(context *fiber.Ctx) error {
	// timeNow := time.Now().UTC()
	if _, err := handler.uc.Check(); err != nil {
		return helpers.CreateResponse(context, &entity.SuccessfulResponse{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		})
	}

	return helpers.CreateResponse(context, &entity.SuccessfulResponse{
		Message:    "OK",
		StatusCode: http.StatusOK,
	})
}
