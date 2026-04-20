package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// RequestMetadata injects standard metadata bounds into each transaction context
func RequestMetadata() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Store time footprint
		c.Locals("req_time", time.Now())

		// Optional: Extract User JWT claims later on down the stack into `c.Locals` too

		return c.Next()
	}
}
