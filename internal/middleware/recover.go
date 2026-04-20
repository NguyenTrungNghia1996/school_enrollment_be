package middleware

import (
	"fmt"
	"runtime/debug"

	"go_be_enrollment/internal/common/httpresponse"
	"go_be_enrollment/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Recovery() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				err := fmt.Errorf("%v", r)
				reqID := c.Locals("requestid")
				logger.Log.Error("Panic recovered",
					zap.Any("req_id", reqID),
					zap.Error(err),
					zap.String("stack", string(debug.Stack())),
				)
				
				// Trả về JSON tiêu chuẩn giấu thông tin core
				_ = httpresponse.ServerError(c)
			}
		}()
		return c.Next()
	}
}
