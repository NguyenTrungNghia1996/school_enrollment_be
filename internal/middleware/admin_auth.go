package middleware

import (
	"strings"

	"go_be_enrollment/internal/common/httpresponse"
	"go_be_enrollment/internal/config"
	"go_be_enrollment/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// AdminAuth parses and validates the JWT Token for admins
func AdminAuth(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return httpresponse.Error(c, fiber.StatusUnauthorized, "UNAUTHORIZED", "Missing or invalid token", nil)
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ParseAdminToken(tokenStr, cfg.JWTSecret)
		if err != nil {
			return httpresponse.Error(c, fiber.StatusUnauthorized, "TOKEN_EXPIRED_OR_INVALID", "Token expired or invalid", nil)
		}

		// Ensure that role is admin
		if claims.Role != "admin" {
			return httpresponse.Error(c, fiber.StatusForbidden, "FORBIDDEN", "You do not have permission to access this resource", nil)
		}

		c.Locals("admin_id", claims.AdminID)
		c.Locals("admin_username", claims.Username)

		return c.Next()
	}
}
