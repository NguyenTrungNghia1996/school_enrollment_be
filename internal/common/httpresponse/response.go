package httpresponse

import "github.com/gofiber/fiber/v2"

// Response defines standard JSON structure
type Response struct {
	Data      interface{} `json:"data,omitempty"`
	Message   string      `json:"message"`
	Status    string      `json:"status"`
	ErrorCode string      `json:"error_code,omitempty"` // Consistent error codes
}

// Success generates a standard 200/20x response
func Success(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	if data == nil {
		data = map[string]interface{}{}
	}
	return c.Status(statusCode).JSON(Response{
		Data:    data,
		Message: message,
		Status:  "success",
	})
}

// Error generates a standard client/business error response
func Error(c *fiber.Ctx, statusCode int, errorCode string, message string, errDetail interface{}) error {
	return c.Status(statusCode).JSON(Response{
		Data:      errDetail,
		Message:   message,
		Status:    "error",
		ErrorCode: errorCode,
	})
}

// ServerError generates a standard 500 error preventing sensitive detail leaks
func ServerError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(Response{
		Message:   "Trạng thái máy chủ xẩy ra lỗi bất ngờ. Xin vui lòng thử lại sau.",
		Status:    "error",
		ErrorCode: "INTERNAL_SERVER_ERROR",
	})
}
