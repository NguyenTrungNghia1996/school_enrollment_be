package middleware

import (
	"regexp"
	"time"

	"go_be_enrollment/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// regex masks key=value where key matches password, token, authorization, or secret.
var sensitiveFilters = regexp.MustCompile(`(?i)(password|token|secret|authorization)=[^&]*`)

func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		cost := time.Since(start)

		// Obfuscate secret pieces in Query string
		rawQuery := string(c.Request().URI().QueryString())
		cleanQuery := sensitiveFilters.ReplaceAllString(rawQuery, "$1=***")

		// Extract request ID optionally
		reqID := c.Locals("requestid")
		if reqID == nil {
			reqID = ""
		}

		logger.Log.Info("Incoming request",
			zap.String("req_id", reqID.(string)),
			zap.Int("status", c.Response().StatusCode()),
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.String("query", cleanQuery),
			zap.String("ip", c.IP()),
			zap.String("user-agent", string(c.Request().Header.UserAgent())),
			zap.Duration("cost", cost),
		)

		return err
	}
}
