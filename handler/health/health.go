package health

import (
	"net/http"

	"api.default.marincor.pt/adapters/logging"
	"api.default.marincor.pt/app/usecases/health"
	"api.default.marincor.pt/config/constants"
	"api.default.marincor.pt/entity"
	"api.default.marincor.pt/pkg/helpers"
)

type Handler struct {
	usecase *health.Usecase
}

func Handle() *Handler {
	return &Handler{
		usecase: health.New(),
	}
}

func (handler *Handler) Check(w http.ResponseWriter, r *http.Request) {
	check, err := handler.usecase.Check()
	if err != nil {
		go logging.Log(&entity.LogDetails{
			Message:    "error to check health",
			Reason:     err.Error(),
			StatusCode: constants.HTTPStatusInternalServerError,
			Severity:   string(constants.SeverityError),
		})

		helpers.CreateResponse(w, &entity.ErrorResponse{
			Message:     "error to check health",
			Description: err.Error(),
			StatusCode:  constants.HTTPStatusInternalServerError,
		})

		return
	}

	go logging.Log(&entity.LogDetails{
		Message:    "successfully health checked",
		StatusCode: constants.HTTPStatusOK,
		Severity:   string(constants.SeverityInfo),
	})

	helpers.CreateResponse(w, check)
}

func (handler *Handler) List(w http.ResponseWriter, r *http.Request) {
	checkList, err := handler.usecase.List()
	if err != nil {
		go logging.Log(&entity.LogDetails{
			Message:    "error to check health",
			Reason:     err.Error(),
			StatusCode: constants.HTTPStatusInternalServerError,
			Severity:   string(constants.SeverityError),
		})

		helpers.CreateResponse(w, &entity.ErrorResponse{
			Message:     "error to check health",
			Description: err.Error(),
			StatusCode:  constants.HTTPStatusInternalServerError,
		})

		return
	}

	go logging.Log(&entity.LogDetails{
		Message:    "successfully health checked",
		StatusCode: constants.HTTPStatusOK,
		Severity:   string(constants.SeverityInfo),
	})

	listResponse := entity.SuccessListResponse{
		Data:  checkList,
		Count: len(checkList),
	}

	helpers.CreateResponse(w, listResponse)
}
