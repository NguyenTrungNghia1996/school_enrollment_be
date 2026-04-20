package health

import (
	"go_be_enrollment/internal/common"
	"go_be_enrollment/internal/database"

	"github.com/gofiber/fiber/v2"
)

func Check(c *fiber.Ctx) error {
	dbStatus := "OK"
	
	if err := database.Ping(c.Context()); err != nil {
		dbStatus = "DOWN"
		// Respond with 503 if DB is down. Modify this if you prefer DB failure to output HTTP 200 with DEGRADED payload.
		return common.ErrorResponse(c, fiber.StatusServiceUnavailable, "Service is degraded - Database connection failed", fiber.Map{
			"status": "DEGRADED",
			"db":     dbStatus,
		})
	}

	return common.SuccessResponse(c, fiber.StatusOK, "Service is up and running", fiber.Map{
		"status": "OK",
		"db":     dbStatus,
	})
}
