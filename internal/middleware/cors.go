package middleware

import (
	"go_be_enrollment/internal/config"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// CORS configurations dynamically bound with your environment setup
func CORS(cfg *config.Config) fiber.Handler {
	origins := cfg.CORSAllowOrigins
	// Can be comma separated or asterisk
	if origins == "" {
		origins = "*"
	}

	return cors.New(cors.Config{
		AllowOrigins: strings.Join(strings.Split(origins, ","), ", "), // Normalize formatting
		AllowHeaders: "Origin, Content-Type, Accept, Authorization, X-Request-ID",
		AllowMethods: "GET, POST, HEAD, PUT, DELETE, PATCH",
	})
}
