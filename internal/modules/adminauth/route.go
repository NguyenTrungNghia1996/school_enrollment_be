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
	svc := service.NewAdminAuthService(repo, cfg)
	hdl := handler.NewAdminAuthHandler(svc)

	adminAuthGroup := router.Group("/admin/auth")
	{
		adminAuthGroup.Post("/login", hdl.Login)

		// Protected endpoints passing through Admin JWT Middleware
		protected := adminAuthGroup.Group("")
		protected.Use(middleware.AdminAuth(cfg))
		protected.Get("/me", hdl.GetMe)
	}
}
