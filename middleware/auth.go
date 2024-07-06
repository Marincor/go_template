package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Authorize() func(context *fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		// // TODO: Implement mechanism to authorize requests from given IP to measure endpoint status and metrics
		// // TODO: Log intent to Authorize request

		// authBearer := context.GetReqHeaders()["Authorization"]
		// authSpec := strings.Split(authBearer, " ")

		// if authSpec[0] != "Bearer" {
		// 	helpers.CreateResponse(context, &entity.ErrorResponse{
		// 		Message:    "Unauthorized",
		// 		StatusCode: http.StatusUnauthorized,
		// 	}, http.StatusUnauthorized)
		// }

		// if authSpec[1] == "" {
		// 	helpers.CreateResponse(context, &entity.ErrorResponse{
		// 		Message:    "Unauthorized",
		// 		StatusCode: http.StatusUnauthorized,
		// 	}, http.StatusUnauthorized)
		// }

		// TODO: CHECK
		// if !jwt.New().Validate(authSpec[1]) {
		// 	helpers.CreateResponse(context, &entity.ErrorResponse{
		// 		Message:    "Unauthorized",
		// 		StatusCode: http.StatusUnauthorized,
		// 	}, http.StatusUnauthorized)
		// }

		return context.Next()
	}
}
