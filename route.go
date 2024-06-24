package main

import (
	"fmt"
	"net/http"
	"time"

	"api.default.marincor.pt/app/appinstance"
	"api.default.marincor.pt/config/constants"
	"api.default.marincor.pt/entity"
	"api.default.marincor.pt/handler/health"
	"api.default.marincor.pt/middleware"
	"api.default.marincor.pt/pkg/app"
	"api.default.marincor.pt/pkg/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func route() *fiber.App {
	allowedOrigins := constants.AllowedOrigins
	if constants.Environment != constants.Production {
		allowedOrigins += fmt.Sprintf(", %s", constants.AllowedStageOrigins)
	}

	appinstance.Data.Server.Use(logger.New())
	appinstance.Data.Server.Use(recover.New())
	appinstance.Data.Server.Use(favicon.New())
	appinstance.Data.Server.Use(cors.New(cors.Config{
		AllowMethods: constants.AllowedMethods,
		AllowOrigins: allowedOrigins,
		AllowHeaders: constants.AllowedHeaders,
	}))
	appinstance.Data.Server.Use(middleware.ValidateContentType())
	appinstance.Data.Server.Use(middleware.SecurityHeaders())
	appinstance.Data.Server.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))

	apiGroup := appinstance.Data.Server.Group("/api")
	apiGroup.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return helpers.Contains(constants.AllowedUnthrottledIPs, c.IP())
		},
		Max:        constants.MaxResquestLimit,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			helpers.CreateResponse(c, &entity.ErrorResponse{
				Message:     "Calls Limit Reached",
				Description: "Rate Limit reached",
				StatusCode:  http.StatusTooManyRequests,
			}, http.StatusTooManyRequests)

			return nil
		},
	}))

	apiGroup.Get("/health", health.Handle().Check, app.Log)

	// secureRoutes := apiGroup.Group("", middleware.Authorize())
	// v1Group := secureRoutes.Group("/v1")

	// Put auth required routes here

	return appinstance.Data.Server
}
