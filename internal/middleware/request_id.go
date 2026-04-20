package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"
)

// RequestID attaches a unique UUID to every request via headers and context
func RequestID() fiber.Handler {
	return requestid.New(requestid.Config{
		Header:     "X-Request-ID",
		Generator:  func() string { return uuid.New().String() },
		ContextKey: "requestid",
	})
}
