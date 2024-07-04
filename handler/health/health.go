package health

import (
	"net/http"

	"api.default.marincor.pt/app/usecases/health"
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

	w.Write([]byte("hello"))
}

// func (handler *Handler) Check(ctx *fiber.Ctx) error {
// 	check, err := handler.usecase.Check()
// 	if err != nil {
// 		ctx.Locals(constants.LogDataKey, &entity.LogDetails{
// 			Message:    "error to check health",
// 			Reason:     err.Error(),
// 			StatusCode: constants.HTTPStatusInternalServerError,
// 		})
// 		ctx.Locals(constants.LogSeverityKey, constants.SeverityError)

// 		helpers.CreateResponse(ctx, &entity.ErrorResponse{
// 			Message:     "error to check health",
// 			Description: err.Error(),
// 			StatusCode:  constants.HTTPStatusInternalServerError,
// 		}, constants.HTTPStatusInternalServerError)

// 		return ctx.Next()
// 	}

// 	ctx.Locals(constants.LogDataKey, &entity.LogDetails{
// 		Message:    "successfully health checked",
// 		StatusCode: constants.HTTPStatusOK,
// 		Response:   check,
// 	})
// 	ctx.Locals(constants.LogSeverityKey, constants.SeverityInfo)

// 	helpers.CreateResponse(ctx, check, constants.HTTPStatusOK)

// 	return ctx.Next()
// }
