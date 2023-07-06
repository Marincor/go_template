package middleware

import "github.com/gofiber/fiber/v2"

func SecurityHeaders() func(context *fiber.Ctx) error {
	return func(context *fiber.Ctx) error {
		context.Response().Header.Add("X-XSS-Protection", "1; mode=block")
		context.Response().Header.Add("X-Content-Type-Options", "nosniff")
		context.Response().Header.Add("X-Frame-Options", "Deny")
		context.Response().Header.Add("Cache-Control", "no-store")
		context.Response().Header.Add("Content-Security-Policy", "frame-ancestors 'none'")
		context.Response().Header.Add("Content-Security-Policy", "default-src 'none'")
		context.Response().Header.Add("Referrer-Policy", "no-referrer")
		context.Response().Header.Add("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")

		return context.Next()
	}
}
