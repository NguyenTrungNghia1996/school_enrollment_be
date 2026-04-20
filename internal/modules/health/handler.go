package health

import (
	"go_be_enrollment/internal/common/httpresponse"
	"go_be_enrollment/internal/database"

	"github.com/gofiber/fiber/v2"
)

func Check(c *fiber.Ctx) error {
	dbStatus := "OK"

	if err := database.Ping(c.Context()); err != nil {
		dbStatus = "DOWN"
		// Render standardized DB disconnection mapped via Error Helper 
		return httpresponse.Error(c, fiber.StatusServiceUnavailable, "DB_OFFLINE", "Service is degraded - Database connection failed", fiber.Map{
			"db": dbStatus,
		})
	}

	return httpresponse.Success(c, fiber.StatusOK, "Service is up and running", fiber.Map{
		"db": dbStatus,
	})
}
