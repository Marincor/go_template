package route

import (
	"fmt"
	"net/http"

	"api.default.marincor.pt/config/constants"
	"api.default.marincor.pt/entity"
	"api.default.marincor.pt/handler/health"
	"api.default.marincor.pt/middleware"
)

func newRouter(mx ...entity.Middleware) *entity.Router {
	return &entity.Router{
		ServeMux: &http.ServeMux{},
		Chain:    mx,
	}
}

func Routes() *entity.Router {
	allowedOrigins := constants.AllowedOrigins
	if constants.Environment != constants.Production {
		allowedOrigins += fmt.Sprintf(", %s", constants.AllowedStageOrigins)
	}

	r := newRouter(middleware.Logger())

	r.Group(func(r *entity.Router) {
		// r.Use(middleware.Logger())

		r.Get("/health/check", health.Handle().Check)
		r.Get("/health/check/list", health.Handle().List)
	})

	// appinstance.Data.Server.Use(logger.New())
	// appinstance.Data.Server.Use(recover.New())
	// appinstance.Data.Server.Use(favicon.New())
	// appinstance.Data.Server.Use(cors.New(cors.Config{
	// 	AllowMethods: constants.AllowedMethods,
	// 	AllowOrigins: allowedOrigins,
	// 	AllowHeaders: constants.AllowedHeaders,
	// }))
	// appinstance.Data.Server.Use(middleware.ValidateContentType())
	// appinstance.Data.Server.Use(middleware.SecurityHeaders())
	// appinstance.Data.Server.Use(compress.New(compress.Config{
	// 	Level: compress.LevelBestCompression,
	// }))

	// apiGroup := appinstance.Data.Server.Group("/api")
	// apiGroup.Use(limiter.New(limiter.Config{
	// 	Next: func(c *fiber.Ctx) bool {
	// 		return helpers.Contains(constants.AllowedUnthrottledIPs, c.IP())
	// 	},
	// 	Max:        constants.MaxResquestLimit,
	// 	Expiration: 1 * time.Minute,
	// 	KeyGenerator: func(c *fiber.Ctx) string {
	// 		return c.Get("x-forwarded-for")
	// 	},
	// 	LimitReached: func(c *fiber.Ctx) error {
	// 		helpers.CreateResponse(c, &entity.ErrorResponse{
	// 			Message:     "Calls Limit Reached",
	// 			Description: "Rate Limit reached",
	// 			StatusCode:  http.StatusTooManyRequests,
	// 		}, http.StatusTooManyRequests)

	// 		return nil
	// 	},
	// }))

	// apiGroup.Get("/health", health.Handle().Check, app.Log)

	// secureRoutes := apiGroup.Group("", middleware.Authorize())
	// v1Group := secureRoutes.Group("/v1")

	// Put auth required routes here

	return r
}
