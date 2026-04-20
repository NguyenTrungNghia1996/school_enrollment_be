package adminauth

import (
	"go_be_enrollment/internal/config"
	"go_be_enrollment/internal/middleware"
	"go_be_enrollment/internal/modules/adminauth/handler"
	"go_be_enrollment/internal/modules/adminauth/repository"
	"go_be_enrollment/internal/modules/adminauth/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterAdminAuthRoutes(router fiber.Router, db *gorm.DB, cfg *config.Config) {
	repo := repository.NewAdminUserRepository(db)
	permRepo := repository.NewPermissionRepository(db)

	svc := service.NewAdminAuthService(repo, cfg)
	permSvc := service.NewPermissionService(permRepo)

	hdl := handler.NewAdminAuthHandler(svc)

	adminAuthGroup := router.Group("/admin/auth")
	{
		adminAuthGroup.Post("/login", hdl.Login)

		// Protected endpoints passing through Admin JWT Middleware
		protected := adminAuthGroup.Group("")
		protected.Use(middleware.AdminAuth(cfg))
		protected.Get("/me", hdl.GetMe)

		// Ví dụ route bảo vệ bằng PermissionGuard (Key: 'system_settings', Bit: 0)
		protected.Get("/settings", middleware.PermissionGuard(permSvc, "system_settings", 0), func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"status": "success",
				"message": "Bạn có quyền truy cập vào thông tin cấu hình hệ thống (bit 0)!",
			})
		})
	}
}
