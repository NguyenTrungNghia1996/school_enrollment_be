package middleware

import (
	"go_be_enrollment/internal/common/httpresponse"
	"go_be_enrollment/internal/modules/adminauth/service"

	"github.com/gofiber/fiber/v2"
)

// PermissionGuard acts as a middleware to check if an admin has adequate permission
func PermissionGuard(permSvc service.PermissionService, permissionKey string, permissionBit int) fiber.Handler {
	return func(c *fiber.Ctx) error {
		adminIDVal := c.Locals("admin_id")
		isSuperAdminVal := c.Locals("admin_is_super_admin")

		if adminIDVal == nil || isSuperAdminVal == nil {
			return httpresponse.Error(c, fiber.StatusUnauthorized, "UNAUTHORIZED", "Missing admin token or context", nil)
		}

		adminID := adminIDVal.(uint)
		isSuperAdmin := isSuperAdminVal.(bool)

		hasPerm, err := permSvc.CheckPermission(adminID, isSuperAdmin, permissionKey, permissionBit)
		if err != nil {
			return httpresponse.Error(c, fiber.StatusInternalServerError, "INTERNAL_ERROR", "Error checking permissions", nil)
		}

		if !hasPerm {
			return httpresponse.Error(c, fiber.StatusForbidden, "FORBIDDEN", "Bạn không có quyền thực hiện thao tác này", nil)
		}

		return c.Next()
	}
}
