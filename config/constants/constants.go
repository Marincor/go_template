package constants

import (
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const (
	Port                      = "9090"
	MainLoggerName            = "health"
	MainServiceName           = MainLoggerName + "_api"
	MaxResquestLimit          = 2
	AccessTokenExpirationTime = 15
	SignedURLExp              = 60
	DefaultLimit              = 10
	DefaultOffset             = 0
	DefaultOTPGenerator       = "iam"
)

var (
	Debug, _         = strconv.ParseBool(os.Getenv("DEBUG"))
	GcpProjectID     = os.Getenv("PROJECT")
	SecretPrefix     = os.Getenv("SEC_PREFIX")
	UseTLS           = strings.ToLower(os.Getenv("USE_TLS")) == "true"
	Environment      = os.Getenv("ENVIRONMENT")
	Prefork          = strings.ToLower(os.Getenv("PREFORK")) != "false"
	UseSecretManager = strings.ToLower(os.Getenv("USE_SECRETMANAGER")) == "true"
)

var (
	AllowedContentTypes   = []string{fiber.MIMEApplicationJSON}
	AllowedHeaders        = "X-Session-Id, Authorization, Content-Type, Accept, Origin"
	AllowedMethods        = "GET,POST,OPTIONS"
	AllowedOrigins        = "https://tbd, https://tbd"
	AllowedStageOrigins   = "https://localhost:3000, http://localhost:3000"
	AllowedUnthrottledIPs = []string{"127.0.0.1"}
)
